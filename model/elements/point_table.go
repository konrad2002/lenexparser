package elements

import "encoding/xml"

type PointTable struct {
	XMLName      xml.Name `xml:"POINTTABLE"`
	Name         string   `xml:"name,attr"`
	PointTableId int      `xml:"pointtableid,attr"`
	Version      string   `xml:"version,attr"`
}
