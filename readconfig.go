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
}

func createMockConfig() Configuration {
	return Configuration{
		BotName:      "Slack to HipCat",
		SlackToken:   "",
		HipToken:     "",
		SlackReport:  "",
		SlackRepTime: 360,
		SlackChannel: "",
		MobHipRoom:   "",
		WebHipRoom:   "",
	}
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
