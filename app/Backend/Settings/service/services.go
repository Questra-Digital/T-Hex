package service

import (
	"Settings/configdb"
	"Settings/graph/model"
	"context"
)

func SetSettings(ctx context.Context, input model.GetSettingInput) (*model.Settings, error) {
	db := configdb.ConnectCockroachDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	settings := model.Settings{ //insert test into DB
		// Status:   "Pending", //table attributes
		// TestPath: input.TestPath,
		Browser:               input.Browser,
		Version:               input.Version,
		StepByStepDebugging:   input.StepByStepDebugging,
		EnableLogs:            input.EnableLogs,
		Parallelism:           input.Parallelism,
		NumberOfParallelTests: input.NumberOfParallelTests,
	}

	if err := db.Table("settings").Create(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}
