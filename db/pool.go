package db

import (
	"fmt"
	"sync"
)

// PoolStats represents usage statistics for a Pool.
type PoolStats struct {
	Makes               int // How many freshly initialized buffers in Borrows
	Borrows             int
	Gives               int
	TotalLengthBorrowed int
	TotalLengthGiven    int
}

func (s PoolStats) String() string {
	return fmt.Sprintf(`%d Makes | %d Borrows | %d Gives
%d borrowed items | %d given items
`,
		s.Makes, s.Borrows, s.Gives, s.TotalLengthBorrowed, s.TotalLengthGiven)
}

// The minimum size for an entry in the pool
var poolMinEntrySize = 10

// IDPool represents a thread-safe pool of reusable
// []ID to avoid allocations and GC pressure.
//
// NOTE: These are trivially pool-able due to their de-initialization
// being absolutely trivial. Maps are more difficult to clear quickly.
type IDPool struct {
	bufs  chan []ID
	stats PoolStats
	*sync.Mutex
}

// NewIDPool creates a properly initialized IDPool
func NewIDPool(maxPerSize int) IDPool {
	p := IDPool{
		bufs: make(chan []ID, maxPerSize),
	}
	p.Mutex = &sync.Mutex{}
	return p
}

// Borrow returns a pre-allocated buffer in the pool
// or a fresh buffer if none are available
func (p *IDPool) Borrow(minLen int) []ID {
	p.Lock()
	defer p.Unlock()

	var buf []ID
	select {
	case buf = <-p.bufs:
	default:
		if minLen < poolMinEntrySize {
			minLen = poolMinEntrySize
		}
		buf = make([]ID, minLen)
		p.stats.Makes++
	}

	p.stats.Borrows++
	p.stats.TotalLengthBorrowed += cap(buf)

	// Clear the buffer as its returned
	buf = buf[:0]

	return buf
}

// Give returns a buffer to the set
//
// If we have more buffers than we are configured to handle,
// further provided buffers are tossed out.
func (p *IDPool) Give(buf []ID) {
	p.Lock()
	defer p.Unlock()

	select {
	case p.bufs <- buf:
	default:
		// Let the buf reference fall out of scope
		// so the GC takes care of it
	}

	p.stats.Gives++
	p.stats.TotalLengthGiven += cap(buf)

}

// idPool is our pool of []ID for general purpose use.
var idPool = NewIDPool(100)

// GiveIDSlice returns a provided []ID to the memory pool.
//
// As []ID are used as return values, they escape the package bounds.
// This allows us to get them back when the caller is done with the,/
func GiveIDSlice(buf []ID) {
	idPool.Give(buf)
}

// IDPoolStats reports statistics for a Pool
func IDPoolStats() PoolStats {
	return idPool.stats
}
