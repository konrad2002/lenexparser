package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Official struct {
	XMLName    xml.Name     `xml:"OFFICIAL"`
	Contact    Contact      `xml:"CONTACT"`
	FirstName  string       `xml:"firstname,attr"`
	Gender     enums.Gender `xml:"gender,attr"`
	Grade      string       `xml:"grade,attr"`
	LastName   string       `xml:"lastname,attr"`
	License    string       `xml:"license,attr"`
	NamePrefix string       `xml:"nameprefix,attr"`
	Nation     enums.Nation `xml:"nation,attr"`
	OfficialId int          `xml:"officialid,attr"`
	Passport   string       `xml:"passport,attr"`
}
