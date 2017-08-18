package main

import (
	"encoding/json"
	"io/ioutil"
)

// FeedbackEvent ... Holds a customer feedback
type FeedbackEvent struct {
	EventType         string `json:"event_type"`
	Name              string `json:"name"`
	Category          string `json:"category"`
	Rating            string `json:"rating"`
	Comment           string `json:"comment"`
	Source            string `json:"source"`
	ClientInformation ClientInformation
	ClientPlatform    ClientPlatform
	ClientDevice      ClientDevice
	ClientEnvironment ClientEnvironment
}

// ClientInformation ... About the device the feedback was sourced on
type ClientInformation struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Build   string `json:"build"`
}

// ClientPlatform ... About the version of app
type ClientPlatform struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ClientDevice ... About the device the feedback was sourced on
type ClientDevice struct {
	Class        string `json:"class"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
}

// ClientEnvironment ... what timezone is the clinet in ?
type ClientEnvironment struct {
	Timezone string `json:"timezone"`
	Locale   string `json:"locale"`
}

/*
{
   "event_type":"feedback_event",
   "name":"feedback event",
   "category":"mobile feedback",
   "client_information":{
      "name":"Ozlotteries for Android",
      "version":"4.9.4",
      "build":"release",
      "client_platform":{"name":"Android","version":"7.1.1"},
      "client_device":{"class":"phone","manufacturer":"Samsung","model":"Galaxy S8"},
      "client_environment":{"timezone":"Australia/Brisbane","locale":"en_AU"}},
   "rating":"' . $rating . '",
   "comment":"App is biased and does not draw my numbers. I can never win!!",
   "source":"feedback"
}
*/

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
