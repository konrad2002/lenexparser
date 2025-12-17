package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Club struct {
	XMLName     xml.Name       `xml:"CLUB"`
	Athletes    []Athlete      `xml:"ATHLETES>ATHLETE"`
	Code        string         `xml:"code,attr"`
	Contact     Contact        `xml:"CONTACT"`
	Name        string         `xml:"name,attr"`
	NameEn      string         `xml:"name.en,attr"`
	Nation      enums.Nation   `xml:"nation,attr"`
	Number      int            `xml:"number,attr"`
	Officials   []Official     `xml:"OFFICIALS>OFFICIAL"`
	Coaches     []Coach        `xml:"COACHES>COACH"`
	Region      string         `xml:"region,attr"`
	Relays      []Relay        `xml:"RELAYS>RELAY"`
	ShortName   string         `xml:"shortname,attr"`
	ShortNameEn string         `xml:"shortname.en,attr"`
	SWRid       int            `xml:"swrid,attr"`
	Type        enums.ClubType `xml:"type,attr"`
}
