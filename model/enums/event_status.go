package enums

type EventStatus string

const (
	EventStatusEntries    EventStatus = "ENTRIES"
	EventStatusSeeded     EventStatus = "SEEDED"
	EventStatusRunning    EventStatus = "RUNNING"
	EventStatusUnOfficial EventStatus = "UNOFFICIAL"
	EventStatusOfficial   EventStatus = "OFFICIAL"
)
