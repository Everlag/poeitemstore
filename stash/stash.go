package stash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Item represents a single item found from the stash api
type Item struct {
	Verified     bool     `json:"verified"`
	Ilvl         int      `json:"ilvl"`
	Icon         string   `json:"icon"`
	League       string   `json:"league"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	TypeLine     string   `json:"typeLine"`
	Identified   bool     `json:"identified"`
	Corrupted    bool     `json:"corrupted"`
	ImplicitMods []string `json:"implicitMods,omitempty"`
	ExplicitMods []string `json:"explicitMods,omitempty"`
	FlavourText  []string `json:"flavourText,omitempty"`
	Note         string   `json:"note,omitempty"`
	Properties   []struct {
		Name        string        `json:"name"`
		Values      []interface{} `json:"values"`
		DisplayMode int           `json:"displayMode"`
	} `json:"properties,omitempty"`
	UtilityMods []string `json:"utilityMods,omitempty"`
	DescrText   string   `json:"descrText,omitempty"`
}

// Response represents expected structure of a stash api call
type Response struct {
	NextChangeID string `json:"next_change_id"`
	Stashes      []struct {
		AccountName       string `json:"accountName"`
		LastCharacterName string `json:"lastCharacterName"`
		ID                string `json:"id"`
		Stash             string `json:"stash"`
		StashType         string `json:"stashType"`
		Items             []Item `json:"items"`
		Public            bool   `json:"public"`
	} `json:"stashes"`
}

// StashAPIBase is the URL the stash api is located at
const StashAPIBase string = "http://www.pathofexile.com/api/public-stash-tabs"

// TestResponseLoc is where testing data is kept
const TestResponseLoc string = "StashResponse.json"

// GetStored gets the stored testing data and returns it
func GetStored() (*Response, error) {
	f, err := os.Open(TestResponseLoc)
	if err != nil {
		return nil, fmt.Errorf("failed to open TestResponseLoc, err=%s", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	var response Response
	err = decoder.Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TestResponseLoc, err=%s", err)
	}

	return &response, nil
}

// FetchAndSetStore grabs the latest stash tab api update
// and stores it in TestResponseLoc
func FetchAndSetStore() error {
	resp, err := http.Get(StashAPIBase)
	if err != nil {
		return fmt.Errorf("failed to call stash api, err=%s", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var response Response
	err = decoder.Decode(&response)
	if err != nil {
		return fmt.Errorf("failed to decode stash tab response, err=%s", err)
	}

	serial, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to marshal stash tab response, err=%s", err)
	}

	if err = ioutil.WriteFile(TestResponseLoc, serial, 0777); err != nil {
		return fmt.Errorf("failed to save stashResult.json, err=%s", err)
	}

	return nil
}
