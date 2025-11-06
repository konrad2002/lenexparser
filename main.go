package main

import (
	"encoding/xml"
	"fmt"
	"github.com/konrad2002/lenexparser/model/elements"
	"log"
	"math/rand"
	"os"
)

func main() {}

func PrintEventList() {
	xmlString, err := os.ReadFile("assets/231209-Marienberg-PR.lef")
	//xmlString, err := os.ReadFile("assets/mdm25.lef")
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

	fmt.Printf("\n\nTeams:\n")

	var athletes []elements.Athlete

	for _, club := range meet.Clubs {
		fmt.Printf("\tClub %s (%s):\n", club.Name, club.Nation)
		for _, athlete := range club.Athletes {
			athletes = append(athletes, athlete)
			fmt.Printf("\t\t%s, %s – %d\n", athlete.Lastname, athlete.Firstname, athlete.Birthdate.Year())
		}
	}

	athlete1 := athletes[rand.Intn(len(athletes))]

	fmt.Printf("\n\nEntries for %s, %s – %d – %s\n", athlete1.Lastname, athlete1.Firstname, athlete1.Birthdate.Year(), athlete1.Club.Name)

	for _, entry := range athlete1.Entries {
		fmt.Printf("\tWK %d – %dm %s: %s\n", entry.EventId, entry.EntryDistance, entry.EntryCourse, entry.EntryTime)
	}
}
