package slowtable

import (
	"bytes"
	"sync"
	"unsafe"

	"github.com/cespare/xxhash/v2"
	"github.com/lemon-mint/slowtable/nocopy"
)

// NewTable : Create a new table
// hash : xxHash64 (default) if hash == nil
// keysize : The number of keys in the table (recommended: 65536)
func NewTable(hash func([]byte) uint64, keysize uint64) *Table {
	pool := sync.Pool{
		New: func() interface{} {
			return &item{}
		},
	}

	if hash == nil {
		hash = xxhash.Sum64
	}
	return &Table{
		entries:  make([]value, keysize),
		hash:     hash,
		keysize:  keysize,
		itempool: pool,
	}
}

func (t *Table) PoolPreload() {
	var tmp [8192]*item
	for i := 0; i < 8192; i++ {
		tmp[i] = t.itempool.Get().(*item)
	}
	for i := 0; i < 8192; i++ {
		t.itempool.Put(tmp[i])
	}
}

func (t *Table) putitem(i *item) {
	i.next = nil
	i.val = nil
	i.key = nil
	t.itempool.Put(i)
}

// Set : Set the value for the given key
// Thread safe
func (t *Table) Set(key []byte, val unsafe.Pointer) {
	hash := t.hash(key) % t.keysize
	t.entries[hash].mu.Lock()
	defer t.entries[hash].mu.Unlock()

	if t.entries[hash].size == 0 {
		t.entries[hash].next = t.itempool.Get().(*item)
		t.entries[hash].next.key = key
		t.entries[hash].next.val = val
		t.entries[hash].size = 1
		return
	}

	var i *item = t.entries[hash].next
	for {
		if bytes.Equal(key, i.key) {
			i.val = val
			return
		} else {
			if i.next == nil {
				i.next = t.itempool.Get().(*item)
				i.next.key = key
				i.next.val = val
				t.entries[hash].size++
				return
			} else {
				i = i.next
			}
		}
	}
}

// Get : Get the value for the given key
// Thread safe
func (t *Table) Get(key []byte) (unsafe.Pointer, bool) {
	hash := t.hash(key) % t.keysize
	t.entries[hash].mu.RLock()
	defer t.entries[hash].mu.RUnlock()

	if t.entries[hash].size == 0 {
		return nil, false
	}

	i := t.entries[hash].next
	for {
		if bytes.Equal(key, i.key) {
			return i.val, true
		} else {
			if i.next == nil {
				return nil, false
			} else {
				i = i.next
			}
		}
	}
}

// Delete : Delete the item from the table
// Thread safe
func (t *Table) Delete(key []byte) {
	hash := t.hash(key) % t.keysize
	t.entries[hash].mu.Lock()
	defer t.entries[hash].mu.Unlock()

	if t.entries[hash].size == 0 {
		return
	}

	i := t.entries[hash].next
	for {
		if bytes.Equal(key, i.key) {
			t.entries[hash].size--
			t.entries[hash].next = i.next
			t.putitem(i)
			return
		} else {
			if i.next == nil {
				return
			} else {
				i = i.next
			}
		}
	}
}

// Clear : Clear the table
// Note: This function drops all the items in the table (GC overhead) No Pooling
func (t *Table) Clear() {
	for i := uint64(0); i < t.keysize; i++ {
		t.entries[i].mu.Lock()
		t.entries[i].size = 0
		t.entries[i].next = nil
		t.entries[i].mu.Unlock()
	}
}

func (t *Table) Exists(key []byte) bool {
	_, ok := t.Get(key)
	return ok
}

// GetS : Get the value for the given key(String)
// Thread safe
func (t *Table) GetS(key string) (unsafe.Pointer, bool) {
	return t.Get(nocopy.S2B(key))
}

// SetS : Set the value for the given key(String)
// Thread safe
func (t *Table) SetS(key string, val unsafe.Pointer) {
	t.Set(nocopy.S2B(key), val)
}

// DeleteS : Delete the item from the table
// Thread safe
func (t *Table) DeleteS(key string) {
	t.Delete(nocopy.S2B(key))
}
