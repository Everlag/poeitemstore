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
	TotalLengthMaked    int
	TotalLengthBorrowed int
	TotalLengthGiven    int
}

func (s PoolStats) String() string {
	return fmt.Sprintf(`%d Makes | %d Borrows | %d Gives
%d make items | %d borrowed items | %d given items
`,
		s.Makes, s.Borrows, s.Gives,
		s.TotalLengthMaked, s.TotalLengthBorrowed, s.TotalLengthGiven)
}

// The minimum size for an entry in the pool
var poolMinEntrySize = 10

// PoolMakerCB is a function that generates a slice
// of minLen of the desired type when the pool
// doesn't have any buffers cached.
type PoolMakerCB func(minLen int) interface{}

// PoolClearCB clears the state of the provided buffer.
// This returns the cleared buffer.
//
// This is necessary as 'buf.([]interface{})[:0]' will always
// fail as a result of []interface having a memory layout different from
// the underlying type of buf.
type PoolClearCB func(buf interface{}) interface{}

// PoolCapCB returns the capacity of a given buffer.
//
// I wish we had generics D:
type PoolCapCB func(buf interface{}) int

// Pool represents a thread-safe pool of reusable
// slices to reduce allocations.
//
// NOTE: Only slices are handled by this pool.
// Maps have a high clearing cost, so we don't pool them.
type Pool struct {
	bufs     chan interface{}
	maker    PoolMakerCB
	clearer  PoolClearCB
	capacity PoolCapCB

	stats PoolStats
	*sync.Mutex
}

// NewPool creates a properly initialized IDPool
func NewPool(maxSize int,
	maker PoolMakerCB, clearer PoolClearCB, capacity PoolCapCB) Pool {
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

var idPool = NewPool(100, func(minLen int) interface{} {
	return make([]ID, 0, minLen)
},
	func(buf interface{}) interface{} {
		return buf.([]ID)[:0]
	},
	func(buf interface{}) int {
		return cap(buf.([]ID))
	})

var idMapPool = NewPool(100, func(minLen int) interface{} {
	return make(map[ID]int, minLen)
},
	func(buf interface{}) interface{} {
		safe := buf.(map[ID]int)
		for k := range safe {
			delete(safe, k)
		}
		return safe
	},
	func(buf interface{}) int {
		return len(buf.(map[ID]int))
	})

var itemPool = NewPool(100, func(minLen int) interface{} {
	return make([]Item, 0, minLen)
},
	func(buf interface{}) interface{} {
		return buf.([]Item)[:0]
	},
	func(buf interface{}) int {
		return cap(buf.([]Item))
	})

// GiveIDSlice returns a provided []ID to the memory pool.
//
// As []ID are used as return values, they escape the package bounds.
// This allows us to get them back when the caller is done with the.
func GiveIDSlice(buf []ID) {
	idPool.Give(buf)
}

// GiveItemSlice returns a provided []Item to the memory pool.
//
// As []Item are used as return values, they escape the package bounds.
// This allows us to get them back when the caller is done with them
func GiveItemSlice(buf []Item) {
	itemPool.Give(buf)
}

// GiveItemSliceSlice returns a provided [][]Item to the []Item memory pool.
//
// As []Item are used as return values, they escape the package bounds.
// This allows us to get them back when the caller is done with them
func GiveItemSliceSlice(buf [][]Item) {
	for i := range buf {
		itemPool.Give(buf[i])
	}
}

// IDPoolStats reports statistics for a Pool
func IDPoolStats() PoolStats {
	return idPool.stats
}

// ItemPoolStats reports statistics for a Pool
func ItemPoolStats() PoolStats {
	return itemPool.stats
}
