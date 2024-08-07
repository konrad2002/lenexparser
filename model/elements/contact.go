package elements

import (
	"encoding/xml"
)

type Contact struct {
	XMLName  xml.Name `xml:"CONTACT"`
	City     string   `xml:"city,attr"`
	Country  string   `xml:"country,attr"`
	Email    string   `xml:"email,attr"`
	Fax      string   `xml:"fax,attr"`
	Internet string   `xml:"internet,attr"`
	Name     string   `xml:"name,attr"`
	Mobile   string   `xml:"mobile,attr"`
	Phone    string   `xml:"phone,attr"`
	State    string   `xml:"state,attr"`
	Street   string   `xml:"street,attr"`
	Street2  string   `xml:"street2,attr"`
	Zip      string   `xml:"zip,attr"`
}
