package main

import (
	"fmt"
	"os"
	"os/signal"

	kafka "github.com/bsm/sarama-cluster"

	log "github.com/sirupsen/logrus"
)

//FeedBackConsumer ... collects feedback handing back to main to process
func FeedBackConsumer(cfg Configuration, messages chan []byte) {

	// init (custom) config, enable errors and notifications
	kafkaConfig := kafka.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Group.Return.Notifications = true

	formatter := &log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05.000",
		FullTimestamp:   true,
	}

	log.SetFormatter(formatter)
	log.Info("Starting consumer")

	// init consumer
	brokers := []string{cfg.Broker}
	topics := []string{cfg.WebTopic, cfg.AppTopic}
	consumer, err := kafka.NewConsumer(brokers, cfg.GroupID, topics, kafkaConfig)
	if err != nil {
		log.Fatal("Had an error : ", err)
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume messages, watch errors and notifications
	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key)

				messages <- msg.Value
				consumer.MarkOffset(msg, "")
			}
		case err, more := <-consumer.Errors():
			if more {
				log.Printf("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				log.Printf("Rebalanced: %+v\n", ntf)
			}
		case <-signals:
			messages <- []byte("consumer-quit")
			return
		}

	}

}
