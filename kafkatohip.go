package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {

	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalln("Had an error : ", err)
		panic(err)
	}

	token, err := LoadToken("hipchat.json")
	if err != nil {
		log.Fatalln("Had an error : ", err)
		panic(err)
	}

	fmt.Printf("%s ready, ^C exits\n", cfg.BotName)

	messages := make(chan []byte)

	go FeedBackConsumer(cfg, messages)

	consumerOK := true
	for consumerOK == true {
		select {
		case msg := <-messages:
			if string(msg) == "consumer-quit" {
				time.Sleep(500 * time.Millisecond)
				consumerOK = false
			} else {
				// fmt.Println("received : ", string(msg))

				var f FeedbackEvent
				err = json.Unmarshal(msg, &f)
				if err != nil {
					log.Fatalln("Had an error : ", err)
				}

				f = SetBackground(f)
				f = SetRoom(f, cfg)
				f = FormatMessage(f)
				err = SendToHipChat(f, cfg, token)
				if err != nil {
					log.Fatalln("Had an error : ", err)
					panic(err)
				}

			}

		}
	}

}
