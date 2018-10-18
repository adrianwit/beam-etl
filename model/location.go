package model
const LocationDQL = "SELECT id, typeId, COALESCE(charge,0) AS charge,  COALESCE(revenueShare, 0) AS revenueShare FROM db1.locations"


type Location struct {
	ID           int     `bigquery:"id"`
	TypeID       int     `bigquery:"typeId"`
	Charge       float64 `bigquery:"charge"`
	RevenueShare float64     `bigquery:"revenueShare"`
}


