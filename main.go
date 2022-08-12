package main

import (
	"context"
	"fmt"

	"github.com/monitorServer/commons"
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

	// Read from the kafka topic at given intervals and display the message
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("cannot read from the kafka queue")
			continue
		}
		fmt.Println("Message->  Offset : ", msg.Offset, " URL : ", string(msg.Key), "Status : ", string(msg.Value))
	}
}