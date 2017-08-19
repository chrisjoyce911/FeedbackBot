package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

func main() {

	config := sarama.NewConfig()
	partition := int32(-1)
	topic := "customer_feedback"

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	config.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err := sarama.NewSyncProducer(strings.Split("localhost:9092", ","), config)
	if err != nil {
		fmt.Printf("Failed to open Kafka producer: %s\n", err)
	}
	//Some house keeping
	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Printf("Failed to close Kafka producer cleanly: %v\n", err)
		}
	}()

	filebites, err := ioutil.ReadFile("../test_configs/test_feedback.json")
	if err != nil {
		log.Fatalln(err)
	}
	jsonString := string(filebites)

	message := &sarama.ProducerMessage{Topic: topic, Partition: partition}
	message.Value = sarama.StringEncoder(jsonString)
	fmt.Println(jsonString)

	_, _, err = producer.SendMessage(message)
	if err != nil {
		fmt.Printf("Failed to produce message: %s", err)
	}

}
