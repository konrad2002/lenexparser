package enums

type HeatStatus string

const (
	HeatStatusScheduled  HeatStatus = "SCHEDULED"
	HeatStatusSeeded     HeatStatus = "SEEDED"
	HeatStatusRunning    HeatStatus = "RUNNING"
	HeatStatusUnofficial HeatStatus = "UNOFFICIAL"
	HeatStatusOfficial   HeatStatus = "OFFICIAL"
)
