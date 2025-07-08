package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Fee struct {
	XMLName  xml.Name       `xml:"FEE"`
	Currency enums.Currency `xml:"currency,attr"`
	Type     string         `xml:"type,attr"`
	Value    string         `xml:"value,attr"`
}
