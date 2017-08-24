package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var cfg Configuration

func main() {

	// See : https://github.com/sirupsen/logrus
	formatter := &log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05.000",
		FullTimestamp:   true,
	}
	log.SetFormatter(formatter)

	// log.Info("Some info. Earth is not flat.")
	// log.Warning("This is a warning")
	// log.Error("Not fatal. An error. Won't stop execution")
	// log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	// log.Panic("Do not panic")

	var err error
	cfg, err = LoadConfig("config.json")
	if err != nil {
		log.Fatal("Had an error : ", err)
		panic(err)
	}
	log.Warning(cfg.RemoteToken)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	token, err := LoadToken("hipchat.json")
	if err != nil {
		log.Fatal("Had an error : ", err)
		panic(err)
	}

	// prometheus monitoring
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	// Add a remote API
	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/health/{token}", GetHealthEndpoint).Methods("GET")
		router.HandleFunc("/loglevel/{level}/{token}", SetLogLevelEndpoint).Methods("PUT")
		log.Fatal(http.ListenAndServe(":12345", router))
	}()

	log.Info(fmt.Sprintf("%s ready, ^C exits", cfg.BotName))

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
				log.Debug("received : ", string(msg))

				var f FeedbackEvent
				err = json.Unmarshal(msg, &f)
				if err != nil {
					log.Error("Had an error : ", err)
				}

				f = SetBackground(f)
				f = SetRoom(f, cfg)
				f = FormatMessage(f)
				err = SendToHipChat(f, cfg, token)
				if err != nil {
					log.Error("Had an error : ", err)
					panic(err)
				}

			}

		}
	}

}

// GetHealthEndpoint ... Remote access health check
func GetHealthEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	mytime := time.Now()

	w.Header().Set("Content-Type", "application/json")
	token := params["token"]
	var reply string
	if TestToken(token) {
		reply = fmt.Sprintf("{\"status\": \"good\", \"token\": \"%s\", \"uptime\": \"%s\"}", token, mytime)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		reply = ""
	}
	w.Write([]byte(reply))
	return
}

// SetLogLevelEndpoint ... Remote set of logging level
func SetLogLevelEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	w.Header().Set("Content-Type", "application/json")
	token := params["token"]
	level := params["level"]

	if TestToken(token) {
		http.Error(w, "StatusAccepted", http.StatusAccepted)
		switch level {
		case "Error":
			log.SetLevel(log.ErrorLevel)
		case "Warn":
			log.SetLevel(log.WarnLevel)
		case "Debug":
			log.SetLevel(log.DebugLevel)
		case "Info":
			log.SetLevel(log.InfoLevel)
		default:
			log.WithFields(log.Fields{
				"level": level,
			}).Info("Incorrect logging level")
			http.Error(w, "BadRequest", http.StatusBadRequest)
		}

	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

	}
	w.Write([]byte(""))
	return
}

// TestToken ..  test access token
func TestToken(token string) bool {
	if cfg.RemoteToken == token {
		return true
	}
	log.WithFields(log.Fields{
		"token": token,
	}).Warning("Incorrect remote access token")

	return false

}
