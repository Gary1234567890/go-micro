package main

import (
	"log"
	"msteams-service/models"
	"net/http"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
)

func (app *Config) SendMessage(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload models.MSTeamsMessage
	_ = app.readJSON(w, r, &requestPayload)

	err := app.SendMessageToTeams(requestPayload)

	var resp = jsonResponse{}

	if err != nil {
		resp.Error = false
		resp.Message = "Sent"
	}	else {
		resp.Error = true
		resp.Message = err.Error()
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}

func(app *Config) SendMessageToTeams(incomingMessage models.MSTeamsMessage) error {

	// Initialize a new Microsoft Teams client.
	mstClient := goteamsnotify.NewClient()

	// Set webhook url.
	webhookUrl := incomingMessage.TargetURL

	// Setup message card.
	msgCard := goteamsnotify.MessageCard{}
	msgCard.Title = incomingMessage.Title
	msgCard.Text = incomingMessage.Body
	msgCard.ThemeColor = "#DF813D"

	// Send the message with default timeout/retry settings.
	err := mstClient.Send(webhookUrl, msgCard)
	if err != nil {
		log.Printf("failed to send message: %v", err)
		return err
	}

	return nil
}