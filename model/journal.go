package model

import "time"

type Journal struct {
	Timestamp time.Time
	Level     string
	Message   string
}


