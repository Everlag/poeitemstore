package pool

import (
	"fmt"
	"sync"
)

// Stats represents usage statistics for a Pool.
type Stats struct {
	Makes               int // How many freshly initialized buffers in Borrows
	Borrows             int
	Gives               int
	TotalLengthMaked    int
	TotalLengthBorrowed int
	TotalLengthGiven    int
}

func (s Stats) String() string {
	return fmt.Sprintf(`%d Makes | %d Borrows | %d Gives
%d make items | %d borrowed items | %d given items
`,
		s.Makes, s.Borrows, s.Gives,
		s.TotalLengthMaked, s.TotalLengthBorrowed, s.TotalLengthGiven)
}

// The minimum size for an entry in the pool
var poolMinEntrySize = 10

// MakerCB is a function that generates a slice
// of minLen of the desired type when the pool
// doesn't have any buffers cached.
type MakerCB func(minLen int) interface{}

// ClearCB clears the state of the provided buffer.
// This returns the cleared buffer.
//
// This is necessary as 'buf.([]interface{})[:0]' will always
// fail as a result of []interface having a memory layout different from
// the underlying type of buf.
type ClearCB func(buf interface{}) interface{}

// CapCB returns the capacity of a given buffer.
//
// I wish we had generics D:
type CapCB func(buf interface{}) int

// Pool represents a thread-safe pool of reusable
// slices to reduce allocations.
//
// NOTE: Only slices are handled by this pool.
// Maps have a high clearing cost, so we don't pool them.
type Pool struct {
	bufs     chan interface{}
	maker    MakerCB
	clearer  ClearCB
	capacity CapCB

	stats Stats
	*sync.Mutex
}

// NewPool creates a properly initialized IDPool
func NewPool(maxSize int,
	maker MakerCB, clearer ClearCB, capacity CapCB) Pool {
	p := Pool{
		bufs:     make(chan interface{}, maxSize),
		maker:    maker,
		clearer:  clearer,
		capacity: capacity,
	}
	p.Mutex = &sync.Mutex{}
	return p
}

// Borrow returns a pre-allocated buffer in the pool
// or a fresh buffer if none are available
func (p *Pool) Borrow(minLen int) interface{} {
	p.Lock()
	defer p.Unlock()

	var buf interface{}
	select {
	case buf = <-p.bufs:
	default:
		if minLen < poolMinEntrySize {
			minLen = poolMinEntrySize
		}
		buf = p.maker(minLen)
		p.stats.Makes++
		p.stats.TotalLengthMaked += minLen
	}

	// Clear the buffer as its returned
	buf = p.clearer(buf)

	p.stats.Borrows++
	p.stats.TotalLengthBorrowed += p.capacity(buf)

	return buf
}

// Give returns a buffer to the set
//
// If we have more buffers than we are configured to handle,
// further provided buffers are tossed out.
func (p *Pool) Give(buf interface{}) {
	p.Lock()
	defer p.Unlock()

	select {
	case p.bufs <- buf:
	default:
		// Let the buf reference fall out of scope
		// so the GC takes care of it
	}

	p.stats.Gives++
	p.stats.TotalLengthGiven += p.capacity(buf)

}

// Stats returns the usage statistics for the Pool
func (p *Pool) Stats() Stats {
	p.Lock()
	defer p.Unlock()
	return p.stats
}
