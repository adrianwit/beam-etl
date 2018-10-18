package model

type DemandSupplyPerformance struct {
	DemandSideId      int
	SupplySideId      int
	EventOneTypeCount int
	EventTwoTypeCount int
	Charge            float64
	Payment           float64
}
