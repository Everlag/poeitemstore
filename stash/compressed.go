package stash

// We use messagepack for serialzing ChangeSet
// as it requires storing large amounts of binary data, which
// messagepack handles significantly easier than json
//go:generate msgp

import (
	"bytes"
	"fmt"
	"os"

	"github.com/golang/snappy"
	"github.com/tinylib/msgp/msgp"
)

// ChangeSet represents the number
type ChangeSet struct {
	// Changes mapping from ID to position in Changes
	ChangeIDToIndex map[string]int
	// Changes as individually compressed
	Changes []CompressedResponse
	// Size as stored
	Size int
}

// NewChangeSet returns an empty, initialized ChangeSet
func NewChangeSet() ChangeSet {
	return ChangeSet{
		ChangeIDToIndex: make(map[string]int),
		Changes:         make([]CompressedResponse, 0),
		Size:            0,
	}
}

// AddResponse includes another CompressedResponse inside the Changes
func (changes *ChangeSet) AddResponse(changeID string,
	comp CompressedResponse) {
	changes.Changes = append(changes.Changes, comp)
	changes.Size += comp.Size
	changes.ChangeIDToIndex[changeID] = len(changes.Changes) - 1
}

// Save stores the marshalled ChangeSet's at the provided location
func (changes *ChangeSet) Save(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("failed to open file, err=%s", err)
	}
	defer f.Close()
	err = msgp.WriteFile(changes, f)
	if err != nil {
		return fmt.Errorf("failed to encode and write file, err=%s", err)
	}
	return nil
}

// GetFirstChange returns the first Change in the set.
func (changes *ChangeSet) GetFirstChange() (*CompressedResponse, error) {
	if len(changes.Changes) == 0 {
		return nil, fmt.Errorf("failed to get change")
	}
	return &changes.Changes[0], nil
}

// GetCompByChangeID returns the CompressedResponse associated with
// a given changeID. Follows the _, ok pattern ala maps if not found
func (changes *ChangeSet) GetCompByChangeID(changeID string) (*CompressedResponse,
	bool) {

	index, ok := changes.ChangeIDToIndex[changeID]
	if !ok {
		return nil, false
	}

	return &changes.Changes[index], true

}

// OpenChangeSet attempts to open a ChangeSet at the provided path
func OpenChangeSet(path string) (*ChangeSet, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file, err=%s", err)
	}

	var set ChangeSet
	err = msgp.ReadFile(&set, f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file or unmarshal, err=%s", err)
	}

	return &set, nil
}

// CompressedResponse represents a Response compressed using this package
type CompressedResponse struct {
	Content []byte
	Size    int
}

// Decompress decompresses the CompressedResponse.
func (comp CompressedResponse) Decompress() (*Response, error) {

	// Stream the read
	buf := bytes.NewBuffer(comp.Content)
	decomp := snappy.NewReader(buf)

	var resp Response
	err := msgp.Decode(decomp, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to msgp decode, err=%s", err)
	}

	return &resp, CleanResponse(&resp)
}

// NewCompressedResponse compresses a Response into compact form
// and returns the result
func NewCompressedResponse(resp *Response) (*CompressedResponse, error) {
	var buf bytes.Buffer

	comp := snappy.NewBufferedWriter(&buf)

	msgpWriter := msgp.NewWriter(comp)
	err := resp.EncodeMsg(msgpWriter)
	if err != nil {
		return nil, fmt.Errorf("failed msgp marshal, err=%s", err)
	}

	if err := msgpWriter.Flush(); err != nil {
		return nil, fmt.Errorf("failed to flush msgp, err=%s", err)
	}

	if err := comp.Close(); err != nil {
		return nil, fmt.Errorf("failed to close snappy compressor, err=%s", err)
	}

	return &CompressedResponse{
		Content: buf.Bytes(),
		Size:    buf.Len(),
	}, nil

}
