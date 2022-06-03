package main

import "time"

type Models struct {
	LogEntry InfuxData
}

type InfuxData struct {
	Bucket string 
	Measurement string	
	Tag string
	FieldUnit string
	Value int 
	CreatedAt time.Time 
}

