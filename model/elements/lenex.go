package elements

import (
	"encoding/xml"
	"github.com/konrad2002/lenexparser/parser"
)

type Lenex struct {
	XMLName           xml.Name           `xml:"LENEX"`
	Constructor       Constructor        `xml:"CONSTRUCTOR"`
	Created           parser.DateTime    `xml:"created,attr"`
	Meets             []Meet             `xml:"MEETS>MEET"`
	RecordLists       []RecordList       `xml:"RECORDLISTS>RECORDLIST"`
	RevisionDate      parser.DateTime    `xml:"revisiondate,attr"`
	TimeStandardLists []TimeStandardList `xml:"TIMESTANDARDLISTS>TIMESTANDARDLIST"`
	Version           string             `xml:"version,attr"`
}
