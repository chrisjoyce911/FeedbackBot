package main

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration ... not sure how this will work yet
type Configuration struct {
	BotName     string `json:"botbame"`
	Broker      string `json:"broker"`
	AppTopic    string `json:"apptopic"`
	WebTopic    string `json:"webtopic"`
	GroupID     string `json:"groupid"`
	ReleseApp   string `json:"release_app"`
	ReleseWeb   string `json:"release_web"`
	Development string `json:"development"`
	RemoteToken string `json:"remoretoken"`
}

// Token  ... As the HipChat token will be sorted in encoded josn we handle it in a different way
type Token struct {
	HipToken string `json:"hiptoken"`
}

//createMockConfig ... will generate a default config
func createMockConfig() Configuration {
	cfg := Configuration{
		BotName:     "Kafka to HipCat",
		Broker:      "127.0.0.1:9092",
		AppTopic:    "my_topic",
		WebTopic:    "my_topic",
		GroupID:     "feedback_to_hipchat",
		ReleseApp:   "Integration Testing",
		ReleseWeb:   "Integration Testing",
		Development: "Integration Testing",
		RemoteToken: "RemoteToken",
	}
	return cfg
}

//createMockToken ... will generate a default config
func createMockToken() Token {
	t := Token{
		HipToken: "HIP_TOKEN",
	}
	return t
}

//LoadConfig reads config
func LoadConfig(configfilename string) (Configuration, error) {

	bytes, err := ioutil.ReadFile(configfilename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}

//LoadToken .. Loads HipChat token
func LoadToken(tokenfilename string) (Token, error) {
	bytes, err := ioutil.ReadFile(tokenfilename)
	if err != nil {
		t := createMockToken()
		terr := saveToken(t, tokenfilename)
		if terr != nil {
			return Token{}, terr
		}
		return Token{}, err
	}

	var t Token
	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return Token{}, err
	}

	return t, nil
}

//saveConfig ... saves config to file
func saveConfig(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

//saveToken ... saves hipchat token to file
func saveToken(t Token, filename string) error {
	bytes, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}
