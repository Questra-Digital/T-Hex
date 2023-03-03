package service

import (
	"Settings/configdb"
	"Settings/graph/model"
	"context"
)

func setSettings(ctx context.Context, input model.GetSettingtInput) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	settings := model.Settings{ //insert test into DB
		// Status:   "Pending", //table attributes
		// TestPath: input.TestPath,
		Browser:input.Browser,
		Version:input.version,
		StepByStepDebugging:input.StepByStepDebugging,
		EnableLogs:input.EnableLogs,
		Parallelism:input.Parallelism,
		NumberOfParallelTests:input.NumberOfParallelTests
	}

	if err := db.Table("settings").Create(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}

func StartTest(ctx context.Context, input model.StartTestInput) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	tests := model.Test{ //insert test into DB
		Status:   "Pending", //table attributes
		TestPath: input.TestPath,
	}

	if err := db.Table("tests").Create(&tests).Error; err != nil {
		return nil, err
	}
	return &tests, nil
}

// func SelemiumPull(ctx context.Context)(string,error){

// }

// List of Tests
func GetAllTest(ctx context.Context) ([]*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var tests []*model.Test
	if err := db.Table("tests").Find(&tests).Error; err != nil {
		return nil, err
	}

	return tests, nil
}

func GetTestByID(ctx context.Context, id int) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var tests model.Test
	if err := db.Table("tests").Where("id = ?", id).Take(&tests).Error; err != nil {
		return nil, err
	}

	return &tests, nil
}

func UpdateTestByID(ctx context.Context, input model.UpdateTestInput) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// var tests model.Test
	if err := db.Table("tests").Where("id = ?", input.ID).Update("status", input.Status).Error; err != nil {
		return nil, err
	}
	if err := db.Table("tests").Where("id = ?", input.ID).Update("test_path", input.TestPath).Error; err != nil {
		return nil, err
	}
	return GetTestByID(ctx, input.ID)
}
