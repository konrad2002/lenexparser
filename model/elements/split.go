package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/parser"
)

type Split struct {
	XMLName  xml.Name        `xml:"SPLIT"`
	Distance int             `xml:"distance,attr"`
	SwimTime parser.SwimTime `xml:"swimtime,attr"`
}
