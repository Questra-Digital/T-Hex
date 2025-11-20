package routes

import (
	"cicd_runner/config"
	"cicd_runner/utils"
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func PullTestCasesFromRepository(c *gin.Context) {
	var requestBody struct {
		RepositoryPath string `json:"repository_path"`
		AccessToken    string `json:"access_token"`
		APIKey         string `json:"api_key"`
		PipelineName   string `json:"pipeline_name"`
	}

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("Error binding JSON", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	//1. First get the test id from the reverse proxy
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/thex/test", config.EnvConfig.ReverseProxyUrl), nil)
	if err != nil {
		log.Printf("Failed to request TestId from THex: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	req.Header.Add("thex-key", requestBody.APIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to request TestId from THex: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	testId := resp.Header.Get("thex-test")
	if testId == "" {
		log.Printf("THex refused to send TestId")
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "THex refused to send TestId"})
		return
	}
	resp.Body.Close()
	log.Printf("Received TestId: %s", testId)

	err = utils.PullTestCases(requestBody.RepositoryPath, requestBody.AccessToken)
	if err != nil {
		log.Println("Error pulling test cases", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

		//2. Setup the parameters for the reverse proxy
	config.ProxyConfig.SetValues(requestBody.APIKey, requestBody.PipelineName, testId)
	// Log the ProxyConfig values being set
	log.Println("ProxyConfig values set successfully")
	//3. Execute the test cases
	err = utils.ExecuteTestCases()
	if err != nil {
		log.Println("Error executing test cases", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	//4. Delete the test cases
	err = utils.DeleteTestCases()
	if err != nil {
		log.Println("Error deleting test cases", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Test cases pulled and executed successfully"})
}
