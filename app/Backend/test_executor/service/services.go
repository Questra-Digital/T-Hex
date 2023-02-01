package service

import (
	"context"
	"test_executor/configdb"
	"test_executor/graph/model"
)

func StartTest(ctx context.Context, input model.StartTestInput) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	tests := model.Test{ //insert test into DB
		Status:    "Pending", //table attributes
		TestPath: input.TestPath,
	}
	
	if err := db.Table("tests").Create(&tests).Error; err != nil {
		return nil, err
	}
	return &tests, nil
}

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
