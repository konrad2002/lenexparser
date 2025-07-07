package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Coach struct {
	XMLName    xml.Name        `xml:"COACH"`
	Contact    Contact         `xml:"CONTACT"`
	FirstName  string          `xml:"firstname,attr"`
	Gender     enums.Gender    `xml:"gender,attr"`
	LastName   string          `xml:"lastname,attr"`
	NamePrefix string          `xml:"nameprefix,attr"`
	Nation     enums.Nation    `xml:"nation,attr"`
	Passport   string          `xml:"passport,attr"`
	Type       enums.CoachType `xml:"type,attr"`
}
