package stash

// We use messagepack for serialzing ChangeSet
// as it requires storing large amounts of binary data, which
// messagepack handles significantly easier than json
//go:generate msgp

import (
	"bytes"
	"fmt"
	"os"

	"github.com/mailru/easyjson"
	"github.com/tinylib/msgp/msgp"
	"github.com/ulikunitz/xz"
)

// ChangeSet represents the number
type ChangeSet struct {
	// Changes as individually compressed
	Changes []CompressedResponse
	// Size as stored
	Size int
}

// AddResponse includes another CompressedResponse inside the Changes
func (changes *ChangeSet) AddResponse(comp CompressedResponse) {
	changes.Changes = append(changes.Changes, comp)
	changes.Size += comp.Size
}

// Save stores the marshalled ChangeSet's at the provided location
func (changes *ChangeSet) Save(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return fmt.Errorf("failed to open file, err=%s", err)
	}
	err = msgp.WriteFile(changes, f)
	if err != nil {
		return fmt.Errorf("failed to encode and write file, err=%s", err)
	}
	return nil
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

	// Setup decompression reader
	buf := bytes.NewBuffer(comp.Content)
	uncomp, err := xz.NewReader(buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create xz reader, err=%s", err)
	}

	// And unmarshal through the decompressor
	var resp Response
	err = easyjson.UnmarshalFromReader(uncomp, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to easyjson unmarshal, err=%s", err)
	}

	return &resp, nil
}

// NewCompressedResponse compresses a Response into compact form
// and returns the result
func NewCompressedResponse(resp *Response) (*CompressedResponse, error) {
	var buf bytes.Buffer
	comp, err := xz.NewWriter(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create xz writer, err=%s", err)
	}

	_, err = easyjson.MarshalToWriter(resp, comp)
	if err != nil {
		return nil, fmt.Errorf("failed easyjson marshal, err=%s", err)
	}

	return &CompressedResponse{
		Content: buf.Bytes(),
		Size:    buf.Len(),
	}, nil

}
