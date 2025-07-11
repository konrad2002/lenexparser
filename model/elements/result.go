package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Result struct {
	XMLName        xml.Name            `xml:"RESULT"`
	Comment        string              `xml:"comment,attr"`
	EventId        int                 `xml:"eventid,attr"`
	Handicap       enums.HandicapClass `xml:"handicap,attr"`
	HeatId         int                 `xml:"heatid,attr"`
	Lane           int                 `xml:"lane,attr"`
	Points         int                 `xml:"points,attr"`
	ReactionTime   string              `xml:"reactiontime,attr"` // TODO data type
	RelayPositions []RelayPosition     `xml:"RELAYPOSITIONS>RELAYPOSITION"`
	ResultId       int                 `xml:"resultid,attr"`
	Status         enums.ResultStatus  `xml:"status,attr"`
	Splits         []Split             `xml:"SPLITS>SPLIT"`
	SwimDistance   int                 `xml:"swimdistance,attr"`
	SwimTime       parser.SwimTime     `xml:"swimtime,attr"`
}
