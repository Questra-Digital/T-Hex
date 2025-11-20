package utils

import (
	"ReverseProxyServer/config"
	"ReverseProxyServer/models"
	"strings"
	"time"
)

// makes a new Testing Session for a key and proj
func TestSessCreate(key string, proj string, pipelineEventID int64) (int64, error) {
	test := models.TestSession{
		Time:            time.Now().Unix(),
		Key:             key,
		Proj:            proj,
		Current:         true,
		PipelineEventID: pipelineEventID,
	}
	err := config.DB.Create(&test).Error
	if err != nil {
		return 0, err
	}
	return test.TestId, nil
}

// ends a testing session
func TestSessEnd(testId int64) error {
	err := config.DB.Model(&models.TestSession{}).
		Where("test_id = ?", testId).
		Update("Current", false).
		Error
	if err != nil {
		return err
	}
	return nil
}

// determines if a session had any critical errors based on logged events
func DetermineSessionStatus(sessionId string) bool {
	var events []models.Event
	err := config.DB.Where("session_id = ?", sessionId).Find(&events).Error
	if err != nil {
		// If we can't fetch events, assume failure for safety
		return false
	}

	for _, event := range events {
		// Check for HTTP error status codes (4xx, 5xx)
		if event.Status >= 400 && event.Status < 600 {
			return false
		}

		// Check for critical selenium error patterns in response body
		responseBody := strings.ToLower(event.Res)
		criticalErrors := []string{
			"nosuchelement",
			"timeoutexception",
			"staleelementreference",
			"elementnotinteractable",
			"invalidselectorexception",
			"webdriverexception",
			"session not created",
			"invalid session id",
			"no such window",
			"no such frame",
		}

		for _, errorPattern := range criticalErrors {
			if strings.Contains(responseBody, errorPattern) {
				return false
			}
		}
	}

	// If no critical errors found, test is considered successful
	return true
}

// updates session status and message based on event analysis
func UpdateSessionStatus(sessionId string) (bool, error) {
	success := DetermineSessionStatus(sessionId)
	message := ""
	if success {
		message = "Test completed successfully"
	} else {
		message = "Test failed"
	}

	err := config.DB.Model(&models.SessionSelenium{}).
		Where("session_id = ?", sessionId).
		Updates(map[string]interface{}{
			"status":  success,
			"message": message,
		}).Error

	return success, err
}
