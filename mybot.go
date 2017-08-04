package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/andybons/hipchat"
)

func main() {

	messagecount := 0
	ws, id := slackStart("FIXME")
	fmt.Println("mybot ready, ^C exits")

	for {
		m, err := getMessage(ws)
		if err != nil {
			fmt.Println(err)
			messagecount = 0
		} else {

			if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
				// if so try to parse if
				parts := strings.Fields(m.Text)
				if len(parts) == 2 && parts[1] == "channel" {
					// looks good, get the quote and reply with the result

					go func(m Message) {
						m.Text = fmt.Sprintf("This is channel %v\n", m.Channel)
						// postMessage(ws, m)
					}(m)

					// NOTE: the Message object is copied, this is intentional
				} else {
					// huh?
					m.Text = fmt.Sprintf("sorry, that does not compute\n")
					postMessage(ws, m)
				}
			}

			if m.Type == "message" {
				fmt.Printf("%s %s %s\n", m.Tstamp, m.Channel, m.Text)

				if m.Channel == "FIXME" {
					c := hipchat.Client{AuthToken: "FIXME"}

					req := hipchat.MessageRequest{
						RoomId:        "Feedback Room",
						From:          "Jonny Boom",
						Message:       fmt.Sprintf("Message forwarded Slack : %s", m.Text),
						Color:         hipchat.ColorPurple,
						MessageFormat: hipchat.FormatText,
						Notify:        true,
					}
					if err := c.PostMessage(req); err != nil {
						log.Printf("Expected no error, but got %q", err)
					}
				}
				messagecount++
			}
		}

	}
}
