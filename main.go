package main

import (
	"fmt"
	"gps/pkg"
	"net/http"
	"os"

	"github.com/fatih/color"
)

const (
	goDocURL string = "https://api.godoc.org/search?&q="
)

var searchTerm string

func main() {
	if len(os.Args) < 2 {
		color.Set(color.FgRed)
		fmt.Println("Usage: gps [search term]")
		color.Unset()
		os.Exit(1)
	} else if len(os.Args) > 2 {
		color.Set(color.FgHiBlue)
		fmt.Println("Usage: 'gps [search term]' :: Limit to one search term")
		color.Unset()
		os.Exit(1)
	}

	searchTerm = os.Args[1]

	resp, err := http.Get(goDocURL + searchTerm)
	if err != nil {
		fmt.Printf("not able to receive response: %v", err)
	}

	pkgs := pkg.BuildPackageList(resp)

	fmt.Println("")
	for i, ii := range pkgs.Results {
		fmt.Println("Entry: ", i+1)
		ii.PrintPackage()
	}

}
