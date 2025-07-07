package elements

import "encoding/xml"

type Constructor struct {
	XMLName      xml.Name `xml:"CONSTRUCTOR"`
	Contact      Contact  `xml:"CONTACT"`
	Name         string   `xml:"name,attr"`
	Registration string   `xml:"registration,attr"`
	Version      string   `xml:"version,attr"`
}
