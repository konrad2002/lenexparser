package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/model/enums"
)

type Ranking struct {
	XMLName  xml.Name    `xml:"RANKING"`
	Medal    enums.Medal `xml:"medal,attr"`
	Order    int         `xml:"order,attr"`
	Place    int         `xml:"place,attr"`
	ResultId int         `xml:"resultid,attr"`
}
