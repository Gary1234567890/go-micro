package models

type MSTeamsMessage struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	TargetURL string `json:"target_url"`
}