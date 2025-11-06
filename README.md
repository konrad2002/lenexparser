# 📝 LENEX Parser

[![Go](https://github.com/konrad2002/lenexparser/actions/workflows/go.yml/badge.svg)](https://github.com/konrad2002/lenexparser/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/konrad2002/lenexparser.svg)](https://pkg.go.dev/github.com/konrad2002/lenexparser)

<img src="lenexparser.png" align="right" alt="dsvparser logo" width="175">

This go package contains a parser for the [LENEX file format](https://www.wikiwand.com/de/Lenex) which is used as exchange format for swim meetings in Europe. The parser will support the current version 3.0 of the LENEX standard.

If you are looking for a DSV parser, check [konrad2002/dsvparser](https://github.com/konrad2002/dsvparser)

## 💻 Development

This project is currently under development and might lack some functionalities.
Feel free to contribute to the project by opening a pull request or getting in touch with [@konrad2002](https://weiss-konrad.de).

### Functionality

Using the `encoding/xml` package this project parses the given LENEX XML file and provides data types for all elements from the LENEX standard (see list below). 

## 📥 Usage

Everyone is allowed to use this project for commercial and non-commercial usage.

### Installation

Import the package

```sh
go get github.com/konrad2002/lenexparser@v0.1.0
```

### Example

The following example prints a list of all events of a meeting.

```go
import (
	...
	"github.com/konrad2002/lenexparser/model/elements"
)

func PrintEventList() {
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
```

## 📋 Standard

LENEX is an UTF-8 encoded XML file that is used as an exchange file for European swim meetings. LENEX files can occur in uncompressed way as `.lef` files as well as the compressed `.lxf` archive. Since March 2010 the current standard 3.0 exists and is valid. LENEX stands for "Ligue Européenne de Natation Exchange Format". The Specifications regarding LENEX can be found [🔗 here](https://wiki.swimrankings.net/index.php/swimrankings:Lenex).

### Data Types

Lenex has a list of used data types. These are converted to the following Go data types:

| LENEX                | short | Go                  | Notes                                                                                                                                            |
|----------------------|-------|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------|
| Currency             | c     | `int`               |                                                                                                                                                  |
| Date                 | d     | `parser.DateTime`   |                                                                                                                                                  |
| Daytime              | t     | `parser.DateTime`   |                                                                                                                                                  |
| Enumeration          | e     | `const ()`          |                                                                                                                                                  |
| Number               | n     | `int`               |                                                                                                                                                  |
| Global Identifier    | uuid  | `string`            |                                                                                                                                                  |
| String               | s     | `string`            | special character entities will be replaced, see [XML char entity references](https://www.wikiwand.com/en/XML_entity?mobile-app=true&theme=dark) |
| String international | si    | `string`            |                                                                                                                                                  |
| Swim time            | st    | `parser.SwimTime`   |                                                                                                                                                  |
| Timestamp            | ts    | `parser.DateTime`   |                                                                                                                                                  |
| Reaction time        | rt    | `parser.SwimTime`   |                                                                                                                                                  |
| Unique id            | uid   | `string`            |                                                                                                                                                  |

### Elements + Implementation Status

| Item          | Item         | Item               |
|---------------|--------------|--------------------|
| ✅ AGEDATE     | ✅ HANDICAP   | ✅ RECORDLIST       |
| ✅ AGEGROUP    | ✅ HEAT       | ✅ RELAY            |
| ✅ ATHLETE     | ✅ JUDGE      | ✅ RELAYPOSITION    |
| ✅ BANK        | ✅ LENEX      | ✅ RESULT           |
| ✅ CLUB        | ✅ MEET       | ✅ SESSION          |
| ✅ COACH       | ✅ MEETINFO   | ✅ SPLIT            |
| ✅ CONSTRUCTOR | ✅ OFFICIAL   | ✅ SWIMSTYLE        |
| ✅ CONTACT     | ✅ POINTTABLE | ✅ TIMESTANDARD     |
| ✅ ENTRY       | ✅ POOL       | ✅ TIMESTANDARDLIST |
| ✅ EVENT       | ✅ QUALIFY    | ✅ TIMESTANDARDREF  |
| ✅ FACILITY    | ✅ RANKING    |                    |
| ✅ FEE         | ✅ RECORD     |                    |



## 🏊‍♀️ SwimResults

During the development and improvement of [SwimResults](https://swimresults.de) we needed to parse LENEX files. In this repository a LENEX file parser is developed independently from SwimResults.

## 📄 Changelogs

### v0.1.0

- add all elements
- add all enums
- parser for time, date and swim time

### v0.0.1

- add example LENEX file
- data types in README.md
- add types:
  - contact

### v0.0.0

- add README.md
