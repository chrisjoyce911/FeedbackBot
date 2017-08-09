package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/andybons/hipchat"
)

func main() {

	cfg := loadcfg("config.json")

	ws, id := slackConnect(cfg.SlackToken)

	fmt.Printf("%s ready, ^C exits\n", cfg.BotName)

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
				parts := strings.Fields(m.Text)
				switch {
				case len(parts) == 2 && parts[1] == "channel":
					go func(m Message) {
						m.Text = fmt.Sprintf("This is channel %v\n", m.Channel)
						postMessage(ws, m)
					}(m)
				default:
					m.Text = fmt.Sprintf("sorry, that does not compute\n")
					postMessage(ws, m)
				}
			}

			if m.Type == "message" {
				hip, background, forward := forwardMessage(m.Channel, m.Text, cfg.Channels)

				fmt.Printf("From : %s\nMessage : %s\nTo Hip : %s\nBackground : %s\n", m.Channel, m.Text, hip, background)
				if forward {
					c := hipchat.Client{AuthToken: cfg.HipToken}
					var HipRoomID = hip

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

func forwardMessage(slackChannel string, message string, inthis []Channel) (hipchatChannel string, background string, forward bool) {
	background = hipchat.ColorGray

	for i := 0; i < len(inthis); i++ {
		if inthis[i].Slack == slackChannel {
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
			return hipchatChannel, background, true
		}
	}
	return "", background, false
}

func loadcfg(configfile string) (cfg Configuration) {
	cfg, err := LoadConfig(configfile)
	if err != nil {
		err = saveConfig(createMockConfig(), configfile)
		cfg = createMockConfig()
		if err != nil {
			panic(err)
		}
	}

	saveConfig(cfg, configfile)

	switch {
	case cfg.SlackToken == "":
		log.Fatalln("Slack token is a required argument")
	case cfg.HipToken == "":
		log.Fatalln("Hipchat token is a required argument")
	}

	return cfg
}
