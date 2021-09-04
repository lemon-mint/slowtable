package slowtable

import (
	"sync"
	"unsafe"

	"github.com/lemon-mint/slowtable/rwmu"
)

type value struct {
	size int64
	mu   rwmu.Mutex

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
