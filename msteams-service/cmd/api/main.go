package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	gRpcPort = "50001"
	webhookURL ="https://outlook.office.com/webhook/YOUR_WEBHOOK_URL_OF_TEAMS_CHANNEL"
)

type Config struct {
	WebHook string
}

func main() {

	app := Config{
		WebHook : webhookURL,
	}

	log.Println("Starting service on port", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}

}