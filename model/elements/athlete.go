package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Athlete struct {
	XMLName     xml.Name            `xml:"ATHLETE"`
	AthleteId   int                 `xml:"athleteid,attr"`
	Birthdate   parser.CustomTime   `xml:"birthdate,attr"`
	Club        Club                `xml:"CLUB"`
	Entries     int                 `xml:"ENTRIES"` // TODO
	Firstname   string              `xml:"firstname,attr"`
	FirstnameEn string              `xml:"firstname.en,attr"`
	Gender      enums.Gender        `xml:"gender,attr"`
	Handicap    int                 `xml:"HANDICAP"` // TODO
	Lastname    string              `xml:"lastname,attr"`
	LastnameEn  string              `xml:"lastname.en,attr"`
	Level       string              `xml:"level,attr"`
	License     string              `xml:"license,attr"`
	LicenseIpc  int                 `xml:"license_ipc,attr"`
	NamePrefix  string              `xml:"nameprefix,attr"`
	Nation      enums.Nation        `xml:"nation,attr"`
	Passport    string              `xml:"passport,attr"`
	Results     int                 `xml:"RESULTS"` // TODO
	Status      enums.AthleteStatus `xml:"status,attr"`
	SWRid       int                 `xml:"swrid,attr"`
}
