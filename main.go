package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	esClient "github.com/yudhapratama10/search-consumer/infrastructure/elasticsearch"
	kafka "github.com/yudhapratama10/search-consumer/infrastructure/kafka"
	"github.com/yudhapratama10/search-consumer/model"
)

var elasticClient *elasticsearch.Client
var err error

func main() {
	elasticClient, err = esClient.GetClient()
	if err != nil {
		log.Fatal("error on elasticsearch client:", err)
	}

	kafkaClient := kafka.GetClient()

	log.Println("Starting search-consumer service")

	for {
		d, err := kafkaClient.ReadMessage(context.Background()) // ReadMessage blocks until recieve new message
		defer kafkaClient.Close()

		if err != nil {
			log.Printf("%T %+v", err, err)
			continue // skip below execution and proceed to new loop
		}

		messages := model.Message{}
		err = json.Unmarshal(d.Value, &messages)
		if err != nil {
			continue
		}

		fmt.Println(messages.Data.Name)

		consume(messages)
	}
}
