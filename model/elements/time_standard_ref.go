package elements

import "encoding/xml"

type TimeStandardRef struct {
	XMLName            xml.Name `xml:"TIMESTANDARDREF"`
	TimeStandardListId int      `xml:"timestandardlistid,attr"`
	Fee                string   `xml:"FEE"` // TODO
	Marker             string   `xml:"marker,attr"`
}
