package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/parser"
)

type TimeStandard struct {
	XMLName   xml.Name        `xml:"TIMESTANDARD"`
	SwimStyle SwimStyle       `xml:"SWIMSTYLE"`
	SwimTime  parser.SwimTime `xml:"swimtime,attr"`
}
