package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func DeleteTestCases() error {

	// Check if test_cases directory exists if it doesn't exist, throw error
	testCasesDir := "/app/test_cases"
	if _, err := os.Stat(testCasesDir); os.IsNotExist(err) {
		return fmt.Errorf("test_cases directory not found")
	}

	// Read all files/directories inside
	entries, err := os.ReadDir(testCasesDir)
	if err != nil {
		return fmt.Errorf("failed to read test_cases directory: %v", err)
	}

	// Delete each entry inside the directory
	for _, entry := range entries {
		entryPath := filepath.Join(testCasesDir, entry.Name())
		if err := os.RemoveAll(entryPath); err != nil {
			return fmt.Errorf("failed to delete %s: %v", entryPath, err)
		}
	}

	log.Println("Test cases deleted successfully")
	return nil
}
