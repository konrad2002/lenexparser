package elements

import "encoding/xml"

type Bank struct {
	XMLName       xml.Name `xml:"BANK"`
	AccountHolder string   `xml:"accountholder,attr"`
	BIC           string   `xml:"bic,attr"`
	IBAN          string   `xml:"iban,attr"`
	Name          string   `xml:"name,attr"`
	Note          string   `xml:"note,attr"`
}
