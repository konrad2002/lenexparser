package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type RelayPosition struct {
	XMLName      xml.Name                  `xml:"RELAYPOSITION"`
	Athlete      Athlete                   `xml:"ATHLETE"`
	AthleteId    int                       `xml:"athleteid,attr"`
	MeetInfo     MeetInfo                  `xml:"MEETINFO"`
	Number       int                       `xml:"number,attr"`
	ReactionTime string                    `xml:"reactiontime,attr"` // TODO data type
	Status       enums.RelayPositionStatus `xml:"status,attr"`
}
