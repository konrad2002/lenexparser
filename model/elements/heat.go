package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Heat struct {
	XMLName    xml.Name         `xml:"HEAT"`
	AgeGroupId int              `xml:"agegroupid,attr"`
	Daytime    parser.DateTime  `xml:"daytime,attr"`
	Final      enums.Final      `xml:"final,attr"`
	HeatId     int              `xml:"heatid,attr"`
	Number     int              `xml:"number,attr"`
	Order      int              `xml:"order,attr"`
	Status     enums.HeatStatus `xml:"status,attr"`
}
