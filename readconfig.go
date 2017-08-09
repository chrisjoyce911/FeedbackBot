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
	Channels     []Channel
}

//Channel ... source and desternation channel
type Channel struct {
	Slack         string `json:"slack"`
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
		BotName:      "Slack to HipCat",
		SlackToken:   "SLACK_TOKEN",
		HipToken:     "HIP_TOKEN",
		SlackReport:  "SLACK_KEEPALICE_CHANNEL",
		SlackRepTime: 600,
		Channels: []Channel{
			{
				Slack:   "SLACK0101",
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
				Slack:   "SLACK0123",
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

//saveConfig ... saves config to file
func saveConfig(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

//getConfig ... load config to file
func getConfig(filename string) Configuration {
	cfg, err := LoadConfig(filename)
	if err != nil {
		err = saveConfig(createMockConfig(), filename)
		cfg = createMockConfig()
		if err != nil {
			panic(err)
		}
	}
	return cfg
}
