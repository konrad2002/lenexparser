package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type MeetInfo struct {
	XMLName           xml.Name          `xml:"MEETINFO"`
	Approved          string            `xml:"approved,attr"`
	City              string            `xml:"city,attr"`
	Course            enums.Course      `xml:"course,attr"`
	Date              parser.CustomTime `xml:"date,attr"`
	Daytime           parser.CustomTime `xml:"daytime,attr"`
	Name              string            `xml:"name,attr"`
	Nation            enums.Nation      `xml:"nation,attr"`
	Pool              Pool              `xml:"POOL"`
	QualificationTime parser.SwimTime   `xml:"qualificationtime,attr"`
	State             string            `xml:"state,attr"`
	Timing            enums.Timing      `xml:"timing,attr"`
}
