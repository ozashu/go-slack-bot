package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
	"time"
)

func main() {
	// Load Env variables from .dot file
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")

	// Create a new client to slack by giving token
	// Set debug to true while developing
	client := slack.New(token, slack.OptionDebug(true))
	// Create the Slack attachment that we will send to the channel
	attachment := slack.Attachment{
		Pretext: "Suber Bot Message",
		Text: "some text",
		// Color Styles the Text, making it possible to heva like Warnings etc.
		Color: "#36a64f",
		// Fields are Optional extra data!
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}
	// PostMessage will send the message away.
	// First parameter is just the channelID, makes no sense to accept it
	_, timestamp, err := client.PostMessage(
		channelID,
		// uncomment the item below to add a extra Header to the message, try it out:)
		// slack.MsgOptionText("New message from bot", false),
		slack.MsgOptionAttachments(attachment),
		)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)
}