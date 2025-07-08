package main

import (
	"encoding/xml"
	"fmt"
	"github.com/konrad2002/lenexparser/model/elements"
	"log"
	"os"
)

func main() {}

func PrintEventList() {
	//xmlString, err := os.ReadFile("assets/231209-Marienberg-PR.lef")
	xmlString, err := os.ReadFile("assets/mdm25.lef")
	if err != nil {
		log.Fatal(err)
	}

	var lenex elements.Lenex
	err = xml.Unmarshal(xmlString, &lenex)
	if err != nil {
		log.Fatal(err)
	}

	meet := lenex.Meets[0]
	fmt.Printf("Events at '%s %d'\n", meet.Name, meet.Sessions[0].Date.Year())

	for _, session := range meet.Sessions {
		fmt.Printf("\tSession %d:\n", session.Number)

		for _, event := range session.Events {
			fmt.Printf("\t\tEvent %d - %s\n", event.Number, event.SwimStyle.Name)
		}
	}
}
