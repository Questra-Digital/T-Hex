package service

import (
	"Settings/configdb"
	"Settings/graph/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func init() {
	// Initialize Redis client
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // replace with your Redis server address
		Password: "",               // no password by default
		DB:       0,                // use default DB
	})
}

// SetProjectInfo stores project info in both CockroachDB and Redis cache
func SetProjectInfo(ctx context.Context, input model.ProjectInfoInput) (*model.Projectinfos, error) {
	// Try to retrieve data from Redis cache first
	cacheKey := fmt.Sprintf("projectInfo:%s", input.GitProjectName)
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data found in cache, parse and return
		var projectinfos model.Projectinfos
		if err := json.Unmarshal([]byte(cachedData), &projectinfos); err == nil {
			return &projectinfos, nil
		}
	}

	// Data not found in cache, fetch from CockroachDB
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	projectinfos := model.Projectinfos{
		GitEmail:            input.GitEmail,
		GitProjectName:      input.GitProjectName,
		GitLanguage:         input.GitLanguage,
		GitNoOfTestCases:    input.GitNoOfTestCases,
		GitTestCaseFileName: input.GitTestCaseFileName,
	}

	// Save to CockroachDB
	if err := db.Table("projectinfos").Create(&projectinfos).Error; err != nil {
		return nil, err
	}

	// Save to Redis cache
	jsonData, _ := json.Marshal(projectinfos)
	redisClient.Set(ctx, cacheKey, jsonData, 24*time.Hour) // You can adjust the expiration time

	return &projectinfos, nil
}

// SetSettings stores settings in both CockroachDB and Redis cache
func SetSettings(ctx context.Context, input model.GetSettingInput) (*model.Settings, error) {
	// Similar implementation as SetProjectInfo, adapt as needed
}

// GetAllSettings retrieves all settings from Redis cache or CockroachDB
func GetAllSettings(ctx context.Context) ([]*model.Settings, error) {
	// Try to retrieve data from Redis cache first
	cacheKey := "allSettings"
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data found in cache, parse and return
		var settings []*model.Settings
		if err := json.Unmarshal([]byte(cachedData), &settings); err == nil {
			return settings, nil
		}
	}

	// Data not found in cache, fetch from CockroachDB
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var settings []*model.Settings
	if err := db.Table("settings").Find(&settings).Error; err != nil {
		return nil, err
	}

	// Save to Redis cache
	jsonData, _ := json.Marshal(settings)
	redisClient.Set(ctx, cacheKey, jsonData, 24*time.Hour) // You can adjust the expiration time

	return settings, nil
}

// GetAllProjectInfos retrieves all project infos from Redis cache or CockroachDB
func GetAllProjectInfos(ctx context.Context) ([]*model.Projectinfos, error) {
	// Try to retrieve data from Redis cache first
	cacheKey := "allProjectInfos"
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data found in cache, parse and return
		var projectinfos []*model.Projectinfos
		if err := json.Unmarshal([]byte(cachedData), &projectinfos); err == nil {
			return projectinfos, nil
		}
	}

	// Data not found in cache, fetch from CockroachDB
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var projectinfos []*model.Projectinfos
	if err := db.Table("projectinfos").Find(&projectinfos).Error; err != nil {
		return nil, err
	}

	// Save to Redis cache
	jsonData, _ := json.Marshal(projectinfos)
	redisClient.Set(ctx, cacheKey, jsonData, 24*time.Hour) // You can adjust the expiration time

	return projectinfos, nil
}

// GetProjectInfoByID retrieves project info by ID from Redis cache or CockroachDB
func GetProjectInfoByID(ctx context.Context, id int) (*model.Projectinfos, error) {
	// Try to retrieve data from Redis cache first
	cacheKey := fmt.Sprintf("projectInfo:%d", id)
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data found in cache, parse and return
		var projectinfos model.Projectinfos
		if err := json.Unmarshal([]byte(cachedData), &projectinfos); err == nil {
			return &projectinfos, nil
		}
	}

	// Data not found in cache, fetch from CockroachDB
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var projectinfos model.Projectinfos
	if err := db.Table("projectinfos").Where("id = ?", id).Take(&projectinfos).Error; err != nil {
		return nil, err
	}

	// Save to Redis cache
	jsonData, _ := json.Marshal(projectinfos)
	redisClient.Set(ctx, cacheKey, jsonData, 24*time.Hour) // You can adjust the expiration time

	return &projectinfos, nil
}

// GetSettingByID retrieves setting by ID from Redis cache or CockroachDB
func GetSettingByID(ctx context.Context, id int) (*model.Settings, error) {
	// Try to retrieve data from Redis cache first
	cacheKey := fmt.Sprintf("setting:%d", id)
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data found in cache, parse and return
		var settings model.Settings
		if err := json.Unmarshal([]byte(cachedData), &settings); err == nil {
			return &settings, nil
		}
	}

	// Data not found in cache, fetch from CockroachDB
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var settings model.Settings
	if err := db.Table("settings").Where("id = ?", id).Take(&settings).Error; err != nil {
		return nil, err
	}

	// Save to Redis cache
	jsonData, _ := json.Marshal(settings)
	redisClient.Set(ctx, cacheKey, jsonData, 24*time.Hour) // You can adjust the expiration time

	return &settings, nil
}

// CloseRedisConnection closes the Redis connection when the application exits
func CloseRedisConnection() {

	_ = redisClient.Close()
}
