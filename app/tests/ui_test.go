package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestHomePageUI(t *testing.T) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to your React app's homepage
	if err := chromedp.Run(ctx, chromedp.Navigate("http://localhost:3000")); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Print a message to the terminal
	log.Println("Navigated to http://localhost:3000")

	// Wait for the hero title to be visible
	if err := chromedp.Run(ctx, chromedp.WaitVisible(`.HomePage_heroTitle___iHRi`)); err != nil {
		log.Fatalf("Failed waiting for hero title: %v", err)
	}
	// Print a message to the terminal
	log.Println("Navigated to hero")

	// Example: Perform assertions on the hero title and description
	var titleText, descriptionText string
	if err := chromedp.Run(ctx,
		chromedp.Text(`.HomePage_heroTitle___iHRi`, &titleText, chromedp.NodeVisible),
		chromedp.Text(`.HomePage_heroDescription__iRdur`, &descriptionText, chromedp.NodeVisible),
	); err != nil {
		log.Fatalf("Failed to get text: %v", err)
	}

	expectedTitle := "Web Test Hub"
	if titleText != expectedTitle {
		log.Fatalf("Unexpected hero title. Expected: '%s', Actual: '%s'", expectedTitle, titleText)
	}

	expectedDescription := "WebTestHub is a cutting-edge test automation software..."
	if descriptionText != expectedDescription {
		log.Fatalf("Unexpected hero description. Expected: '%s', Actual: '%s'", expectedDescription, descriptionText)
	}

	// Perform additional interactions and assertions as needed

	// Example: Click the 'Get Started' button
	if err := chromedp.Run(ctx, chromedp.Click(`.HomePage_heroButton__8RLDq`)); err != nil {
		log.Fatalf("Failed to click 'Get Started' button: %v", err)
	}

	// Wait for the page to navigate (adjust as needed)
	time.Sleep(2 * time.Second)

	// Example: Assert on the new page title
	var newPageTitle string
	if err := chromedp.Run(ctx, chromedp.Title(&newPageTitle)); err != nil {
		log.Fatalf("Failed to get new page title: %v", err)
	}

	expectedNewPageTitle := "Clone Repository"
	if newPageTitle != expectedNewPageTitle {
		log.Fatalf("Unexpected new page title. Expected: '%s', Actual: '%s'", expectedNewPageTitle, newPageTitle)
	}
}
