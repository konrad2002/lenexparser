package elements

import "encoding/xml"

type Athlete struct {
	XMLName   xml.Name `xml:"ATHLETE"`
	AthleteId int      `xml:"athleteid,attr"`
}
