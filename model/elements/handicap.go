package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Handicap struct {
	XMLName      xml.Name             `xml:"HANDICAP"`
	Breast       enums.HandicapClass  `xml:"breast,attr"`
	BreastStatus enums.HandicapStatus `xml:"breaststatus,attr"`
	Exception    string               `xml:"exception,attr"`
	Free         enums.HandicapClass  `xml:"free,attr"`
	FreeStatus   enums.HandicapStatus `xml:"freestatus,attr"`
	Medley       enums.HandicapClass  `xml:"medley,attr"`
	MedleyStatus enums.HandicapStatus `xml:"medleystatus,attr"`
}
