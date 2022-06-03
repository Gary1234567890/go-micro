package main

import "net/http"

type JSONPayload struct {
	Name string `json:"name"`
	Data InfuxData `json:"data"`

}

func (app *Config) SendData(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	// insert data
	err := app.WriteData(requestPayload.Data)

	resp := jsonResponse{}

	if err != nil{
		resp.Error = true
		resp.Message = "Error logging data: " + err.Error()
	} else{
		resp.Error = false
		resp.Message = "logged"
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}