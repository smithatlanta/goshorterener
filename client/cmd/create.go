package cmd

import (
	"fmt"

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
	var url = args[0]
	var desiredShortcut = args[1]
	shortenedURL, err := Create(url, desiredShortcut)
	if err != nil {
		fmt.Println("Create Failed")
	}
	println(shortenedURL)
	return
}

//Create -
func Create(url string, desiredShortcut string) (string, error) {
	CreateShortenedURL(desiredShortcut, url)
	return "", nil
}
