package enums

type EventStatus string

const (
	EventStatusEntries    EventStatus = "ENTRIES"
	EventStatusSeeded     EventStatus = "SEEDED"
	EventStatusRunning    EventStatus = "RUNNING"
	EventStatusUnofficial EventStatus = "UNOFFICIAL"
	EventStatusOfficial   EventStatus = "OFFICIAL"
)
