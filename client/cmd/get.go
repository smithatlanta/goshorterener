package cmd

import (
	"fmt"

	"github.com/smithatlanta/goshortener/lib"
	"github.com/spf13/cobra"
)

// getCmd represents the create command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a shortened url",
	Long:  `The get command prompts for your shortcut and returns a the full url.`,
	Run:   get,
}

func init() {
	RootCmd.AddCommand(getCmd)
}

func get(cmd *cobra.Command, args []string) {
	var shortcut = args[0]
	var projectID = args[1]
	fullURL, err := Get(shortcut, projectID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fullURL)
	return
}

//Get -
func Get(shortcut string, projectID string) (string, error) {
	url, err := lib.GetShortenedURL(shortcut, projectID)
	return url, err
}
