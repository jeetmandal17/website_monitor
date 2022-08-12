package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/monitorServer/commons"
	"github.com/monitorServer/monitor_server/types"

	kafka "github.com/segmentio/kafka-go"
)

func main(){
	// Initialize the context 
	ctx := context.Background()

	// We will create the kafka event stream to send the data among the 2 services
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{commons.BrokerAddress1,commons.BrokerAddress2},
		Topic: commons.ResponseTopic,
	})

	// Get the websites information into the global in-memory collection
	newWebsitesList := []string{"www.google.com","www.facebook.com","xyz"}

	// Send the newList into in-memory database [Send String Array]
	types.AddWebsites(newWebsitesList)

	// Run the go routine to ping the server every 1 minute
	for {
		finalResponses := types.PingWebsites()

		for _, item := range finalResponses{
			fmt.Println("URL : ", item.URL, "Status : ", item.Active)
			// Write this response into the Kafka Topic
			err := w.WriteMessages(ctx, kafka.Message{
				Key: []byte(item.URL),
				Value: []byte(strconv.FormatBool(item.Active)),
			},
			)

			if err != nil {
				fmt.Println("Cannot Write the message into Kafka Topic")
			}
		}

		//Repeat this every minute
		time.Sleep(10*time.Second) 
	}
}