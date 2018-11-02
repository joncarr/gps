package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	// GoDocURL is the URL for godoc.org api
	GoDocURL string = "https://api.godoc.org/search?&q="
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "gps",
	Long: `
		Go Package Search (gps) provides a means to search godoc.org via the command-line.
		It's extremely easy to use!  Just use:

		**************************************
		gps find [search-term] [flag] [option]
		**************************************

		You will be able to install the packages directly from the terminal once provided
		the search results.
		`,
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.whatev.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
