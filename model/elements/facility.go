package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Facility struct {
	XMLName xml.Name     `xml:"FACILITY"`
	City    string       `xml:"city,attr"`
	Nation  enums.Nation `xml:"nation,attr"`
	Name    string       `xml:"name,attr"`
	State   string       `xml:"state,attr"`
	Street  string       `xml:"street,attr"`
	Street2 string       `xml:"street2,attr"`
	Zip     string       `xml:"zip,attr"`
}
