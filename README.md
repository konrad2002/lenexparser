# 📝 LENEX Parser

[![Go](https://github.com/konrad2002/lenexparser/actions/workflows/go.yml/badge.svg)](https://github.com/konrad2002/lenexparser/actions/workflows/go.yml)

<img src="lenexparser.png" align="right" alt="dsvparser logo" width="175">

This go package contains a parser for the [LENEX file format](https://www.wikiwand.com/de/Lenex) which is used as exchange format for swim meetings in Europe. The parser will support the current version 3.0 of the LENEX standard.

If you are looking for a DSV parser, check [konrad2002/dsvparser](https://github.com/konrad2002/dsvparser)

## 💻 Development

This project is currently under development and lacking lots of functionalities.
Feel free to contribute to the project by opening a pull request or getting in touch with [@konrad2002](https://weiss-konrad.de).

### Functionality

*tbd*

## 📥 Usage

Everyone is allowed to use this project for commercial and non-commercial usage.

### Installation

Import the package

```sh
go get github.com/konrad2002/lenexparser@v0.0.0
```

### Example

*tbd*

## 📋 Standard

LENEX is an UTF-8 encoded XML file that is used as an exchange file for European swim meetings. LENEX files can occur in uncompressed way as `.lef` files as well as the compressed `.lxf` archive. Since March 2010 the current standard 3.0 exists and is valid. LENEX stands for "Ligue Européenne de Natation Exchange Format". The Specifications regarding LENEX can be found [🔗 here](https://wiki.swimrankings.net/index.php/swimrankings:Lenex).

### Data Types

Lenex has a list of used data types. These are converted to the following Go data types:

| LENEX                | short | Go              | Notes                                                                                                                                       |
|----------------------|-------|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| String               | s     | `string`        | special character entities will be replaced [XML char entity references](https://www.wikiwand.com/en/XML_entity?mobile-app=true&theme=dark) |
| String international | si    | `string`        |                                                                                                                                             |
| Number               | n     | `int`           |                                                                                                                                             |
| Enumeration          | e     | `const ()`      |                                                                                                                                             |
| Date                 | d     | `time.Time`     |                                                                                                                                             |
| Daytime              | t     | `time.Time`     |                                                                                                                                             |
| Currency             | c     | `int`           |                                                                                                                                             |
| Swim time            | st    | `time.Duration` |                                                                                                                                             |
| Reaction time        | rt    | `time.Duration` |                                                                                                                                             |
| Unique id            | uid   | `string`        |                                                                                                                                             |

## 🏊‍♀️ SwimResults

During the development and improvement of [SwimResults](https://swimresults.de) we needed to parse LENEX files. In this repository a LENEX file parser is developed independently from SwimResults.

## 📄 Changelogs

### v0.0.1

- add example LENEX file
- data types in README.md
- add types:
  - contact

### v0.0.0

- add README.md
