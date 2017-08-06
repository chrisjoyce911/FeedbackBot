package main

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration ... not sure how this will work yet
type Configuration struct {
	BotName      string `json:"botbame"`
	SlackToken   string `json:"slacktoken"`
	HipToken     string `json:"hiptoken"`
	SlackReport  string `json:"slackreport"`
	SlackRepTime int    `json:"slackreptime"`
	SlackChannel string `json:"slackroom"`
	MobHipRoom   string `json:"mobilehipchat"`
	WebHipRoom   string `json:"webhipchat"`
	Channels     []Channel
}

//Channel ... source and desternation channel
type Channel struct {
	Slack   string `json:"slack"`
	HipChat string `json:"hipchat"`
}

func createMockConfig() Configuration {
	cfg := Configuration{
		BotName:      "Slack to HipCat",
		SlackToken:   "",
		HipToken:     "",
		SlackReport:  "",
		SlackRepTime: 600,
		SlackChannel: "",
		MobHipRoom:   "",
		WebHipRoom:   "",
		Channels: []Channel{
			{
				Slack:   "C5ADCS9FV",
				HipChat: "Dev Test Channel"},
			{
				Slack:   "G6J2D2HCY",
				HipChat: "Integration Testing"},
		},
	}
	return cfg
}

//LoadConfig reads config
func LoadConfig(filename string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
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

func saveConfig(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}
