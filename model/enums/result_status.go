package enums

type ResultStatus string

const (
	ResultStatusEXH  ResultStatus = "EXH"
	ResultStatusDSQ  ResultStatus = "DSQ"
	ResultStatusDNS  ResultStatus = "DNS"
	ResultStatusDNF  ResultStatus = "DNF"
	ResultStatusSICK ResultStatus = "SICK"
	ResultStatusWDR  ResultStatus = "WDR"
)
