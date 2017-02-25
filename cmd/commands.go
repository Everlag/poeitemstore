package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"unsafe"

	"encoding/hex"

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
			return
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
			return
		}
		fmt.Printf("read cached stash update, %d entries found\n", len(resp.Stashes))
	},
}

var tryCompactyCmd = &cobra.Command{
	Use:   "tryCompact",
	Short: "attempt to compact all stashes in cached stash update",
	Long:  "get the stash update from disk, deserialize it, and try compacting stashes, this will result in db writes",
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := stash.GetStored()
		if err != nil {
			fmt.Printf("failed to read cached stash data, err=%s\n", err)
			return
		}

		// Flatten the items
		_, cItems, err := db.StashStashToCompact(resp.Stashes, bdb)
		if err != nil {
			fmt.Printf("failed to convert fat stashes to compact, err=%s\n", err)
			return
		}
		compacItemtSize := unsafe.Sizeof(db.Item{})
		fmt.Printf("compact done, item size is %d bytes\n", int(compacItemtSize)*len(cItems))
	},
}

var storeItemsCmd = &cobra.Command{
	Use:   "storeItems",
	Short: "attempt to store all items in cached stash update",
	Long:  "get the stash update from disk, deserialize it, compact it, and write it to the db",
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := stash.GetStored()
		if err != nil {
			fmt.Printf("failed to read cached stash data, err=%s\n", err)
			return
		}

		// Flatten the items
		_, cItems, err := db.StashStashToCompact(resp.Stashes, bdb)
		if err != nil {
			fmt.Printf("failed to convert fat stashes to compact, err=%s\n", err)
			return
		}

		overwritten, err := db.AddItems(cItems, bdb)
		if err != nil {
			fmt.Printf("failed to store items, err=%s\n", err)
			return
		}

		var count int
		count, err = db.ItemStoreCount(bdb)
		if err != nil {
			fmt.Printf("failed to get item count, err=%s\n", err)
			return
		}

		representativeItem := db.Item{}
		serialItemSize := representativeItem.Msgsize()

		fmt.Printf("items stored done, %d items added, %d item overwritten, %d items total, itemSize=%d bytes\n",
			len(cItems), overwritten, count, serialItemSize)
	},
}

var listLeaguesCmd = &cobra.Command{
	Use:   "listLeagues",
	Short: "attempt to fetch the names of stored leagues",
	Long:  "open the db and find all leagues",
	Run: func(cmd *cobra.Command, args []string) {

		leagues, err := db.ListLeagues(bdb)
		if err != nil {
			fmt.Printf("failed to read cached stash data, err=%s\n", err)
			return

		}

		fmt.Printf("found leagues '%v'\n", leagues)
	},
}

var lookupItemCmd = &cobra.Command{
	Use:   "lookup [\"itemid\"]",
	Short: "lookup an item with a specific id",
	Long:  "get the database and lookup an item with our short, hashed format. This searches every league",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("please provide id to lookup")
			return
		}
		idString := args[0]

		idBytes, err := hex.DecodeString(idString)
		if err != nil {
			fmt.Printf("failed to decode id, err=%s\n", err)
			return
		}
		if len(idBytes) != db.IDSize {
			fmt.Printf("id wrong decoded size; got %d, expected %d\n", db.IDSize, len(idBytes))
			return
		}
		var id db.ID
		copy(id[:], idBytes)

		item, err := db.GetItemByIDGlobal(id, bdb)
		if err != nil {
			fmt.Printf("failed to find item, err=%s", err)
		}

		fmt.Println(item)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(tryCompactyCmd)
	rootCmd.AddCommand(storeItemsCmd)
	rootCmd.AddCommand(listLeaguesCmd)
	rootCmd.AddCommand(lookupItemCmd)
}

// HandleCommands runs commands after setting up
// necessary preconditions
func HandleCommands(db *bolt.DB) {
	bdb = db

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("command failed, err=%s", err)
		return
	}
}
