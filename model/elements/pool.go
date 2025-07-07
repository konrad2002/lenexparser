package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Pool struct {
	XMLName     xml.Name       `xml:"POOL"`
	LaneMax     int            `xml:"lanemax,attr"`
	LaneMin     int            `xml:"lanemin,attr"`
	Temperature int            `xml:"temperature,attr"`
	Type        enums.PoolType `xml:"type,attr"`
}
