package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "urlshortener",
	Short: "shorten a url",
	Long:  ``,
}

// Version is the version of this app
var Version string

// BuildDate is the date this binary was built
var BuildDate string

// Verbose determines whether or not verbose output is enabled
var Verbose bool

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, buildDate string) {
	Version = version
	BuildDate = buildDate

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Show more output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
