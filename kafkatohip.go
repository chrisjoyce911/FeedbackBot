package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	router := mux.NewRouter()

	router.HandleFunc("/health/{token}", healthEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))

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

func healthEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	mytime := time.Now()

	w.Header().Set("Content-Type", "application/json")
	token := params["token"]

	reply := fmt.Sprintf("{\"status\": \"good\", \"token\": \"%s\", \"uptime\": \"%s\"}", token, mytime)

	w.Write([]byte(reply))
	return
}
