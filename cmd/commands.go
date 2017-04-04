package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"unsafe"

	"encoding/hex"

	"strconv"

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
		cStashes, cItems, err := db.StashStashToCompact(resp.Stashes, bdb)
		if err != nil {
			fmt.Printf("failed to convert fat stashes to compact, err=%s\n", err)
			return
		}

		stats, err := db.AddStashes(cStashes, cItems, bdb)
		if err != nil {
			fmt.Printf("failed to store stashes, err=%s\n", err)
			return
		}

		var itemCount int
		itemCount, err = db.ItemStoreCount(bdb)
		if err != nil {
			fmt.Printf("failed to get item count, err=%s\n", err)
			return
		}

		var indexCount int
		indexCount, err = db.IndexEntryCount(bdb)
		if err != nil {
			fmt.Printf("failed to get index entry count, err=%s\n", err)
			return
		}

		representativeItem := db.Item{}
		serialItemSize := representativeItem.Msgsize()

		fmt.Printf(`done,
%s
%d items total, itemSize=%d bytes
%d index entries
`,
			stats, itemCount, serialItemSize, indexCount)
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

		// Decode and validate identifier
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

		// Grab it
		item, err := db.GetItemByIDGlobal(id, bdb)
		if err != nil {
			fmt.Printf("failed to find item, err=%s", err)
			return
		}

		// Inflate it
		inflated := item.Inflate(bdb)

		// Serialize so we can read the entire damn thing
		inflatedBytes, err := inflated.MarshalJSON()
		if err != nil {
			fmt.Printf("failed to marshal item item, err=%s", err)
			return
		}

		fmt.Printf("got item:\n%s\n", string(inflatedBytes))
	},
}

var lookupStringCmd = &cobra.Command{
	Use:   "string [\"StringHeapID\"]",
	Short: "lookup a string on the heap with a specific id(hex encoded)",
	Long:  "get the database and look for the string",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("please provide id to lookup")
			return
		}
		idString := args[0]

		// Decode and validate identifier
		idBytes, err := hex.DecodeString(idString)
		if err != nil {
			fmt.Printf("failed to decode id, err=%s\n", err)
			return
		}
		id := db.StringHeapIDFromBytes(idBytes)
		fmt.Printf("id decoded as %d\n", id)

		fmt.Printf("resolved: '%s'\n", id.Inflate(bdb))
	},
}

var lookupStringIDCmd = &cobra.Command{
	Use:   "stringid [\"string\"]",
	Short: "lookup the StringHeapID for a provided string",
	Long:  "get the database and look for the StringHeapID",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("please provide id to lookup")
			return
		}
		request := args[0]

		ids, err := db.GetStrings([]string{request}, bdb)
		if err != nil {
			fmt.Printf("failed to fetch string, err=%s\n", err)
			return
		}
		if len(ids) != 1 {
			fmt.Printf("id fetch failed for some reason D:")
			return
		}

		hexID := hex.EncodeToString(ids[0].ToBytes())
		fmt.Printf("resolved: '%s'\n", hexID)
	},
}

var searchItemByModCmd = &cobra.Command{
	Use:     "searchMinMod [\"maxMatches root flavor mod minValue\"]",
	Short:   "Find a items matching criteria up to maxMatches",
	Long:    "Search for an item with a given root and flavor item types with specific mod with minimum value",
	Example: "searchMinMod 3 Armour Boots \"\\\"#% increased Movement Speed\\\"\" 10",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 4 {
			fmt.Printf("invalid use, ex: %s\n", cmd.Example)
			return
		}
		// Ignore the arguments for a minute...
		maxMatchesString := args[0]
		root := args[1]
		flavor := args[2]
		mod := args[3]
		modMinValueString := args[4]

		maxMatches, err := strconv.Atoi(maxMatchesString)
		if err != nil {
			fmt.Printf("cannot read maxMatches '%s' as a number", maxMatchesString)
			return
		}
		minModValue, err := strconv.Atoi(modMinValueString)
		if err != nil {
			fmt.Printf("cannot read minValue '%s' as a number", modMinValueString)
			return
		}

		// Lookup the root, flavor, and mod
		// TODO not hardcode
		strings := []string{
			root,
			flavor,
			mod,
		}
		ids, err := db.GetStrings(strings, bdb)
		if err != nil {
			fmt.Printf("failed to fetch string, err=%s\n", err)
			return
		}
		// And we we need to fetch the league
		// TODO not hardcode
		leagueIDs, err := db.GetLeagues([]string{"Standard"}, bdb)
		if err != nil {
			fmt.Printf("failed to fetch league, err=%s\n", err)
			return
		}

		// OH, this is ugly D:
		query := db.NewIndexQuery(ids[0], ids[1],
			[]db.StringHeapID{ids[2]}, []uint16{uint16(minModValue)},
			leagueIDs[0], maxMatches)
		resultIDs, err := query.Run(bdb)
		if err != nil {
			fmt.Printf("failed to search items, err=%s\n", err)
			return
		}

		fmt.Println("result:")
		for _, id := range resultIDs {
			fmt.Printf("    %x\n", id)
		}
	},
}

var searchItemMultiMod = &cobra.Command{
	Use:     "searchMultiMod [\"path to MultiModSearch json\"]",
	Short:   "Find an item with types and mods",
	Example: "searchMultiMod ./query.json",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Printf("invalid use, ex: %s\n", cmd.Example)
			return
		}
		search, err := FetchMultiModSearch(args[0])
		if err != nil {
			fmt.Printf("failed to get search, err=%s\n", err)
			return
		}

		if len(search.Mods) == 0 {
			fmt.Println("no mods provided")
			return
		}

		if len(search.MinValues) != len(search.Mods) {
			fmt.Println("each mod must have a minvalue")
			return
		}

		// Lookup the root, flavor, and mod
		strings := []string{search.RootType, search.RootFlavor}
		ids, err := db.GetStrings(strings, bdb)
		if err != nil {
			fmt.Printf("failed to fetch rootType or RootFlavor id, err=%s\n", err)
			return
		}
		modIds, err := db.GetStrings(search.Mods, bdb)
		if err != nil {
			fmt.Printf("failed to fetch mod id, err=%s\n", err)
			return
		}

		// And we we need to fetch the league
		leagueIDs, err := db.GetLeagues([]string{search.League}, bdb)
		if err != nil {
			fmt.Printf("failed to fetch league, err=%s\n", err)
			return
		}

		// OH, this is ugly D:
		query := db.NewIndexQuery(ids[0], ids[1],
			modIds, search.MinValues, leagueIDs[0], search.MaxDesired)
		resultIDs, err := query.Run(bdb)
		if err != nil {
			fmt.Printf("failed to search items, err=%s\n", err)
			return
		}

		fmt.Println("result:")
		for _, id := range resultIDs {
			fmt.Printf("    %x\n", id)
		}
	},
}

var searchItemMultiModSlow = &cobra.Command{
	Use:     "searchMultiModSlow [\"path to MultiModSearch json\"]",
	Short:   "Find an item with types and mods. This is very slow but accurate",
	Example: "searchMultiModSlow ./query.json",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Printf("invalid use, ex: %s\n", cmd.Example)
			return
		}
		search, err := FetchMultiModSearch(args[0])
		if err != nil {
			fmt.Printf("failed to get search, err=%s\n", err)
			return
		}

		if len(search.Mods) == 0 {
			fmt.Println("no mods provided")
			return
		}

		if len(search.MinValues) != len(search.Mods) {
			fmt.Println("each mod must have a minvalue")
			return
		}

		// Lookup the root, flavor, and mod
		strings := []string{search.RootType, search.RootFlavor}
		ids, err := db.GetStrings(strings, bdb)
		if err != nil {
			fmt.Printf("failed to fetch rootType or RootFlavor id, err=%s\n", err)
			return
		}
		modIds, err := db.GetStrings(search.Mods, bdb)
		if err != nil {
			fmt.Printf("failed to fetch mod id, err=%s\n", err)
			return
		}

		// And we we need to fetch the league
		leagueIDs, err := db.GetLeagues([]string{search.League}, bdb)
		if err != nil {
			fmt.Printf("failed to fetch league, err=%s\n", err)
			return
		}

		// OH, this is ugly D:
		query := db.NewItemStoreQuery(ids[0], ids[1],
			modIds, search.MinValues, leagueIDs[0], search.MaxDesired)
		resultIDs, err := query.Run(bdb)
		if err != nil {
			fmt.Printf("failed to search items, err=%s\n", err)
			return
		}

		fmt.Println("result:")
		for _, id := range resultIDs {
			fmt.Printf("    %x\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(tryCompactyCmd)
	rootCmd.AddCommand(storeItemsCmd)
	rootCmd.AddCommand(listLeaguesCmd)
	rootCmd.AddCommand(lookupItemCmd)
	rootCmd.AddCommand(lookupStringCmd)
	rootCmd.AddCommand(lookupStringIDCmd)
	rootCmd.AddCommand(searchItemByModCmd)
	rootCmd.AddCommand(searchItemMultiMod)
	rootCmd.AddCommand(searchItemMultiModSlow)
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
