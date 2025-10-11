package db

import (
	"cicd_runner/config"
	"cicd_runner/models"
)

func CreatePipeline(pipeline *models.Pipeline, accessToken string) error {

	// First create the pipeline to get an ID
	if err := config.DB.Create(&pipeline).Error; err != nil {
		return err
	}

	// Then create the access token with the pipeline ID
	accessTokenModel := models.AccessTokens{
		AccessToken: accessToken,
		PipelineID:  pipeline.ID,
	}

	if err := config.DB.Create(&accessTokenModel).Error; err != nil {
		return err
	}

	return nil
}

func CreatePipelineEvent(pipelineEvent *models.PipelineEvent) error {
	return config.DB.Create(&pipelineEvent).Error
}

func GetPipelineIdByRepositoryPath(repositoryPath string) (uint, error) {
	var pipeline models.Pipeline
	if err := config.DB.Where("repository_path = ?", repositoryPath).First(&pipeline).Error; err != nil {
		return 0, err
	}
	return pipeline.ID, nil
}

func GetPipelines() ([]models.Pipeline, error) {
	var pipelines []models.Pipeline
	if err := config.DB.Find(&pipelines).Error; err != nil {
		return nil, err
	}
	var err error
	for i, pipeline := range pipelines {
		pipelines[i].Events, err = GetPipelineEvents(pipeline.ID)
		if err != nil {
			return nil, err
		}
	}

	return pipelines, nil
}

func GetPipelineEvents(pipelineID uint) ([]models.PipelineEvent, error) {
	var events []models.PipelineEvent
	if err := config.DB.Where("pipeline_id = ?", pipelineID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
