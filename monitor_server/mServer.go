package main

import (
	"fmt"
	"time"

	"github.com/monitorServer/monitor_server/types"
	// kafka "github.com/segmentio/kafka-go"
)
const (
	topic = "webstie-uptime"
	brokerAddress1 = "localhost:9092"
	brokerAddress2 = "localhost:9093"
)


func main(){

	// We will create the kafka event stream to send the data among the 2 services
	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{brokerAddress1,brokerAddress2},
	// 	Topic: topic,
	// })

	// Get the websites information into the global in-memory collection
	newWebsitesList := []string{"www.google.com","www.facebook.com"}

	// Send the newList into in-memory database [Send String Array]
	types.AddWebsites(newWebsitesList)

	// Run the go routine to ping the server every 1 minute
	for {
		finalResponses := types.PingWebsites()

		for _, item := range finalResponses{
			fmt.Println("URL : ", item.URL, "Status : ", item.Active)
		}

		//Repeat this every minute
		time.Sleep(10*time.Second) 
	}
}