package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Judge struct {
	XMLName    xml.Name          `xml:"JUDGE"`
	Number     int               `xml:"number,attr"`
	OfficialId int               `xml:"officialid,attr"`
	Remarks    string            `xml:"remarks,attr"`
	Role       enums.JudgeRole   `xml:"role,attr"`
	Status     enums.JudgeStatus `xml:"status,attr"`
}
