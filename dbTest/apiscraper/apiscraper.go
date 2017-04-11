package main

import (
	"fmt"

	"os"

	"flag"

	"time"

	"github.com/Everlag/poeitemstore/stash"
)

// FetchAndCompress fetches a given changeID and returns its compressed
// representation alongside its NextChangeID field
func FetchAndCompress(changeID string) (string,
	*stash.CompressedResponse, error) {
	response, err := stash.FetchUpdate(changeID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to fetch update, err=%s", err)
	}

	comp, err := stash.NewCompressedResponse(response)
	return response.NextChangeID, comp, err
}

// FetchTillLimit grabs and saves stash updates starting from
// the provided changeID until the storage space used exceeds
// the provided limit in terms of bytes.
func FetchTillLimit(changeID string,
	sizeLimit int) (*stash.ChangeSet, error) {

	set := stash.NewChangeSet()

	var comp *stash.CompressedResponse
	var err error
	for set.Size < sizeLimit {
		fmt.Printf("id=%s fetching\n", changeID)
		changeID, comp, err = FetchAndCompress(changeID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch and compress, err=%s", err)
		}

		set.AddResponse(changeID, *comp)
		fmt.Printf("fetched size=%d; total size=%d\n", comp.Size, set.Size)

		// Wait so not DOSing
		time.Sleep(WaitDuration)
	}

	return &set, nil
}

// Unpack extracts the given Response with the provided changeID
// from the ChangeSet at the specific path
func Unpack(changeID, path string) (*stash.Response, error) {

	set, err := stash.OpenChangeSet(path)
	if err != nil {
		return nil,
			fmt.Errorf("failed to open changeset, err=%s", err)
	}

	comp, ok := set.GetCompByChangeID(changeID)
	if !ok {
		return nil, fmt.Errorf("failed to find response with corresponding id")
	}

	resp, err := comp.Decompress()
	if err != nil {
		return nil, fmt.Errorf("failed to decompress change, err=%s", err)
	}

	return resp, nil

}

// SaveResponse saves the JSON of the provided Response to
// the CWD  as `$changeID.json`.
func SaveResponse(changeID string, resp stash.Response) error {
	json, err := resp.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal response, err=%s", err)
	}

	fileName := fmt.Sprintf("%s.json", changeID)

	f, err := os.OpenFile(fileName, os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("failed to open file, err=%s", err)
	}
	defer f.Close()

	_, err = f.Write(json)
	if err != nil {
		return fmt.Errorf("failed to write file, err=%s", err)
	}
	return nil
}

// WaitDuration is the time spent between update requests
//
// This is required to not get rate limited...
const WaitDuration = time.Second * 4

func main() {

	size := flag.Int("maxsize", -1, "Maximum stored size in bytes")
	path := flag.String("dest", "", "location for content")
	changeID := flag.String("id", "", "starting changeID(optional)")
	unpackWhich := flag.Bool("unpack", false, "changeID to extract(optional)")

	flag.Parse()

	if *path == "" {
		fmt.Println("please provide dest(ination)")
		os.Exit(1)
	}

	if *changeID == "" {
		fmt.Println("empty id, using default changeID")
	}

	if *unpackWhich {
		fmt.Printf("unpacking changeid to %s.json\n", *changeID)
		resp, err := Unpack(*changeID, *path)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		err = SaveResponse(*changeID, *resp)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if *size == -1 {
		fmt.Println("please provide maxsize")
		os.Exit(1)
	}

	set, err := FetchTillLimit(*changeID, *size)
	if err != nil {
		fmt.Printf("encountered error while fetching, err=%s\n", err)
		os.Exit(1)
	}

	err = set.Save(*path)
	if err != nil {
		fmt.Printf("failed to save compressed set, err=%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("done, %d bytes fetched\n", set.Size)

}
