package enums

type QualifyConversion string

const (
	QualifyConversionNone              QualifyConversion = "NONE"
	QualifyConversionFinaPoints        QualifyConversion = "FINA_POINTS"
	QualifyConversionPercentLinear     QualifyConversion = "PERCENT_LINEAR"
	QualifyConversionNonConformingLast QualifyConversion = "NON_CONFORMING_LAST"
)
