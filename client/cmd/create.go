package cmd

import (
	"fmt"

	"github.com/smithatlanta/goshortener/lib"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a shortened url",
	Long:  `The create command prompts for your url and returns a new shortened url.`,
	Run:   create,
}

func init() {
	RootCmd.AddCommand(createCmd)
}

func create(cmd *cobra.Command, args []string) {
	var desiredShortcut = args[0]
	var url = args[1]
	var projectID = args[2]
	err := Create(desiredShortcut, url, projectID)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//Create -
func Create(desiredShortcut string, url string, projectID string) error {
	err := lib.CreateShortenedURL(desiredShortcut, url, projectID)
	return err
}
