package redis

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/geometry-labs/icon-transactions/config"
)

type Client struct {
	client *redis.Client
	pubsub *redis.PubSub
}

var redisClient *Client
var redisClientOnce sync.Once

func GetRedisClient() *Client {
	redisClientOnce.Do(func() {
		addr := config.Config.RedisHost + ":" + config.Config.RedisPort

		redisClient = new(Client)

		// Init connection
		if config.Config.RedisSentinelClientMode == false {
			// Use default client
			redisClient.client = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: config.Config.RedisPassword,
				DB:       0,
			})
		} else {
			// Use sentinel client
			redisClient.client = redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:    config.Config.RedisSentinelClientMasterName,
				SentinelAddrs: []string{addr},
			})
		}

		if redisClient.client == nil {
			zap.S().Fatal("RedisClient: Unable to create to redis client")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Test connection
		_, err := redisClient.client.Ping(ctx).Result()
		if err != nil {
			zap.S().Fatal("RedisClient: Unable to connect to redis", err.Error())
		}

		// Init pubsub
		redisClient.pubsub = redisClient.client.Subscribe(ctx, config.Config.RedisChannel)

		// Test pubsub
		_, err = redisClient.pubsub.Receive(ctx)
		if err != nil {
			zap.S().Fatal("RedisClient: Unable to create pubsub channel")
		}
	})

	return redisClient
}

func (c *Client) Publish(data []byte) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Publish
		err := c.client.Publish(ctx, config.Config.RedisChannel, string(data)).Err()
		if err != nil {
			// Failure
			zap.S().Warn("Redis Publish: Cannot publish message...retrying in 3 second")
			time.Sleep(3 * time.Second)

			continue
		}

		// Success
		break
	}
}

func (c *Client) StartSubscriber() {

	go func() {
		subscriberChannel := c.pubsub.Channel()
		inputChannel := GetBroadcaster().InputChannel

		for {
			redisMsg := <-subscriberChannel

			inputChannel <- []byte(redisMsg.Payload)
		}
	}()
}
