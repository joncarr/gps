package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

// PKG is package data
type PKG struct {
	Name        string  `json:"name"`
	Path        string  `json:"path"`
	ImportCount int     `json:"import_count"`
	Synopsis    string  `json:"synopsis"`
	Stars       int     `json:"stars"`
	Score       float32 `json:"score"`
}

// PackageList is self explanatory
type PackageList struct {
	Results []PKG `json:"results"`
}

// PrintPackage prints the package information
func (p *PKG) PrintPackage() {

	fmt.Printf("Name: ")
	color.Set(color.FgCyan)
	fmt.Printf("%s\n", p.Name)
	color.Unset()
	fmt.Printf("Path: ")
	color.Set(color.FgYellow)
	fmt.Printf("%s\n", p.Path)
	color.Unset()
	fmt.Printf("Score: ")
	color.Set(color.FgGreen)
	fmt.Printf("%f\n", p.Score)
	color.Unset()
	fmt.Printf("Imports: %d\n", p.ImportCount)
	fmt.Printf("Stars: %d\n", p.Stars)
	fmt.Printf("Synopsis: ")
	color.Set(color.FgWhite)
	if p.Synopsis == "" {
		color.Set(color.FgRed)
		fmt.Printf("%s\n", "No synopsis was provided")
		color.Unset()
	}
	fmt.Printf("%s\n", p.Synopsis)
	color.Unset()
	color.Set(color.FgBlue)
	fmt.Printf("======================================================================\n\n")
	color.Unset()
}

// PrintPackageList prints the list of collected packages
func PrintPackageList(pkgs []*PKG) {
	for _, p := range pkgs {
		p.PrintPackage()
	}
}

// BuildPackageList builds a package list from response body
func BuildPackageList(d *http.Response) PackageList {
	pkgs := PackageList{}
	json.NewDecoder(d.Body).Decode(&pkgs)
	defer d.Body.Close()

	return pkgs

}
