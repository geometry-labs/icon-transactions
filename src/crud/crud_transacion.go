package crud

import (
	"strings"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/geometry-labs/icon-transactions/models"
)

// TransactionModel - type for transaction table model
type TransactionModel struct {
	db        *gorm.DB
	model     *models.Transaction
	modelORM  *models.TransactionORM
	WriteChan chan *models.Transaction
}

var transactionModel *TransactionModel
var transactionModelOnce sync.Once

// GetTransactionModel - create and/or return the transactions table model
func GetTransactionModel() *TransactionModel {
	transactionModelOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		transactionModel = &TransactionModel{
			db:        dbConn,
			model:     &models.Transaction{},
			WriteChan: make(chan *models.Transaction, 1),
		}

		err := transactionModel.Migrate()
		if err != nil {
			zap.S().Fatal("TransactionModel: Unable migrate postgres table: ", err.Error())
		}
	})

	return transactionModel
}

// Migrate - migrate transactions table
func (m *TransactionModel) Migrate() error {
	// Only using TransactionRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

// Insert - Insert transaction into table
func (m *TransactionModel) Insert(transaction *models.Transaction) error {

	err := backoff.Retry(func() error {
		query := m.db.Create(transaction)
		if query.Error != nil && !strings.Contains(query.Error.Error(), "duplicate key value violates unique constraint") {
			zap.S().Warn("POSTGRES Insert Error : ", query.Error.Error())
			return query.Error
		}

		return nil
	}, backoff.NewExponentialBackOff())

	return err
}

// SelectMany - select from transactions table
// Returns: models, total count (if filters), error (if present)
func (m *TransactionModel) SelectMany(
	limit int,
	skip int,
	hash string,
	from string,
	to string,
) ([]models.TransactionAPI, int64, error) {
	db := m.db
	computeCount := false

	// Set table
	db = db.Model(&[]models.Transaction{})

	// Latest transactions first
	db = db.Order("block_number desc")

	// Hash
	if hash != "" {
		computeCount = true
		db = db.Where("hash = ?", hash)
	}

	// from
	if from != "" {
		computeCount = true
		db = db.Where("from_address = ?", from)
	}

	// to
	if to != "" {
		computeCount = true
		db = db.Where("to_address = ?", to)
	}

	// Count, if needed
	count := int64(-1)
	if computeCount {
		db.Count(&count)
	}

	// Limit is required and defaulted to 1
	// Note: Count before setting limit
	db = db.Limit(limit)

	// Skip
	// Note: Count before setting skip
	if skip != 0 {
		db = db.Offset(skip)
	}

	transactions := []models.TransactionAPI{}
	db = db.Find(&transactions)

	return transactions, count, db.Error
}

// SelectOne - select from transactions table
func (m *TransactionModel) SelectOne(
	hash string,
) (models.Transaction, error) {
	db := m.db

	// Hash
	if hash != "" {
		db = db.Where("hash = ?", hash)
	}

	transaction := models.Transaction{}
	db = db.First(&transaction)

	return transaction, db.Error
}

// StartTransactionLoader starts loader
func StartTransactionLoader() {
	go func() {
		var transaction *models.Transaction
		postgresLoaderChan := GetTransactionModel().WriteChan

		for {
			// Read transaction
			transaction = <-postgresLoaderChan

			// Load transaction to database
			GetTransactionModel().Insert(transaction)

			zap.S().Debugf("Loader Transaction: Loaded in postgres table Transactions, Block Number: %d", transaction.BlockNumber)
		}
	}()
}