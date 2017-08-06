package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/andybons/hipchat"
)

func main() {

	cfg, err := LoadConfig("config.json")
	if err != nil {
		err = saveConfig(createMockConfig(), "config.json")
		cfg = createMockConfig()
		if err != nil {
			panic(err)
		}
	}

	slackPtr := flag.String("s", cfg.SlackToken, "Slack token")
	hipPtr := flag.String("h", cfg.HipToken, "HipChat token")
	channelPtr := flag.String("c", cfg.SlackChannel, "Slack channel")
	roomMobPtr := flag.String("m", cfg.MobHipRoom, "Mobile HipChat room")
	roomWebPtr := flag.String("w", cfg.WebHipRoom, "Web HipChat room")

	flag.Parse()

	cfg.SlackToken = *slackPtr
	cfg.HipToken = *hipPtr
	cfg.MobHipRoom = *roomMobPtr
	cfg.WebHipRoom = *roomWebPtr
	cfg.SlackChannel = *channelPtr

	saveConfig(cfg, "config.json")

	switch {
	case *slackPtr == "":
		log.Fatalln("Slack token is a required argument")
	case *hipPtr == "":
		log.Fatalln("Hipchat token is a required argument")
	}

	for i := 0; i < len(cfg.Channels); i++ {
		fmt.Println(cfg.Channels[i].Slack)
	}

	log.Fatalln("Stop here")
	ws, id := slackConnect(cfg.SlackToken)

	fmt.Printf("%s ready, ^C exits", cfg.BotName)

	go func() {
		c := time.Tick(time.Duration(cfg.SlackRepTime) * time.Second)
		for now := range c {
			rmsg := Message{
				Type:    "message",
				Channel: cfg.SlackReport,
				Text:    fmt.Sprintf("I'm alive %v\n", now),
			}
			postMessage(ws, rmsg)
		}

		return
	}()

	for {
		m, err := getMessage(ws)
		if err != nil {
			fmt.Println(err)
		} else {

			if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
				// if so try to parse if
				parts := strings.Fields(m.Text)
				switch {
				case len(parts) == 2 && parts[1] == "channel":
					go func(m Message) {
						m.Text = fmt.Sprintf("This is channel %v\n", m.Channel)
						postMessage(ws, m)
					}(m)
				default:
					// huh?
					m.Text = fmt.Sprintf("sorry, that does not compute\n")
					postMessage(ws, m)
				}
			}

			if m.Type == "message" {
				fmt.Printf("%s %s\n", m.Channel, m.Text)
				if m.Channel == cfg.SlackChannel {
					c := hipchat.Client{AuthToken: cfg.HipToken}

					var background = "gray"

					switch {
					case strings.Contains(m.Text, "Rating: Satisfied"):
						background = hipchat.ColorGreen
					case strings.Contains(m.Text, "Rating: Neutral"):
						background = hipchat.ColorYellow
					case strings.Contains(m.Text, "Rating: Not Satisfied"):
						background = hipchat.ColorRed
					}

					var HipRoomID = cfg.MobHipRoom
					if strings.Contains(m.Text, "OzLotteries for Web") {
						HipRoomID = cfg.WebHipRoom
					}

					req := hipchat.MessageRequest{
						RoomId:        HipRoomID,
						From:          cfg.BotName,
						Message:       m.Text,
						Color:         background,
						MessageFormat: hipchat.FormatText,
						Notify:        true,
					}
					if err := c.PostMessage(req); err != nil {
						log.Fatalln("Expected no error, but got %q", err)
					}
				}
			}
		}
	}
}

// Sum .. for testing
func Sum(a, b int) int {
	return a + b
}
