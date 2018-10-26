package main

import (
	"bufio"
	"fmt"
	"gps/pkg"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

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

	fmt.Println("")
	color.Set(color.FgHiGreen)
	fmt.Print("Entry number to install (or enter \"q\" to quit): ")
	color.Unset()
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	var entry int
	if input == "q\n" {
		os.Exit(0)
	} else {
		entry, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			panic(err)
		}
	}

	cmd := exec.Command("go", "get", pkgs.Results[entry-1].Path)
	color.Set(color.FgMagenta)
	fmt.Println("Installing:", pkgs.Results[entry-1].Path)
	color.Unset()
	cmd.Run()

}
