package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
)

func main() {

	const (
		hubURL = "http://localhost:4444/wd/hub"
	)

	// Create a new instance of the Chrome driver
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, hubURL)
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}
	defer wd.Quit()

	// Get the Google homepage
	if err := wd.Get("https://www.google.com"); err != nil {
		log.Fatalf("Failed to load page: %v", err)
	}

	// Find the search field and enter "Selenium"
	searchField, err := wd.FindElement(selenium.ByName, "q")
	if err != nil {
		log.Fatalf("Failed to find search field: %v", err)
	}
	if err := searchField.SendKeys("Selenium"); err != nil {
		log.Fatalf("Failed to enter text intosearch field: %v", err)
	}

	//Copy code
	// Submit the search form
	searchField.Submit()

	// Wait for the search results to load
	//time.Sleep(time.Second * 5)

	// Get the title of the search results page
	title, err := wd.Title()
	if err != nil {
		log.Fatalf("Failed to get page title: %v", err)
	}

	fmt.Printf("Search results page title: %s\n", title)
}
