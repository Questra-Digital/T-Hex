package service

import (
	"Settings/configdb"
	"Settings/graph/model"
	"Settings/redisdb"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func SetProjectInfo(ctx context.Context, input model.ProjectInfoInput) (*model.Projectinfos, error) {
	// First, check if the data exists in the cache
	cachedProjectInfo, err := GetProjectInfoFromCache(ctx, input.GitProjectName)
	if err == nil && cachedProjectInfo != nil {
		return cachedProjectInfo, nil // Return the cached data
	}

	// If data doesn't exist in the cache or there's an error, fetch it from the database
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	fmt.Printf("services.go")

	projectinfos := model.Projectinfos{
		GitEmail:            input.GitEmail,
		GitProjectName:      input.GitProjectName,
		GitLanguage:         input.GitLanguage,
		GitNoOfTestCases:    input.GitNoOfTestCases,
		GitTestCaseFileName: input.GitTestCaseFileName,
	}

	// Store data in the database
	if err := db.Table("projectinfos").Create(&projectinfos).Error; err != nil {
		return nil, err
	}

	// Store data in cache
	err = SetProjectInfoToCache(ctx, input.GitProjectName, &projectinfos)
	if err != nil {
		log.Println("Failed to store data in cache:", err)
	}

	return &projectinfos, nil
}

// Function to retrieve project info from cache
// Function to retrieve project info from cache
func GetProjectInfoFromCache(ctx context.Context, projectName string) (*model.Projectinfos, error) {
	client := redisdb.ConnectRedis()

	val, err := client.Get(projectName).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // Cache miss
		}
		return nil, err
	}

	// Unmarshal cached value
	var projectInfo model.Projectinfos
	err = json.Unmarshal([]byte(val), &projectInfo)
	if err != nil {
		return nil, err
	}

	return &projectInfo, nil
}

func SetProjectInfoToCache(ctx context.Context, projectName string, projectInfo *model.Projectinfos) error {
	client := redisdb.ConnectRedis()

	// Marshal project info
	data, err := json.Marshal(projectInfo)
	if err != nil {
		return err
	}

	// Store data in cache
	err = client.Set(projectName, data, 0).Err() // 0 means no expiration
	if err != nil {
		return err
	}

	return nil
}

func SetSettings(ctx context.Context, input model.GetSettingInput) (*model.Settings, error) {

	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	fmt.Printf("services.go")

	settings := model.Settings{
		Browser:             input.Browser,
		Version:             input.Version,
		StepByStepDebugging: input.StepByStepDebugging,
		EnableLogs:          input.EnableLogs,
		Parallelism:         input.Parallelism,
	}

	if err := db.Table("settings").Create(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}

func GetAllSettings(ctx context.Context) ([]*model.Settings, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var settings []*model.Settings
	if err := db.Table("settings").Find(&settings).Error; err != nil {
		return nil, err
	}

	return settings, nil
}
func GetAllProjectInfos(ctx context.Context) ([]*model.Projectinfos, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var projectinfos []*model.Projectinfos
	if err := db.Table("projectinfos").Find(&projectinfos).Error; err != nil {
		return nil, err
	}

	return projectinfos, nil
}

func GetProjectInfoByID(ctx context.Context, id int) (*model.Projectinfos, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var projectinfos model.Projectinfos
	if err := db.Table("projectinfos").Where("id = ?", id).Take(&projectinfos).Error; err != nil {
		return nil, err
	}

	return &projectinfos, nil
}
func GetSettingByID(ctx context.Context, id int) (*model.Settings, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var settings model.Settings
	if err := db.Table("settings").Where("id = ?", id).Take(&settings).Error; err != nil {
		return nil, err
	}

	return &settings, nil
}
