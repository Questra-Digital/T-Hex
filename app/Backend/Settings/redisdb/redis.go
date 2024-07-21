package redisdb

import (
	"log"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func ConnectRedis() *redis.Client {
	// Connect to Redis server
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Ping the Redis server to check if the connection was successful
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal("failed to connect to Redis:", err)
	}
	log.Println("Connected to Redis:", pong)

	return RedisClient
}
