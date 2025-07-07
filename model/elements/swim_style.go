package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type SwimStyle struct {
	XMLName     xml.Name        `xml:"SWIMSTYLE"`
	Code        string          `xml:"code,attr"`
	Distance    int             `xml:"distance,attr"`
	Name        string          `xml:"name,attr"`
	RelayCount  int             `xml:"relaycount,attr"`
	Stroke      enums.Stroke    `xml:"stroke,attr"`
	SwimStyleId int             `xml:"swimstyleid,attr"`
	Technique   enums.Technique `xml:"technique,attr"`
}
