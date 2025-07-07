package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Event struct {
	XMLName          xml.Name          `xml:"EVENT"`
	AgeGroups        string            `xml:"AGEGROUPS"` // TODO
	Daytime          parser.CustomTime `xml:"daytime,attr"`
	EventId          int               `xml:"eventid,attr"`
	Fee              string            `xml:"FEE"` // TODO
	Gender           enums.EventGender `xml:"gender,attr"`
	Heats            string            `xml:"HEATS"` // TODO
	MaxEntries       int               `xml:"maxentries,attr"`
	Number           int               `xml:"number,attr"`
	Order            int               `xml:"order,attr"`
	PreEventId       int               `xml:"preveventid,attr"`
	Round            enums.Round       `xml:"round,attr"`
	Run              int               `xml:"run,attr"`
	Status           enums.EventStatus `xml:"status,attr"`
	SwimStyle        string            `xml:"SWIMSTYLE"`        // TODO
	TimeStandardRefs string            `xml:"TIMESTANDARDREFS"` // TODO
	Timing           enums.Timing      `xml:"timing,attr"`
	Type             enums.EventType   `xml:"type,attr"`
}
