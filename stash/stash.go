package stash

// JSON for GGG responses
// and msgpack for faster testing setups
//go:generate easyjson $GOFILE
//go:generate msgp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"strings"

	"github.com/mailru/easyjson"
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
	Values   []uint16
}

// MarshalJSON implements custom serialization for this type
func (m *ItemMod) MarshalJSON() ([]byte, error) {
	result := []byte(m.Template)

	for _, value := range m.Values {

		loc := hashRegex.FindIndex(result)
		if loc == nil {
			return nil, fmt.Errorf("failed to match template string slot with expected value")
		}

		insert := []byte(strconv.FormatUint(uint64(value), 10))

		portion := hashRegex.ReplaceAll(result[:loc[1]], insert)
		result = append(portion, result[loc[1]:]...)
	}

	return json.Marshal(string(result))
}

// UnmarshalJSON implements custom deserialization for this type
//
// Typically, the GGG api will return string which we run through
// regexp for ease
func (m *ItemMod) UnmarshalJSON(b []byte) error {
	// Find magnitudes in the mod
	matches := numberRegex.FindAll(b, -1)

	// Grab the template string representing the mod
	template := string(numberRegex.ReplaceAll(b, []byte("#")))

	// Convert the matches to numbers
	m.Values = make([]uint16, len(matches))
	for i, match := range matches {
		parsed, err := strconv.ParseUint(string(match), 10, 16)
		if err != nil {
			return fmt.Errorf("failed to convert match to int, err=%s", err)
		}
		m.Values[i] = uint16(parsed)
	}

	// If enclosed by quotes, get rid of them.
	//
	// This theoretically makes boltdb searches faster but
	// it just improves my quality of life...
	if len(template) > 2 &&
		strings.HasPrefix(template, "\"") &&
		strings.HasSuffix(template, "\"") {

		template = strings.TrimPrefix(template, "\"")
		template = strings.TrimSuffix(template, "\"")
	}

	m.Template = []byte(template)

	return nil
}

// Item represents a single item found from the stash api
//easyjson:json
type Item struct {
	Verified     bool      `json:"verified"`
	League       string    `json:"league"`
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	TypeLine     string    `json:"typeLine"`
	Identified   bool      `json:"identified"`
	Corrupted    bool      `json:"corrupted"`
	ImplicitMods []ItemMod `json:"implicitMods,omitempty"`
	ExplicitMods []ItemMod `json:"explicitMods,omitempty"`
	Note         string    `json:"note,omitempty"`
	UtilityMods  []string  `json:"utilityMods,omitempty"`
	DescrText    string    `json:"descrText,omitempty"`
	// Additional data not present in response
	StashID    string `json:"-"`
	RootType   string `json:"-"`
	RootFlavor string `json:"-"`
}

// GetMods concats both ExplicitMods and ImplicitMods, returning the result
func (item *Item) GetMods() []ItemMod {
	mods := make([]ItemMod, 0)
	mods = append(mods, item.ImplicitMods...)
	mods = append(mods, item.ExplicitMods...)
	return mods
}

// Stash represents a stash tab with items and associated metadata
//easyjson:json
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
//easyjson:json
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

	serial, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to ReadlAll from file")
	}

	return RespFromJSON(serial)
}

// CleanResponse adds on Type data as well as ensures the response
// will satisfy our expectations, as wildly unreasonable as they can be
func CleanResponse(response *Response) error {
	for _, stash := range response.Stashes {
		for i, item := range stash.Items {
			item.StashID = stash.ID
			// Handle cases where the item is identified purely by its typeline
			if len(item.Name) == 0 {
				item.Name = item.TypeLine
			}

			if len(item.Note) == 0 {
				item.Note = "unknown"
			}

			// Resolve the typeLine on an item to its flavor and root
			flavor, root, ok := MatchTypeline(item.TypeLine)
			if !ok {
				fmt.Println(stash.ID)
				fmt.Println(item)
				return fmt.Errorf("failed to discover flavor and root for Item, name=%s, typeline=%s, id=%s",
					item.Name, item.TypeLine, item.ID)
			}
			item.RootType = root
			item.RootFlavor = flavor

			stash.Items[i] = item
		}
	}

	return nil
}

// RespFromJSON attempts to deserialize the provided data
// and return it as a StashResponse
func RespFromJSON(serial []byte) (*Response, error) {
	var response Response
	err := easyjson.Unmarshal(serial, &response)
	// err = decoder.Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode TestResponseLoc, err=%s", err)
	}

	return &response, CleanResponse(&response)
}

// FetchUpdate grabs the update indicated by the changeID.
//
// If empty changeID is provided, it grabs the default update.
func FetchUpdate(changeID string) (*Response, error) {
	endpoint := StashAPIBase
	if changeID != "" {
		endpoint = fmt.Sprintf("%s?id=%s", StashAPIBase, changeID)
	}

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil,
			fmt.Errorf("failed to call stash api, err=%s", err)
	}
	defer resp.Body.Close()

	var response Response
	err = easyjson.UnmarshalFromReader(resp.Body, &response)
	if err != nil {
		return nil,
			fmt.Errorf("failed to decode stash tab response, err=%s", err)
	}

	return &response, CleanResponse(&response)
}

// FetchAndSetStore grabs the latest stash tab api update
// and stores it in TestResponseLoc
func FetchAndSetStore() error {

	response, err := FetchUpdate("")
	if err != nil {
		return fmt.Errorf("failed to fetch update, err=%s", err)
	}

	serial, err := response.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal stash tab response, err=%s", err)
	}

	if err = ioutil.WriteFile(TestResponseLoc, serial, 0777); err != nil {
		return fmt.Errorf("failed to save stashResult.json, err=%s", err)
	}

	return nil
}
