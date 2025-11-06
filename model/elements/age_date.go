package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type AgeDate struct {
	XMLName xml.Name          `xml:"AGEDATE"`
	Type    enums.AgeDateType `xml:"type,attr"`
	Value   parser.DateTime   `xml:"value,attr"`
}
