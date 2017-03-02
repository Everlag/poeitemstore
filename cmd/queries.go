package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

// MultiModSearch specifies a search to perform for items
type MultiModSearch struct {
	MaxDesired int
	RootType   string
	RootFlavor string
	League     string
	Mods       []string
	MinValues  []uint16
}

// FetchMultiModSearch returns a MultiModSearch deserialized
// from the provided path on disk
func FetchMultiModSearch(path string) (*MultiModSearch, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file, err=%s", err)
	}
	decoder := json.NewDecoder(f)
	var search MultiModSearch
	err = decoder.Decode(&search)
	if err != nil {
		return nil, fmt.Errorf("failed to read query, err=%s", err)
	}

	return &search, nil
}
