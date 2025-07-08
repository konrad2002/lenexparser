package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Record struct {
	XMLName   xml.Name           `xml:"RECORD"`
	Athlete   Athlete            `xml:"ATHLETE"`
	Comment   string             `xml:"comment,attr"`
	MeetInfo  MeetInfo           `xml:"MEETINFO"`
	Relay     Relay              `xml:"RELAY"`
	Splits    []Split            `xml:"SPLITS>SPLIT"`
	SwimStyle SwimStyle          `xml:"SWIMSTYLE"`
	SwimTime  parser.SwimTime    `xml:"swimtime,attr"`
	Status    enums.RecordStatus `xml:"status,attr"`
}
