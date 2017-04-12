package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

func (search *MultiModSearch) String() string {
	modPrints := make([]string, len(search.Mods))
	var modString string
	if len(search.Mods) != len(search.MinValues) {
		modString = "invalid mods: len(Mods) != len(MinValues)"
	} else {
		for i, mod := range search.Mods {
			modPrints[i] = fmt.Sprintf("%s: %d", mod, search.MinValues[i])
		}
		if len(modPrints) == 0 {
			modString = "no mods present"
		}
		modString = strings.Join(modPrints, "\n")
	}
	return fmt.Sprintf(`RootType: %s, RootFlavor: %s,
League: %s, MaxDesired: %d
%s`,
		search.RootType, search.RootFlavor,
		search.League, search.MaxDesired, modString)
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
