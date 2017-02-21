package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Everlag/gothing/db"
	"github.com/Everlag/gothing/stash"
	"github.com/boltdb/bolt"
)

// db is a pointer to a database
// that should be valid on calling any command.
var bdb *bolt.DB

// rootCmd is the root command...
var rootCmd = &cobra.Command{
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

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check if cached stash update is valid",
	Long:  "get the stash update from disk and try to deserialize",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := stash.GetStored()
		if err != nil {
			fmt.Printf("failed to read cached stash data, err=%s\n", err)
			os.Exit(-1)
		}
		fmt.Printf("read cached stash update, %d entries found\n", len(resp.Stashes))
	},
}

var addNamesCmd = &cobra.Command{
	Use:   "addNames",
	Short: "add all names in cached stash update",
	Long:  "get the stash update from disk, deserialize it, and add all property names it contains to the database",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := stash.GetStored()
		if err != nil {
			fmt.Printf("failed to read cached stash data, err=%s\n", err)
			os.Exit(-1)
		}

		if err := db.AddPropertyNamesFromResponse(resp, bdb); err != nil {
			fmt.Printf("failed to add property names, err=%s\n", err)
			os.Exit(-1)
		}

		var count int
		count, err = db.PropertyNameCount(bdb)
		if err != nil {
			fmt.Printf("failed to get property name count, err=%s\n", err)
			os.Exit(-1)
		}

		fmt.Printf("added property names, %d properties exist\n", count)
	},
}

var lookupPropertyCmd = &cobra.Command{
	Use:   "lookupProperty [\"property to lookup\"]",
	Short: "lookup the integer identifier for a property",
	Long:  "get the database and lookup a property",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("please provide property")
			os.Exit(-1)
		}
		property := args[0]

		index, err := db.GetPropertyID(property, bdb)

		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(-1)
		}

		fmt.Printf("%s = %d\n", property, index)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(addNamesCmd)
	rootCmd.AddCommand(lookupPropertyCmd)
}

// HandleCommands runs commands after setting up
// necessary preconditions
func HandleCommands(db *bolt.DB) {
	bdb = db

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
