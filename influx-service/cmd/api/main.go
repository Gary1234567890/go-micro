package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type Config struct {
	APIPort string
	InfluxDb string
	AdminUser string
	Password string
	InfluxHost string
	InfluxClient influxdb2.Client
}

var app Config

func main() {

	app := Config{
		AdminUser: os.Getenv("INFLUXDB_ADMIN_USER"),
		Password: os.Getenv("INFLUXDB_ADMIN_PASSWORD"),
		InfluxHost: "influxdb",
		APIPort: ":80",
		InfluxDb: "DB0",
	}

	// create new client with default option for server url authenticate by token
	app.InfluxClient = influxdb2.NewClient("http://" + app.InfluxHost + ":8086", app.AdminUser + ":" + app.Password)
	
	// start web server
	// go app.serve()
	log.Println("Starting service on port", app.APIPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s", app.APIPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	// Ensures background processes finish
	app.InfluxClient.Close()
}