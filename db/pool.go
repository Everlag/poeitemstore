package db

import "github.com/Everlag/poeitemstore/pool"

var idPool = pool.NewPool(100, func(minLen int) interface{} {
	return make([]ID, 0, minLen)
},
	func(buf interface{}) interface{} {
		return buf.([]ID)[:0]
	},
	func(buf interface{}) int {
		return cap(buf.([]ID))
	})

var idMapPool = pool.NewPool(100, func(minLen int) interface{} {
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

var itemPool = pool.NewPool(100, func(minLen int) interface{} {
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
func IDPoolStats() pool.Stats {
	return idPool.Stats()
}

// ItemPoolStats reports statistics for a Pool
func ItemPoolStats() pool.Stats {
	return itemPool.Stats()
}
