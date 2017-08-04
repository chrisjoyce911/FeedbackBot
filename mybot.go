package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/andybons/hipchat"
)

func main() {

	slackPtr := flag.String("s", "", "Slack token")
	hipPtr := flag.String("h", "", "HipChat token")
	channelPtr := flag.String("c", "Error Logs", "Slack channel")
	roomPtr := flag.String("r", "Integration Testing", "HipChat room")

	flag.Parse()

	var SlackToken = *slackPtr
	var HipToken = *hipPtr
	var HipRoomID = *roomPtr
	var SlackChannel = *channelPtr

	ws, id := slackConnect(SlackToken)

	fmt.Println("mybot ready, ^C exits")

	for {
		m, err := getMessage(ws)
		if err != nil {
			fmt.Println(err)
		} else {

			if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
				// if so try to parse if
				parts := strings.Fields(m.Text)
				if len(parts) == 2 && parts[1] == "channel" {
					go func(m Message) {
						m.Text = fmt.Sprintf("This is channel %v\n", m.Channel)
					}(m)
				} else {
					// huh?
					m.Text = fmt.Sprintf("sorry, that does not compute\n")
					postMessage(ws, m)
				}
			}

			if m.Type == "message" {
				fmt.Printf("%s %s\n", m.Channel, m.Text)
				if m.Channel == SlackChannel {
					c := hipchat.Client{AuthToken: HipToken}

					req := hipchat.MessageRequest{
						RoomId:        HipRoomID,
						From:          "Jonny Boom",
						Message:       m.Text,
						Color:         hipchat.ColorPurple,
						MessageFormat: hipchat.FormatText,
						Notify:        true,
					}
					if err := c.PostMessage(req); err != nil {
						log.Printf("Expected no error, but got %q", err)
					}
				}
			}
		}

	}
}
