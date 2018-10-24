package main

import (
	"flag"
	"gps/pkg"
	"gps/scraper"
)

const (
	goDocURL string = "https://godoc.org/?q="
)

var searchTerm string

func init() {
	const (
		defaultValue = "gorilla"
		usage        = "A package search term you are looking for"
	)
	flag.StringVar(&searchTerm, "find", defaultValue, usage)
	flag.StringVar(&searchTerm, "f", defaultValue, usage+" (shorthand)")
}

func main() {
	res := scraper.BuildResults(goDocURL, "twitter")
	pkgs := pkg.BuildPackageList(res)

	pkg.PrintPackageList(pkgs)
}
