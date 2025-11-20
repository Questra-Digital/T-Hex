package utils

import (
	"cicd_runner/config"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ExecuteTestCases() error {

	// Check if test_cases directory exists if it doesn't exist, throw error
	testCasesDir := "/app/test_cases"
	if _, err := os.Stat(testCasesDir); os.IsNotExist(err) {
		return fmt.Errorf("test_cases directory not found")
	}

	executorUrl := fmt.Sprintf("http://host.docker.internal%s/reverse_proxy", config.EnvConfig.Port)

	//Read files in the test_cases directory and execute them
	files, err := os.ReadDir(testCasesDir)
	if err != nil {
		return fmt.Errorf("failed to read test cases directory: %v", err)
	}

	for _, file := range files {

		filePath := filepath.Join(testCasesDir, file.Name())
		log.Printf("Executing test case: %s", filePath)

		cmdPython := exec.Command("python", filePath, executorUrl)
		cmdPython.Stdout = os.Stdout
		cmdPython.Stderr = os.Stderr
		if err := cmdPython.Run(); err != nil {
			return fmt.Errorf("failed to execute test case: %v", err)
		}

		log.Printf("Test case %s executed successfully", file.Name())

	}

	log.Println("All test cases executed successfully")

	return nil
}
