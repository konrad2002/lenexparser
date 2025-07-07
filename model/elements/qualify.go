package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Qualify struct {
	XMLName    xml.Name                `xml:"QUALIFY"`
	Conversion enums.QualifyConversion `xml:"conversion,attr"`
	From       parser.CustomTime       `xml:"from,attr"`
	Percent    int                     `xml:"percent,attr"`
	Until      parser.CustomTime       `xml:"until,attr"`
}
