package utils

import (
	"server/models"
	"fmt"
)


func DeleteState(state string) bool{
	_, exists := models.StateStore.Map[state]
	if exists {
		delete(models.StateStore.Map, state)
		fmt.Println("State validated and deleted:", state)
	}
	return exists
}

// ValidateState checks if the state exists and deletes it to prevent replay attacks
func ValidateState(state string) bool {
	models.StateStore.Lock()
	defer models.StateStore.Unlock()

	exists := DeleteState(state)
	return exists
}