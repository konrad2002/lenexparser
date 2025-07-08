package enums

type RecordStatus string

const (
	RecordStatusApproved        RecordStatus = "APPROVED"
	RecordStatusPending         RecordStatus = "PENDING"
	RecordStatusInvalid         RecordStatus = "INVALID"
	RecordStatusApprovedHistory RecordStatus = "APPROVED.HISTORY"
	RecordStatusPendingHistory  RecordStatus = "PENDING.HISTORY"
	RecordStatusTargetTime      RecordStatus = "TARGETTIME"
)
