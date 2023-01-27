package service

import (
	"context"
	"go-app/configdb"
	"go-app/graph/model"
)

// Create User in DB
func AddUser(ctx context.Context, input model.AddUserInput) (*model.User, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	users := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := db.Table("users").Create(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

// Creates Test in Db
func AddTest(ctx context.Context, input model.AddTestInput) (*model.Test, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// var user model.User // user to send to test table
	// user.Name = "Tester User"
	// user.Email = "tester@gmail.com"
	// user.Password = "tester"

	tests := model.Test{ //insert test into DB
		Status:    input.Status, //table attributes
		Startedby: input.Startedby,
	}
	if err := db.Table("tests").Create(&tests).Error; err != nil {
		return nil, err
	}

	return &tests, nil
}

// List of user
func GetAllUser(ctx context.Context) ([]*model.User, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users []*model.User
	if err := db.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
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

// Get User By NAme	//User NAme as Parameter
func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users model.User
	if err := db.Table("users").Where("name = ?", name).Take(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
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
