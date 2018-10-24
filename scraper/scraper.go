package scraper

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var nodeResults []*html.Node

func BuildResults(url, searchTerm string) []string {
	resp, err := http.Get(url + searchTerm)
	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	nodes := parseRootNode(root, resp)
	nodes = nodes[1:]

	var rtn []string
	for i, d := range nodes {
		str := fmt.Sprintf("%d %s", i+1, scrape.Text(d))
		rtn = append(rtn, str)
	}

	return rtn

}

// ParseRootNode returns []*html.Node or an error
// In this particular instance it will be all Table Rows
func parseRootNode(root *html.Node, resp *http.Response) []*html.Node {

	tables := scrape.FindAll(root, scrape.ByTag(atom.Table))

	table := tables[0]

	td := scrape.FindAll(table, scrape.ByTag(atom.Tr))

	return td
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
