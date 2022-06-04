package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type InfluxPayload struct {
	Bucket string `json:"bucket"`
	Measurement string	`json:"measurement"`
	Tag string `json:"tag"`
	FieldUnit string `json:"field_unit"`
	Value int `json:"value"`
	CreatedAt int64 `json:"created_at"` 
}

func (app *Config) SendData(w http.ResponseWriter, r *http.Request) {
	log.Println("got something to log.")
	var b bytes.Buffer
	b.ReadFrom(r.Body)
	log.Println(string(b.String()))

	// read json into var
	var requestPayload Payload
 	err := json.Unmarshal(b.Bytes(),&requestPayload)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Println("Error decoding object: ", err.Error())
		return
    }

	log.Println(requestPayload.Data)

	req := InfluxPayload{}
	json.Unmarshal([]byte(requestPayload.Data),&req)

	// insert data
	err = app.WriteData(req)

	resp := jsonResponse{}

	if err != nil{
		log.Println("Error logging data: " + err.Error())
		resp.Error = true
		resp.Message = "Error logging data: " + err.Error()
	} else{
		resp.Error = false
		resp.Message = "logged"
		log.Println("Sent to Rabbit: ", requestPayload)
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}