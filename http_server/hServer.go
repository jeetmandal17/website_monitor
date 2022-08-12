package main

import (
	"context"
	"sync"
	"time"

	"github.com/monitorServer/commons"
	"github.com/monitorServer/http_server/data"
	kafka "github.com/segmentio/kafka-go"
)

func main(){
	
	// This interacts with the monitor_server via kafka server
	// Initialize the context
	ctx := context.Background()

	// Initialize the kafka reader for the respective topic

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{commons.BrokerAddress1, commons.BrokerAddress2},
		Topic: commons.ResponseTopic,
	})

	// Initiazlize the map
	newWebsitesList := []string{"www.google.com","www.facebook.com","xyz"}
	data.InitializeMap(newWebsitesList)

	// For synchronization 
	wg := new(sync.WaitGroup)

	// Read from the Kafka topic
	wg.Add(1)
		go data.ReadFromTopic(r, ctx)
	
	time.Sleep(11*time.Second)
	data.AccessAllData()
	wg.Wait()
}