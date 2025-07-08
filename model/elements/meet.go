package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Meet struct {
	XMLName           xml.Name           `xml:"MEET"`
	AgeDate           AgeDate            `xml:"AGEDATE"`
	Bank              Bank               `xml:"BANK"`
	Altitude          int                `xml:"altitude,attr"`
	City              string             `xml:"city,attr"`
	CityEn            string             `xml:"city.en,attr"`
	Clubs             []Club             `xml:"CLUBS>CLUB"`
	Contact           Contact            `xml:"CONTACT"`
	Course            enums.Course       `xml:"course,attr"`
	Deadline          parser.CustomTime  `xml:"deadline,attr"`
	DeadlineTime      parser.CustomTime  `xml:"deadlinetime,attr"`
	EntryStartDate    parser.CustomTime  `xml:"entrystartdate,attr"`
	EntryType         enums.EntryType    `xml:"entrytype,attr"`
	Facility          Facility           `xml:"FACILITY"`
	Fees              []Fee              `xml:"FEES>FEE"`
	HostClub          string             `xml:"hostclub,attr"`
	HostClubUrl       string             `xml:"hostclub.url,attr"`
	MaxEntriesAthlete int                `xml:"maxentriesathlete,attr"`
	MaxEntriesRelay   int                `xml:"maxentriesrelay,attr"`
	Name              string             `xml:"name,attr"`
	NameEn            string             `xml:"name.en,attr"`
	Nation            enums.Nation       `xml:"nation,attr"`
	Number            string             `xml:"number,attr"`
	Organizer         string             `xml:"organizer,attr"`
	OrganizerUrl      string             `xml:"organizer.url,attr"`
	PointTable        PointTable         `xml:"POINTTABLE"`
	Pool              Pool               `xml:"POOL"`
	Qualify           Qualify            `xml:"QUALIFY"`
	ReserveCount      int                `xml:"reservecount,attr"`
	ResultUrl         string             `xml:"result.url,attr"`
	Sessions          []Session          `xml:"SESSIONS>SESSION"`
	StartMethod       enums.StartMethod  `xml:"startmethod,attr"`
	Status            enums.MeetStatus   `xml:"status,attr"`
	SWRid             string             `xml:"swrid,attr"`
	Timing            enums.Timing       `xml:"timing,attr"`
	TouchpadMode      enums.TouchpadMode `xml:"touchpadmode,attr"`
	Type              string             `xml:"type,attr"`
	WithdrawUntil     parser.CustomTime  `xml:"withdrawuntil,attr"`
}
