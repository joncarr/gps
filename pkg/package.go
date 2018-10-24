package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Flag byte

const (
	Imports Flag = 1 << iota
	Fork
	Stars
	Description
)

func isNumerical(s string) bool {
	if _, err := strconv.ParseFloat(string(s[0]), 32); err != nil {
		return false
	}
	return true
}

// PKG is package data
type PKG struct {
	EntryNumber string
	Path        string
	Synopsis    string
	Imports     string
	Fork        string
	Stars       string
}

// NewPKG builds a package respresentative
func newPKG(entryNum, path, synopsis, imports, stars, fork string) *PKG {
	return &PKG{
		EntryNumber: entryNum,
		Path:        path,
		Synopsis:    synopsis,
		Imports:     imports,
		Stars:       stars,
		Fork:        fork,
	}
}

func (p *PKG) PrintPackage() {

	fmt.Printf("# ")
	color.Set(color.FgCyan)
	fmt.Printf("%s\t\t", p.EntryNumber)
	color.Unset()
	fmt.Printf("Path: ")
	color.Set(color.FgYellow)
	fmt.Printf("%s\n\n", p.Path)
	color.Unset()
	fmt.Printf("Imports: %s\t", p.Imports)
	fmt.Printf("Fork: ")
	if p.Fork == "Yes" {
		color.Set(color.FgGreen)
	} else {
		color.Set(color.FgRed)
	}
	fmt.Printf("%s\t", p.Fork)
	color.Unset()
	fmt.Printf("Stars: %s\n\n", p.Stars)
	fmt.Printf("Synopsis:\t")
	color.Set(color.FgWhite)
	fmt.Printf("%s\n\n", p.Synopsis)
	color.Unset()
	color.Set(color.FgBlue)
	fmt.Printf("======================================================================\n\n")
	color.Unset()

}

func BuildPackageList(results []string) []*PKG {

	var PKGs []*PKG

	for _, d := range results {
		dSplit := strings.Split(d, " ")

		classification := classifyEntry(dSplit)

		entryNum := dSplit[0]
		url := dSplit[1]
		imports := dSplit[2]
		var fork string
		var synopsis string
		var stars string

		switch classification {
		case (Imports):
			fork = "No"
			stars = "0"
			synopsis = ""
		case (Imports | Fork):
			fork = "Yes"
			stars = "0"
			synopsis = ""
		case (Imports | Stars):
			fork = "No"
			stars = dSplit[5]
			synopsis = ""
		case (Imports | Fork | Stars):
			fork = "Yes"
			stars = dSplit[7]
			synopsis = ""
		case (Imports | Description):
			fork = "No"
			stars = "0"
			synopsis = strings.Join(dSplit[4:], " ")
		case (Imports | Fork | Description):
			fork = "Yes"
			stars = "0"
			synopsis = strings.Join(dSplit[6:], " ")
		case (Imports | Stars | Description):
			fork = "No"
			stars = dSplit[5]
			synopsis = strings.Join(dSplit[7:], " ")
		default:
			fork = "No"
			stars = dSplit[5]
			synopsis = strings.Join(dSplit[7:], " ")
		}

		if synopsis == "" {
			synopsis = "No description available"
		}

		pkg := newPKG(entryNum, url, synopsis, imports, stars, fork)
		PKGs = append(PKGs, pkg)
	}

	return PKGs
}

func PrintPackageList(pkgs []*PKG) {
	for _, p := range pkgs {
		p.PrintPackage()
	}
}

func classifyEntry(s []string) Flag {
	classification := Imports
	// imports 0001
	// fork 0010
	// stars 0100
	// description 1000

	if isFork(s) {
		classification += Fork
	}
	if hasStars(s) {
		classification += Stars
	}
	if hasDescription(s) {
		classification += Description
	}

	return classification

}

func hasImports(s []string) bool {
	if s[3] == "imports" {
		return true
	}
	return false
}

func hasStars(s []string) bool {
	if len(s) < 7 {
		return false
	}
	if s[6] == "stars" || s[8] == "stars" {
		return true
	}
	return false
}

func isFork(s []string) bool {
	if len(s) < 6 {
		return false
	}
	if s[5] == "fork" {
		return true
	}
	return false
}

func hasDescription(s []string) bool {
	if hasImports(s) && isFork(s) && hasStars(s) && len(s) >= 10 {
		return true
	} else if hasImports(s) && !isFork(s) && hasStars(s) && len(s) >= 8 {
		return true
	} else if hasImports(s) && isFork(s) && !hasStars(s) && len(s) >= 7 {
		return true
	} else if hasImports(s) && !isFork(s) && !hasStars(s) && len(s) >= 5 {
		return true
	}
	return false
}

// 0								[0] index
// github.com/ChimeraCoder/anaconda [1] url
// 214								[2] imports
// imports
// ·
// 917								[5] stars
// stars
// Package							[7:] synopsis
// anaconda
// provides
// structs
// and
// functions
// for
// accessing
// version
// 1.1
// of
// the
// Twitter
// API.
