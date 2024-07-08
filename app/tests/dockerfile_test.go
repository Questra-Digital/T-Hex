package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
)

func TestCreateDockerFileUI(t *testing.T) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to your React app's homepage (adjust the URL if needed)
	if err := chromedp.Run(ctx, chromedp.Navigate("http://localhost:3000")); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Wait for the Dockerfile textarea to be visible
	if err := chromedp.Run(ctx, chromedp.WaitVisible(`textarea`)); err != nil {
		log.Fatalf("Failed waiting for Dockerfile textarea: %v", err)
	}

	// Example: Type text into the Dockerfile textarea
	if err := chromedp.Run(ctx, chromedp.SetValue(`textarea`, "FROM ubuntu:latest\n")); err != nil {
		log.Fatalf("Failed to set Dockerfile textarea value: %v", err)
	}

	// Example: Click the 'Confirm' button
	if err := chromedp.Run(ctx, chromedp.Click(`button`)); err != nil {
		log.Fatalf("Failed to click 'Confirm' button: %v", err)
	}

	// Wait for the confirmation (adjust as needed)
	time.Sleep(2 * time.Second)

	// Perform assertions or additional interactions as needed
}
