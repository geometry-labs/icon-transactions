package crud

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/geometry-labs/icon-transactions/config"
)

func formatPostgresDSN(host string, port string, user string, password string, dbname string, sslmode string, timezone string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
}

func getPostgresConn() *gorm.DB {
	dsn := formatPostgresDSN(
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbUser,
		config.Config.DbPassword,
		config.Config.DbName,
		config.Config.DbSslmode,
		config.Config.DbTimezone,
	)

	session, err := retryGetPostgresSession(dsn)
	if err != nil {
		zap.S().Warn("Cannot create a connection to postgres", err)
		return nil
	}

	zap.S().Info("Successful connection to postgres")
	return session
}

func createSession(dsn string) (*gorm.DB, error) {

	slowThreshold := (time.Duration(config.Config.GormLoggingThresholdMilli) * time.Millisecond)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             slowThreshold, // Slow SQL threshold
			LogLevel:                  logger.Warn,   // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		zap.S().Info("err:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(config.Config.DbMaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.Config.DbMaxOpenConnections)

	return db, err
}

func retryGetPostgresSession(dsn string) (*gorm.DB, error) {
	var session *gorm.DB
	operation := func() error {
		sess, err := createSession(dsn)
		if err != nil {
			zap.S().Info("POSTGRES SESSION Error : ", err.Error())
		} else {
			session = sess
		}
		return err
	}
	neb := backoff.NewConstantBackOff(time.Second * 3)
	err := backoff.Retry(operation, neb)
	return session, err
}
