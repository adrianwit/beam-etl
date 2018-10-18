package model

import "time"

const EventDQL = "SELECT id, sessionId, eventTypeId, demandSideId, subjectId, supplySideSideId, locationId, charge, payment, timestamp FROM db1.events"

type Event struct {
	ID               int       `bigquery:"id"`
	SessionID        string    `bigquery:"sessionId"`
	EventTypeId      int       `bigquery:"eventTypeId"`
	DemandSideId     int       `bigquery:"demandSideId"`
	SubjectId        int       `bigquery:"subjectId"`
	SupplySideSideId int       `bigquery:"supplySideSideId"`
	LocationID       int       `bigquery:"locationId"`
	Charge           float64   `bigquery:"charge"`
	Payment          float64   `bigquery:"payment"`
	Timestamp        time.Time `bigquery:"timestamp"`
	Location         Location
}

