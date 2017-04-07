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

	set := stash.ChangeSet{}

	var comp *stash.CompressedResponse
	var err error
	for set.Size < sizeLimit {
		fmt.Printf("id=%s fetching\n", changeID)
		changeID, comp, err = FetchAndCompress(changeID)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch and compress, err=%s", err)
		}

		set.AddResponse(*comp)
		fmt.Printf("fetched size=%d; total size=%d\n", comp.Size, set.Size)

		// Wait so not DOSing
		time.Sleep(WaitDuration)
	}

	return &set, nil
}

// WaitDuration is the time spent between update requests
//
// This is required to not get rate limited...
const WaitDuration = time.Second * 2

func main() {

	size := flag.Int("maxsize", -1, "Maximum stored size in bytes")
	path := flag.String("dest", "", "location for content")
	changeID := flag.String("id", "", "starting changeID(optional)")

	flag.Parse()

	if *size == -1 {
		fmt.Println("please provide maxsize")
		os.Exit(1)
	}

	if *path == "" {
		fmt.Println("please provide dest(ination)")
		os.Exit(1)
	}

	if *changeID == "" {
		fmt.Println("empty id, using default changeID")
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
