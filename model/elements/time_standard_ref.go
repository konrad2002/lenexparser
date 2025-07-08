package elements

import "encoding/xml"

type TimeStandardRef struct {
	XMLName            xml.Name `xml:"TIMESTANDARDREF"`
	TimeStandardListId int      `xml:"timestandardlistid,attr"`
	Fee                Fee      `xml:"FEE"`
	Marker             string   `xml:"marker,attr"`
}
