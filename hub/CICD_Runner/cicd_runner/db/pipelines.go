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

func GetPipelines(userId int64) ([]models.Pipeline, error) {
	var pipelines []models.Pipeline
	if err := config.DB.Where("user_id = ?", userId).Find(&pipelines).Error; err != nil {
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

// User Related Functions
func GetUserIdByEmail(email string) (int64, error) {
	var userId int64
	if err := config.DB.Model(&models.User{}).Where("email = ?", email).Select("id").First(&userId).Error; err != nil {
		return 0, err
	}
	return userId, nil
}

// /Access Related functions
func GetAccessTokenByPipelineId(pipelineId uint) (string, error) {
	var accessToken models.AccessTokens
	if err := config.DB.Where("pipeline_id = ?", pipelineId).First(&accessToken).Error; err != nil {
		return "", err
	}
	return accessToken.AccessToken, nil
}

func GetAPIKeyByUserId(userId int64) (string, error) {
	var apiKey models.Keys
	if err := config.DB.Where("user_id = ?", userId).First(&apiKey).Error; err != nil {
		return "", err
	}
	return apiKey.APIKey, nil
}

func GetUserIdByPipelineId(pipelineId uint) (int64, error) {
	var userId int64
	if err := config.DB.Model(&models.Pipeline{}).Where("id = ?", pipelineId).Select("user_id").First(&userId).Error; err != nil {
		return 0, err
	}
	return userId, nil
}

func GetPipelineNameByPipelineId(pipelineId uint) (string, error) {
	var pipelineName string
	if err := config.DB.Model(&models.Pipeline{}).Where("id = ?", pipelineId).Select("name").First(&pipelineName).Error; err != nil {
		return "", err
	}
	return pipelineName, nil
}

func UpdatePipelineEvent(pipelineEvent *models.PipelineEvent) error {
	return config.DB.Model(&models.PipelineEvent{}).Where("id = ?", pipelineEvent.ID).Updates(pipelineEvent).Error
}

// //Getters for Test Sessions/////////////////////////////////////////////////////////////
func GetTestSession(pipelineId uint, pipelineEventId uint, userId int64) (models.TestSession, error) {
	//First get the pipeline
	var pipeline models.Pipeline
	if err := config.DB.Where("id = ? AND user_id = ?", pipelineId, userId).First(&pipeline).Error; err != nil {
		return models.TestSession{}, err
	}
	//Then get the pipeline event
	var pipelineEvent models.PipelineEvent
	if err := config.DB.Where("id = ? AND pipeline_id = ?", pipelineEventId, pipeline.ID).First(&pipelineEvent).Error; err != nil {
		return models.TestSession{}, err
	}
	//Then get the test session
	var testSession models.TestSession
	if err := config.DB.Where("pipeline_event_id = ?", pipelineEvent.ID).First(&testSession).Error; err != nil {
		return models.TestSession{}, err
	}

	return testSession, nil
}

func VerifyUserId(userId int64) (int64, error) {
	var user models.User

	if err := config.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func GetSeleniumSessionsByTestId(testId int64) ([]models.SessionSelenium, error) {
	var seleniumSessions []models.SessionSelenium
	if err := config.DB.Where("test_id = ?", testId).Find(&seleniumSessions).Error; err != nil {
		return nil, err
	}
	return seleniumSessions, nil
}

func GetSeleniumSessionBySessionId(sessionId string, testId int64) (models.SessionSelenium, error) {
	var session models.SessionSelenium
	if err := config.DB.Where("session_id = ? AND test_id = ?", sessionId, testId).First(&session).Error; err != nil {
		return models.SessionSelenium{}, err
	}
	return session, nil
}

func GetSeleniumEventsBySessionId(sessionId string) ([]models.Event, error) {
	var events []models.Event
	if err := config.DB.Where("session_id = ?", sessionId).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}