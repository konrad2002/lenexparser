package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Entry struct {
	XMLName       xml.Name          `xml:"ENTRY"`
	AgeGroupId    string            `xml:"agegroupid,attr"`
	EntryCourse   enums.Course      `xml:"entrycourse,attr"`
	EntryDistance int               `xml:"entrydistance,attr"`
	EntryTime     parser.SwimTime   `xml:"entrytime,attr"`
	EventId       int               `xml:"eventid,attr"`
	Handicap      string            `xml:"handicap,attr"` // TODO: enum
	HeatId        int               `xml:"heatid,attr"`
	Lane          int               `xml:"lane,attr"`
	MeetInfo      string            `xml:"MEETINFO"`                      // TODO
	ReplayOptions string            `xml:"RELAYPOSITIONS>REPLAYPOSITION"` // TODO
	Status        enums.EntryStatus `xml:"status,attr"`
}
