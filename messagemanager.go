package main

import (
	"strings"

	"github.com/andybons/hipchat"
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
	HipChat           HipChat
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

// HipChat ... details for HipChat client
type HipChat struct {
	Room       string `json:"room"`
	Background string `json:"background"`
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

func forwardMessage(message string, inthis []Channel) (hipchatChannel string, background string) {
	background = hipchat.ColorGray

	for i := 0; i < len(inthis); i++ {
		hipchatChannel = inthis[i].HipChat
		for f := 0; f < len(inthis[i].RedirectRules); f++ {
			if strings.Contains(message, inthis[i].RedirectRules[f].ContainsText) {
				hipchatChannel = inthis[i].RedirectRules[f].HipChat
				for b := 0; b < len(inthis[i].RedirectRules[f].BackgroundRules); b++ {
					if strings.Contains(message, inthis[i].RedirectRules[f].BackgroundRules[b].ContainsText) {
						background = inthis[i].RedirectRules[f].BackgroundRules[b].Background
					}
				}
			}
		}
		return hipchatChannel, background
	}
	return "", background
}
