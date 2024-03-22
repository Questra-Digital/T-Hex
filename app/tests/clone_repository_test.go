package tests

import (
	"context"
	"log"
	"testing"

	"github.com/chromedp/chromedp"
)

func TestCloneRepositoryUI(t *testing.T) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the Clone Repository page
	if err := chromedp.Run(ctx, chromedp.Navigate("http://localhost:3000/clone-repository")); err != nil {
		log.Fatalf("Failed to navigate to Clone Repository page: %v", err)
	}

	// Print a message to the terminal
	log.Println("Navigated to http://localhost:3000/clone-repository")

	// Wait for the input field to be visible
	if err := chromedp.Run(ctx, chromedp.WaitVisible(`input[type="text"]`)); err != nil {
		log.Fatalf("Failed waiting for input field: %v", err)
	}

	// Example: Enter repository URL
	repoURL := "https://github.com/example/repo"
	if err := chromedp.Run(ctx, chromedp.SetValue(`input[type="text"]`, repoURL, chromedp.ByQuery)); err != nil {
		log.Fatalf("Failed to set repository URL: %v", err)
	}

	// Example: Click the 'Add Repository' button
	if err := chromedp.Run(ctx, chromedp.Click(`button[data-testid="add-repo-button"]`)); err != nil {
		log.Fatalf("Failed to click 'Add Repository' button: %v", err)
	}

	// Wait for the success message to be visible
	if err := chromedp.Run(ctx, chromedp.WaitVisible(`.text-green-700`)); err != nil {
		log.Fatalf("Failed waiting for success message: %v", err)
	}

	// Example: Assert on the success message text
	var successMessageText string
	if err := chromedp.Run(ctx,
		chromedp.Text(`.text-green-700`, &successMessageText, chromedp.NodeVisible),
	); err != nil {
		log.Fatalf("Failed to get success message text: %v", err)
	}

	expectedSuccessMessage := "Project had been added successfully."
	if successMessageText != expectedSuccessMessage {
		log.Fatalf("Unexpected success message. Expected: '%s', Actual: '%s'", expectedSuccessMessage, successMessageText)
	}
}
