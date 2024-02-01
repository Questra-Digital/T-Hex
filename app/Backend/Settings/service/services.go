package service

import (
	"Settings/configdb"
	"Settings/graph/model"
	"context"
	"fmt"
)

func SetProjectInfo(ctx context.Context, input model.ProjectInfoInput) (*model.Projectinfos, error) {
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

	if err := db.Table("projectinfos").Create(&projectinfos).Error; err != nil {
		return nil, err
	}
	return &projectinfos, nil

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
