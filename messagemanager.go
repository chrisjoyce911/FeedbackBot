package main

import (
	"fmt"
)

// FeedbackEvent ... Holds a customer feedback
type FeedbackEvent struct {
	EventType         string            `json:"event_type"`
	Name              string            `json:"name"`
	CustomerID        string            `json:"customerid"`
	Category          string            `json:"category"`
	Rating            string            `json:"rating"`
	Comment           string            `json:"comment"`
	Source            string            `json:"source"`
	ClientInformation ClientInformation `json:"client_information"`
	HipChat           HipChat
}

// ClientInformation ... About the device the feedback was sourced on
type ClientInformation struct {
	Name              string            `json:"name"`
	Version           string            `json:"version"`
	Build             string            `json:"build"`
	ClientPlatform    ClientPlatform    `json:"client_platform"`
	ClientDevice      ClientDevice      `json:"client_device"`
	ClientEnvironment ClientEnvironment `json:"client_environment"`
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
	Room        string `json:"room"`
	Background  string `json:"background"`
	FormatedMsg string `json:"formated_message"`
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
   "rating":"Satisfied",
   "comment":"App is biased and does not draw my numbers. I can never win!!",
   "source":"feedback"
}
*/

// SetBackground ... Set background based on the rating
func SetBackground(f FeedbackEvent) FeedbackEvent {
	switch f.Rating {
	case "Not Satisfied":
		f.HipChat.Background = "red"
	case "Neutral":
		f.HipChat.Background = "yellow"
	case "Satisfied":
		f.HipChat.Background = "green"
	default:
		f.HipChat.Background = "gray"
	}
	return f
}

// SetRoom ... What room do we display in
func SetRoom(f FeedbackEvent, cfg Configuration) FeedbackEvent {
	switch f.ClientInformation.Build {
	case "release":
		switch f.Category {
		case "mobile feedback":
			f.HipChat.Room = cfg.ReleseApp
		default:
			f.HipChat.Room = cfg.ReleseWeb
		}
	default:
		f.HipChat.Room = cfg.Development
	}
	return f
}

// FormatMessage ... Make a nice layout
func FormatMessage(f FeedbackEvent) FeedbackEvent {

	thismsg := fmt.Sprintf("%s\tCustomerID : %s\t Source : %s\n%s\n", f.Rating, f.CustomerID, f.Source, f.Comment)
	thismsg = fmt.Sprintf("%s\n", thismsg)

	f.HipChat.FormatedMsg = thismsg
	return f
}
