package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Entry struct {
	XMLName       xml.Name            `xml:"ENTRY"`
	AgeGroupId    string              `xml:"agegroupid,attr"`
	EntryCourse   enums.Course        `xml:"entrycourse,attr"`
	EntryDistance int                 `xml:"entrydistance,attr"`
	EntryTime     parser.SwimTime     `xml:"entrytime,attr"`
	EventId       int                 `xml:"eventid,attr"`
	Handicap      enums.HandicapClass `xml:"handicap,attr"`
	HeatId        int                 `xml:"heatid,attr"`
	Lane          int                 `xml:"lane,attr"`
	MeetInfo      MeetInfo            `xml:"MEETINFO"`
	ReplayOptions string              `xml:"RELAYPOSITIONS>REPLAYPOSITION"` // TODO
	Status        enums.EntryStatus   `xml:"status,attr"`
}
