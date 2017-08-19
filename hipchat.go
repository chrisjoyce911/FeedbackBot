package main

import (
	"fmt"
	"log"

	"github.com/andybons/hipchat"
)

// SendToHipChat ... Post feedback to HipChat
func SendToHipChat(f FeedbackEvent, cfg Configuration, token Token) error {
	fmt.Printf("From : %s\nMessage : %s\nTo Hip : %s\nBackground : %s\n", f.Category, f.Comment, f.HipChat.Room, f.HipChat.Background)
	c := hipchat.Client{AuthToken: token.HipToken}

	req := hipchat.MessageRequest{
		RoomId:        f.HipChat.Room,
		From:          cfg.BotName,
		Message:       f.Comment,
		Color:         f.HipChat.Background,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}
	err := c.PostMessage(req)
	if err != nil {
		log.Printf("Expected no error, but got %q\n", err)
		return err
	}
	return err
}
