package types

import (
	"fmt"
	"net/http"
)

// Make a in-memory storage to save the websites to be monitored
// GLOBAL Array
var websiteCollection []Website

// Utility to add a website in the in-memory list
func AddWebsites(newWebsiteCollection []string) {
	// Check if the websiteCollection is already in use
	if len(websiteCollection) != 0 {
		fmt.Println("Overwriting the previous monitoring list")
	}

	// Add these list into the in-memory
	newTransformedList := []Website{}

	// Add it into the newTransformed list
	for _,item := range newWebsiteCollection {
		newWebsite := NewWebsite(item)
		
		// Append it into the list
		newTransformedList = append(newTransformedList, *newWebsite)
	}
	// Initialize the new website lists
	websiteCollection = newTransformedList
}

// Go-Routine to ping to the list of websites
func PingWebsites() (responseCollections []WebsiteResponse) {
	
	// This stores the responses
	receivedResponses := []WebsiteResponse{}

	// Create a new slice for storing the 
	// Create clients to store the responses 
	for _, item := range websiteCollection{

		// We now make individual connection to each website
		webResponse, err := WebsiteStatusUpdate(item.URL)
		if err != nil {
			fmt.Println("failed to connect to the site")
		}
		receivedResponses = append(receivedResponses, *webResponse)
	}
	
	// this contains the received responses
	return receivedResponses
}

// Function to ping to the URL we got from the POST request
func WebsiteStatusUpdate(URL string) (*WebsiteResponse, error) {

	// Here we ping the individual website for status
	httpURL := "https://" + URL
	_, err := http.Get(httpURL)

	// Check for the errors 
	if err != nil {
		return NewResponseWebsite(httpURL, false), nil
	}

	// Check for the responses and return the ResponseStructure
	return NewResponseWebsite(httpURL, true), nil
}