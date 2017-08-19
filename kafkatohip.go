package main

import (
	"fmt"
	"time"
)

func main() {

	cfg, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	// token, err := LoadToken("hipchat.json")
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("%s ready, ^C exits\n", cfg.BotName)

	// this is just a placeholder for what a message may look like
	// f := FeedbackEvent{
	// 	Category: "mobile feedback",
	// 	Comment:  "Text ",
	// }

	messages := make(chan string)

	go FeedBackConsumer(cfg, messages)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-messages:
			fmt.Println("received", msg)
			time.Sleep(time.Second * 1)
		}
	}

}
