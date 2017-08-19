package main

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration ... not sure how this will work yet
type Configuration struct {
	BotName  string `json:"botbame"`
	Broker   string `json:"broker"`
	Topic    string `json:"topic"`
	GroupID  string `json:"groupid"`
	Channels []Channel
}

// Token  ... As the HipChat token will be sorted in encoded josn we handle it in a different way
type Token struct {
	HipToken string `json:"hiptoken"`
}

//Channel ... source and desternation channel
type Channel struct {
	HipChat       string `json:"hipchat"`
	RedirectRules []RedirectRules
}

//RedirectRules ... if the text matches redirtect
type RedirectRules struct {
	HipChat         string `json:"hipchat"`
	ContainsText    string `json:"containstext"`
	BackgroundRules []BackgroundRules
}

//BackgroundRules ... if the text matches redirtect
type BackgroundRules struct {
	Background   string `json:"background"`
	ContainsText string `json:"containstext"`
}

//createMockConfig ... will generate a default config
func createMockConfig() Configuration {
	cfg := Configuration{
		BotName: "Kafka to HipCat",
		Broker:  "127.0.0.1:9092",
		Topic:   "my_topic",
		GroupID: "feedback_to_hipchat",
		Channels: []Channel{
			{
				HipChat: "Dev Test Channel",
				RedirectRules: []RedirectRules{
					{
						HipChat:      "Match oneChannel",
						ContainsText: "Match oneText",
						BackgroundRules: []BackgroundRules{
							{
								Background:   "green",
								ContainsText: "Rating: Satisfied"},
							{
								Background:   "yellow",
								ContainsText: "Rating: Neutral"},
							{
								Background:   "red",
								ContainsText: "Rating: Not Satisfied"},
						},
					},
					{
						HipChat:      "Match twoChannel",
						ContainsText: "Match twoText"},
				},
			},
			{
				HipChat: "Integration Testing"},
		},
	}
	return cfg
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
func saveToken(c Token, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}
