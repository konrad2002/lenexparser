package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type AgeGroup struct {
	XMLName    xml.Name                `xml:"AGEGROUP"`
	AgeGroupId int                     `xml:"agegroupid,attr"`
	AgeMax     int                     `xml:"agemax,attr"`
	AgeMin     int                     `xml:"agemin,attr"`
	Gender     enums.AgeGroupGender    `xml:"gender,attr"`
	Calculate  enums.AgeGroupCalculate `xml:"calculate,attr"`
	Handicap   enums.Handicap          `xml:"handicap,attr"`
	LevelMax   string                  `xml:"levelmax,attr"`
	LevelMin   string                  `xml:"levelmin,attr"`
	Levels     string                  `xml:"levels,attr"`
	Name       string                  `xml:"name,attr"`
	Rankings   []Ranking               `xml:"RANKINGS>RANKING"`
}
