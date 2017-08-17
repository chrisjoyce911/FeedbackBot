package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/andybons/hipchat"
	kafka "github.com/bsm/sarama-cluster"
)

func main() {

	cfg := loadcfg("config.json")

	fmt.Printf("%s ready, ^C exits\n", cfg.BotName)

	// this is just a placeholder for waht a message may looklike
	m := Message{
		ID:      1234,
		Type:    "mobile",
		Channel: "channel",
		Text:    "Text ",
	}

	// init (custom) config, enable errors and notifications
	kafkaConfig := kafka.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Group.Return.Notifications = true

	// init consumer
	brokers := []string{"127.0.0.1:9092"}
	topics := []string{"my_topic"}
	consumer, err := kafka.NewConsumer(brokers, "customer-feedback-tohipchat-group", topics, kafkaConfig)
	if err != nil {
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
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)

				// will some rework
				if m.Type == "message" {
					hip, background, forward := forwardMessage(m.Text, cfg.Channels)

					fmt.Printf("From : %s\nMessage : %s\nTo Hip : %s\nBackground : %s\n", m.Channel, m.Text, hip, background)
					if forward {
						c := hipchat.Client{AuthToken: cfg.HipToken}
						var HipRoomID = hip

						req := hipchat.MessageRequest{
							RoomId:        HipRoomID,
							From:          cfg.BotName,
							Message:       m.Text,
							Color:         background,
							MessageFormat: hipchat.FormatText,
							Notify:        true,
						}
						if err := c.PostMessage(req); err != nil {
							log.Fatalln("Expected no error, but got %q", err)
						}
					}
				}
				// mark message as processed
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
			return
		}
	}
}

func forwardMessage(message string, inthis []Channel) (hipchatChannel string, background string, forward bool) {
	background = hipchat.ColorGray

	for i := 0; i < len(inthis); i++ {
		hipchatChannel = inthis[i].HipChat
		for f := 0; f < len(inthis[i].RedirectRules); f++ {
			if strings.Contains(message, inthis[i].RedirectRules[f].ContainsText) {
				hipchatChannel = inthis[i].RedirectRules[f].HipChat
				for b := 0; b < len(inthis[i].RedirectRules[f].BackgroundRules); b++ {
					if strings.Contains(message, inthis[i].RedirectRules[f].BackgroundRules[b].ContainsText) {
						background = inthis[i].RedirectRules[f].BackgroundRules[b].Background
					}
				}
			}
		}
		return hipchatChannel, background, true
	}
	return "", background, false
}

func loadcfg(configfile string) (cfg Configuration) {
	cfg, err := LoadConfig(configfile)
	if err != nil {
		err = saveConfig(createMockConfig(), configfile)
		cfg = createMockConfig()
		if err != nil {
			panic(err)
		}
	}

	saveConfig(cfg, configfile)

	switch {
	case cfg.HipToken == "":
		log.Fatalln("Hipchat token is a required argument")
	}

	return cfg
}
