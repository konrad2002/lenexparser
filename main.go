package main

import (
	"encoding/xml"
	"fmt"
	"github.com/konrad2002/lenexparser/model/elements"
	"log"
	"os"
)

func main() {
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
	fmt.Println(lenex)
	fmt.Printf("Lenex (%s) generated with '%s' from '%s' in version %s\n", lenex.Version, lenex.Constructor.Name, lenex.Constructor.Contact.Name, lenex.Constructor.Version)

	for _, meet := range lenex.Meets {
		fmt.Printf("At '%s' there are %d teams registered\n", meet.Name, len(meet.Clubs))
	}
}
