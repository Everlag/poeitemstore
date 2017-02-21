package stash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// PropertyValue holds a string value alongside an
// associated PrintKey
type PropertyValue struct {
	Value    string
	PrintKey int
}

// MarshalJSON implements custom serialization for this type
func (v *PropertyValue) MarshalJSON() ([]byte, error) {
	raw := make([]interface{}, 2)
	raw[0] = v.Value
	raw[1] = v.PrintKey

	return json.Marshal(raw)
}

// UnmarshalJSON implements custom deserialization for this type
//
// Typically, the GGG api will return [string, int] which is very unhelpful,
// so we take care of that right here
func (v *PropertyValue) UnmarshalJSON(b []byte) error {
	var raw []interface{}

	err := json.Unmarshal(b, &raw)
	if err != nil {
		return fmt.Errorf("invalid property pair, unparseable, err=%s", err)
	}

	if len(raw) != 2 {
		return fmt.Errorf("invalid property pair, (len()==%d)!=2", len(raw))
	}

	var ok bool
	v.Value, ok = raw[0].(string)
	if !ok {
		return fmt.Errorf("invalid property pair, first element not string")
	}

	// We have to use the widest possible type here and narrow...
	var broad float64
	broad, ok = raw[1].(float64)
	if !ok {
		return fmt.Errorf("invalid property pair, second element not float64")
	}
	v.PrintKey = int(broad)

	return nil
}

// Regexes we use for ItemMod parsing
var numberRegex = regexp.MustCompile(`\d{1,}`) // Grabbing digits
var hashRegex = regexp.MustCompile("#")        // Filling templates

// ItemMod is a modifier an item can have
type ItemMod struct {
	Template []byte
	Values   []int
}

// MarshalJSON implements custom serialization for this type
func (m *ItemMod) MarshalJSON() ([]byte, error) {
	result := []byte(m.Template)

	for _, value := range m.Values {

		loc := hashRegex.FindIndex(result)
		if loc == nil {
			return nil, fmt.Errorf("failed to match template string slot with expected value")
		}

		insert := []byte(strconv.Itoa(value))

		portion := hashRegex.ReplaceAll(result[:loc[1]], insert)
		result = append(portion, result[loc[1]:]...)
	}

	return result, nil
}

// UnmarshalJSON implements custom deserialization for this type
//
// Typically, the GGG api will return string which we run through
// regexp for ease
func (m *ItemMod) UnmarshalJSON(b []byte) error {
	// Find magnitudes in the mod
	matches := numberRegex.FindAll(b, -1)

	// Grab the template string representing the mod
	m.Template = numberRegex.ReplaceAll(b, []byte("#"))

	// Convert the matches to numbers
	m.Values = make([]int, len(matches))
	var err error
	for i, match := range matches {
		m.Values[i], err = strconv.Atoi(string(match))
		if err != nil {
			return fmt.Errorf("failed to convert match to int, err=%s", err)
		}
	}

	return nil
}

// Property represents a single item property
type Property struct {
	Name        string          `json:"name"`
	Values      []PropertyValue `json:"values"`
	DisplayMode int             `json:"displayMode"`
}

// Item represents a single item found from the stash api
type Item struct {
	Verified     bool       `json:"verified"`
	Ilvl         int        `json:"ilvl"`
	Icon         string     `json:"icon"`
	League       string     `json:"league"`
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	TypeLine     string     `json:"typeLine"`
	Identified   bool       `json:"identified"`
	Corrupted    bool       `json:"corrupted"`
	ImplicitMods []ItemMod  `json:"implicitMods,omitempty"`
	ExplicitMods []ItemMod  `json:"explicitMods,omitempty"`
	FlavourText  []string   `json:"flavourText,omitempty"`
	Note         string     `json:"note,omitempty"`
	Properties   []Property `json:"properties,omitempty"`
	UtilityMods  []string   `json:"utilityMods,omitempty"`
	DescrText    string     `json:"descrText,omitempty"`
	// Additional data not present in response
	StashID    string `json:"-"`
	RootType   string `json:"-"`
	RootFlavor string `json:"-"`
}

// Stash represents a stash tab with items and associated metadata
type Stash struct {
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	ID                string `json:"id"`
	Stash             string `json:"stash"`
	StashType         string `json:"stashType"`
	Items             []Item `json:"items"`
	Public            bool   `json:"public"`
}

// Response represents expected structure of a stash api call
type Response struct {
	NextChangeID string  `json:"next_change_id"`
	Stashes      []Stash `json:"stashes"`
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

	// Add additioanl data to items
	for _, stash := range response.Stashes {
		for _, item := range stash.Items {
			item.StashID = stash.ID
			flavor, root, ok := MatchBase(item.TypeLine)
			if !ok {
				item.RootType = item.TypeLine
				item.RootFlavor = item.TypeLine
			} else {
				item.RootType = root
				item.RootFlavor = flavor
			}
		}
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
