package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

type HipChatNotificationEvent struct {
	Event string
	Item  HipChatEventItem
}

type HipChatEventItem struct {
	Message HipChatEventMessage
	Room    HipChatRoom
}

type HipChatEventMessage struct {
	Color         string
	Id            string
	MessageFormat string
	From          string
	Message       string
}

type HipChatRoom struct {
	Id   string
	Name string
}

type SlackMessage struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

//https://api.slack.com/docs/formatting
func reformat(source string) string {
	linkRegex := regexp.MustCompile("<a href=\"([^\"]+)\">([^<]+)</a>")
	message := linkRegex.ReplaceAllString(source, "<$1|$2>")

	newlineRegex := regexp.MustCompile("<br[ ]?/>")
	message = newlineRegex.ReplaceAllString(message, "\n")

	return message
}

func sendToSlack(webhookUrl string, channel string, sourceMessage HipChatEventMessage) (*http.Response, error) {

	message := reformat(sourceMessage.Message)

	slackMessage := SlackMessage{Channel: channel, Username: sourceMessage.From, Text: message}

	payload, err := json.Marshal(slackMessage)
	if err != nil {
		return nil, err
	}
	payloadStr := string(payload)
	fmt.Println(payloadStr)
	return http.PostForm(webhookUrl, url.Values{"payload": {payloadStr}})
}

func handler(w http.ResponseWriter, r *http.Request, slackUrl string, slackChannel string) {
	var notification HipChatNotificationEvent

	json.NewDecoder(r.Body).Decode(&notification)

	_, err := sendToSlack(slackUrl, slackChannel, notification.Item.Message)
	if err != nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func main() {
	incomingWebhookUrl := os.Getenv("SLACK_URL")
	channel := os.Getenv("SLACK_CHANNEL")
	port := os.Getenv("PORT")
	if len(incomingWebhookUrl) == 0 || len(channel) == 0 || len(port) == 0 {
		fmt.Println("environment variables SLACK_URL, SLACK_CHANNEL, PORT are not set correcly")
		return
	}
	fmt.Println("Slack Incoming Webhook URL: " + incomingWebhookUrl)
	fmt.Println("Slack Channel: " + channel)
	fmt.Println("Port: " + port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, incomingWebhookUrl, channel)
	})
	http.ListenAndServe(":"+port, nil)
}
