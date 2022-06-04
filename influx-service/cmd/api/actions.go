package main

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func(app *Config) WriteData(data InfluxPayload) error {
	// user blocking write client for writes to desired bucket
	writeAPI := app.InfluxClient.WriteAPIBlocking("", data.Bucket)

	// create point using fluent style
	p := influxdb2.NewPointWithMeasurement(data.Measurement).
		AddTag("unit", data.Tag).
		AddField(data.FieldUnit, data.Value).
		SetTime(time.Now())
	err := writeAPI.WritePoint(context.Background(), p)
	return err
}