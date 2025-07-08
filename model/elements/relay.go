package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Relay struct {
	XMLName        xml.Name            `xml:"RELAY"`
	AgeMax         int                 `xml:"agemax,attr"`
	AgeMin         int                 `xml:"agemin,attr"`
	AgeTotalMax    int                 `xml:"agetotalmax,attr"`
	AgeTotalMin    int                 `xml:"agetotalmin,attr"`
	Club           Club                `xml:"CLUB"`
	Entries        []Entry             `xml:"ENTRIES>ENTRY"`
	Gender         enums.EventGender   `xml:"gender,attr"`
	Handicap       enums.HandicapClass `xml:"handicap,attr"`
	Name           string              `xml:"name,attr"`
	Number         int                 `xml:"number,attr"`
	RelayPositions []RelayPosition     `xml:"RELAYPOSITIONS,RELAYPOSITION"`
	Results        []Result            `xml:"RESULTS>RESULT"`
}
