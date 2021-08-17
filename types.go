package slowtable

import (
	"sync"
	"unsafe"
)

type value struct {
	size int64
	mu   sync.RWMutex

	next *item
}

type item struct {
	key  []byte
	val  unsafe.Pointer
	next *item
}

type Table struct {
	entries []value
	hash    func([]byte) uint64
	keysize uint64

	itempool sync.Pool
}
