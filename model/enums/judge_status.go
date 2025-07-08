package enums

type JudgeStatus string

const (
	JudgeStatusAvailable JudgeStatus = "AVAILABLE"
	JudgeStatusScheduled JudgeStatus = "SCHEDULED"
	JudgeStatusCompleted JudgeStatus = "COMPLETED"
)
