package tests

import (
	"encoding/xml"
	"fmt"
	"github.com/konrad2002/lenexparser/model/elements"
	"github.com/konrad2002/lenexparser/model/enums"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"os"
	"testing"
)

func TestExample1(t *testing.T) {
	xmlString, err := os.ReadFile("../assets/231209-Marienberg-PR.lef")
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

	assert.Equal(t, 4, len(meet.Sessions))
	assert.Equal(t, 11, len(meet.Sessions[0].Events))
	assert.Equal(t, 27, len(meet.Clubs))

	fmt.Printf("Events at '%s %d'\n", meet.Name, meet.Sessions[0].Date.Year())

	for _, session := range meet.Sessions {
		fmt.Printf("\tSession %d:\n", session.Number)

		for _, event := range session.Events {
			var style string
			var gender string
			var round string
			distance := ""

			switch event.SwimStyle.Stroke {
			case enums.StrokeFly:
				style = "Schmetterling"
				break
			case enums.StrokeBack:
				style = "Rücken"
				break
			case enums.StrokeBreast:
				style = "Brust"
				break
			case enums.StrokeFree:
				style = "Freistil"
				break
			case enums.StrokeMedley:
				style = "Lagen"
				break
			}

			switch event.Gender {
			case enums.EventGenderMale:
				gender = "männlich"
				break
			case enums.EventGenderFemale:
				gender = "weiblich"
				break
			case enums.EventGenderMixed:
				gender = "mixed"
				break
			case enums.EventGenderAll:
				gender = "alle"
				break
			}

			switch event.Round {
			case enums.RoundPRE:
				round = "(Vorlauf)"
				break
			case enums.RoundFIN:
				round = "(Finale)"
				break
			default:
				round = ""
			}

			if event.SwimStyle.RelayCount > 1 {
				distance = fmt.Sprintf("%dx%dm", event.SwimStyle.RelayCount, event.SwimStyle.Distance)
			} else {
				distance = fmt.Sprintf("%dm", event.SwimStyle.Distance)
			}

			fmt.Printf("\t\tWK %d – %s %s %s %s\n", event.Number, distance, style, gender, round)
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

	println("----- AGE GROUPS -----")

	for _, session := range meet.Sessions {
		fmt.Printf("\tSession %d:\n", session.Number)

		for _, event := range session.Events {
			fmt.Printf("\t\tEvent %d:\n", event.Number)

			for _, ageGroup := range event.AgeGroups {
				fmt.Printf("\t\t\tAge Group %v:\n", ageGroup)
			}
		}
	}
}

func TestExample2(t *testing.T) {
	xmlString, err := os.ReadFile("../assets/MeldungenStolle25.lef")
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

	assert.Equal(t, 5, len(meet.Sessions))
	assert.Equal(t, 6, len(meet.Sessions[0].Events))
	assert.Equal(t, 66, len(meet.Clubs))

	fmt.Printf("Events at '%s %d'\n", meet.Name, meet.Sessions[0].Date.Year())

	for _, session := range meet.Sessions {
		fmt.Printf("\tSession %d:\n", session.Number)

		for _, event := range session.Events {
			var style string
			var gender string
			var round string
			distance := ""

			switch event.SwimStyle.Stroke {
			case enums.StrokeFly:
				style = "Schmetterling"
				break
			case enums.StrokeBack:
				style = "Rücken"
				break
			case enums.StrokeBreast:
				style = "Brust"
				break
			case enums.StrokeFree:
				style = "Freistil"
				break
			case enums.StrokeMedley:
				style = "Lagen"
				break
			}

			switch event.Gender {
			case enums.EventGenderMale:
				gender = "männlich"
				break
			case enums.EventGenderFemale:
				gender = "weiblich"
				break
			case enums.EventGenderMixed:
				gender = "mixed"
				break
			case enums.EventGenderAll:
				gender = "alle"
				break
			}

			switch event.Round {
			case enums.RoundPRE:
				round = "(Vorlauf)"
				break
			case enums.RoundFIN:
				round = "(Finale)"
				break
			default:
				round = ""
			}

			if event.SwimStyle.RelayCount > 1 {
				distance = fmt.Sprintf("%dx%dm", event.SwimStyle.RelayCount, event.SwimStyle.Distance)
			} else {
				distance = fmt.Sprintf("%dm", event.SwimStyle.Distance)
			}

			fmt.Printf("\t\tWK %d – %s %s %s %s\n", event.Number, distance, style, gender, round)
		}
	}
}
