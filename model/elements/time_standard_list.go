package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type TimeStandardList struct {
	XMLName            xml.Name                   `xml:"TIMESTANDARDLIST"`
	AgeGroup           AgeGroup                   `xml:"AGEGROUP"`
	Course             enums.Course               `xml:"course,attr"`
	Gender             enums.EventGender          `xml:"gender,attr"`
	Handicap           enums.HandicapClass        `xml:"handicap,attr"`
	Name               string                     `xml:"name,attr"`
	TimeStandardListId int                        `xml:"timestandardlistid,attr"`
	TimeStandards      []TimeStandard             `xml:"TIMESTANDARDS>TIMESTANDARD"`
	Type               enums.TimeStandardListType `xml:"type,attr"`
}
