package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/konrad2002/lenexparser/parser"
)

type RecordList struct {
	XMLName  xml.Name             `xml:"RECORDLIST"`
	AgeGroup AgeGroup             `xml:"AGEGROUP"`
	Course   enums.Course         `xml:"course,attr"`
	Gender   enums.EventGender    `xml:"gender,attr"`
	Handicap enums.HandicapClass  `xml:"handicap,attr"`
	Name     string               `xml:"name,attr"`
	Nation   string               `xml:"nation,attr"`
	Order    int                  `xml:"order,attr"`
	Records  []Record             `xml:"RECORDS>RECORD"`
	Region   string               `xml:"region,attr"`
	Updated  parser.CustomTime    `xml:"updated,attr"`
	Type     enums.RecordListType `xml:"type,attr"`
}
