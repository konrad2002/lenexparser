package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type Session struct {
	XMLName           xml.Name           `xml:"SESSION"`
	Course            enums.Course       `xml:"course,attr"`
	Date              parser.CustomTime  `xml:"date,attr"`
	Daytime           parser.CustomTime  `xml:"daytime,attr"`
	Endtime           parser.CustomTime  `xml:"endtime,attr"`
	Events            []Event            `xml:"EVENTS>EVENT"`
	Fees              []Fee              `xml:"FEES>FEE"`
	Judges            []Judge            `xml:"JUDGES>JUDGE"`
	MaxEntriesAthlete int                `xml:"maxentriesathlete,attr"`
	MaxEntriesRelay   int                `xml:"maxentriesrelay,attr"`
	Name              string             `xml:"name,attr"`
	Number            int                `xml:"number,attr"`
	OfficialMeeting   parser.CustomTime  `xml:"officialmeeting,attr"`
	Pool              Pool               `xml:"POOL"`
	RemarksJudge      string             `xml:"remarksjudge,attr"`
	Status            enums.MeetStatus   `xml:"status,attr"`
	TeamLeaderMeeting parser.CustomTime  `xml:"teamleadermeeting,attr"`
	Timing            enums.Timing       `xml:"timing,attr"`
	TouchpadMode      enums.TouchpadMode `xml:"touchpadmode,attr"`
	WarmupFrom        parser.CustomTime  `xml:"warmupfrom,attr"`
	WarmupUntil       parser.CustomTime  `xml:"warmupuntil,attr"`
}
