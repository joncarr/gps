package cmd

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
	"github.com/spf13/cobra"
)

var (
	searchTerm string
	limitOpt   = 100
	sortOpt    = "imports"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:     "find",
	Short:   fmt.Sprint("\nTell gps what you're searching for, it's like Tom-Tom but...it's not"),
	Aliases: []string{"f"},
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		searchTerm = args[0]

		resp, err := http.Get(GoDocURL + searchTerm)
		if err != nil {
			return err
		}

		pkgs := pkg.BuildPackageList(resp)
		pkgs.Sort(sortOpt)

		for i := 0; i < limitOpt; i++ {
			fmt.Println("Entry: ", i+1)
			pkgs.Results[i].PrintPackage()
		}

		color.Set(color.FgHiGreen)
		fmt.Print("\n\nEntry number to install (or enter \"q\" to quit): ")
		color.Unset()

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}

		var entry int
		if input == "q\n" {
			os.Exit(0)
		} else {
			entry, err = strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				return err
			}
		}

		command := exec.Command("go", "get", "-u", "-v", pkgs.Results[entry-1].Path)
		color.Set(color.FgMagenta)
		fmt.Println("Installing:", pkgs.Results[entry-1].Path)
		color.Unset()
		command.Run()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// findCmd.Flags().StringP("sort", "s", "", "Method by which to sort. Options:[\"alpha\": to sort alphabetically by package name]")
	findCmd.Flags().IntVarP(&limitOpt, "limit", "l", 100, "Limit your search results (1-100)")
	findCmd.Flags().StringVarP(&sortOpt, "sort", "s", "imports", sortOptionsDesc())

}

func sortOptionsDesc() string {
	return `Method by which to sort.

Available sort options:

"alpha": to sort alphabetically by package name
"score": to sort by score
"stars": to sort by stars
"imports": to sort by imports
`
}
