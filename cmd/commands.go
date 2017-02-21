package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Everlag/gothing/stash"
)

// RootCmd is the root command...
var RootCmd = &cobra.Command{
	Use:   "thing",
	Short: "run the thing",
	Long:  "run the thing and this is supposed to be helpful D:",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hell yeah boi")
	},
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "get the latest stash update",
	Long:  "get the latest stash update, deserialize, and serialize to our testing format",
	Run: func(cmd *cobra.Command, args []string) {
		err := stash.FetchAndSetStore()
		if err != nil {
			fmt.Printf("failed to fetch and update stash data, err=%s\n", err)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(fetchCmd)
}
