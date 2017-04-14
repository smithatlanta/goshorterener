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
	fmt.Println(shortcut)
	fmt.Println(projectID)
	fullURL, err := Get(shortcut, projectID)
	if err != nil {
		fmt.Println(err)
	}
	println(fullURL)
	return
}

//Get -
func Get(shortcut string, projectID string) (string, error) {
	url, err := lib.GetShortenedURL(shortcut, projectID)
	fmt.Printf(url)
	fmt.Println(err)
	return url, err
}
