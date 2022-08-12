package data

import (
	"context"
	"fmt"
	"strconv"

	kafka "github.com/segmentio/kafka-go"
)

// In-Memory map to access the data from the responses
var responseCollections map[string]bool

// Initalize the Map to get the response from the server
func InitializeMap(websites []string) {
	
	responseCollections = map[string]bool{}

	// Create the key values for the websites
	for _, item := range websites{

		// initalize the map with default values
		responseCollections[item] = false
	}
}

// Function for the GET query
func AccessQueryData(queryWebsite string) {
	fmt.Println("URL : ", queryWebsite, " Active : ", responseCollections[queryWebsite])
}

// Function for accessing all the websites
func AccessAllData() {
	//Display all the data from the websites
	for key := range responseCollections{
		fmt.Println("URL : ", key, " Active : ", responseCollections[key])
	}
}

// Function to read from the Kafka queue
func ReadFromTopic(r *kafka.Reader, ctx context.Context){
	
	// Read from the kafka topic at given intervals and display the message
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("cannot read from the kafka queue")
			continue
		}
		// fmt.Println("Message->  Offset : ", msg.Offset, " URL : ", string(msg.Key), "Status : ", string(msg.Value))
		convBool, _ := strconv.ParseBool(string(msg.Value))
		responseCollections[string(msg.Key)] = convBool
	}
}