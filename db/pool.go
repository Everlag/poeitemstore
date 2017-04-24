package db

// The minimum size for an entry in the pool
var poolMinEntrySize = 100

// IDMapPool represents a thread-safe pool of reusable
// sets for IndexQuery
type IDMapPool struct {
	bufs chan map[ID]struct{}
}

// NewIDMapPool creates a properly initialized IDMapPool
func NewIDMapPool(maxPerSize int) IDMapPool {
	p := IDMapPool{
		bufs: make(chan map[ID]struct{}, maxPerSize),
	}
	return p
}

// Borrow returns a pre-allocated buffer in the pool
// or a fresh buffer if none are available
func (p *IDMapPool) Borrow(minLen int) map[ID]struct{} {
	var buf map[ID]struct{}
	select {
	case buf = <-p.bufs:
	default:
		if minLen < poolMinEntrySize {
			minLen = poolMinEntrySize
		}
		buf = make(map[ID]struct{}, minLen)
	}

	// Clear out the map
	for k := range buf {
		delete(buf, k)
	}

	return buf
}

// Give returns a buffer to the set
//
// If we have more buffers than we are configured to handle,
// further provided buffers are tossed out.
func (p *IDMapPool) Give(buf map[ID]struct{}) {
	select {
	case p.bufs <- buf:
	default:
		// Let the buf reference fall out of scope
		// so the GC takes care of it
	}
}
