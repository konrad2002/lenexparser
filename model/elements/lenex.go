package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/parser"
)

type Lenex struct {
	XMLName     xml.Name          `xml:"LENEX"`
	Constructor Constructor       `xml:"CONSTRUCTOR"`
	Created     parser.CustomTime `xml:"created,attr"`
	Meets       []Meet            `xml:"MEETS>MEET"` // TODO
	//RecordLists       string    `xml:"RECORDLISTS"` // TODO
	RevisionDate parser.CustomTime `xml:"revisiondate,attr"`
	//TimeStandardLists string    `xml:"TIMESTANDARDLISTS"` // TODO
	Version string `xml:"version,attr"`
}
