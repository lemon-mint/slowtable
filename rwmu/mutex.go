package rwmu

import (
	"runtime"
	"sync/atomic"
)

type Mutex struct {
	state int32
}

func (m *Mutex) Lock() {
	for !atomic.CompareAndSwapInt32(&m.state, 0, 1) {
		runtime.Gosched()
	}
}

func (m *Mutex) Unlock() {
	atomic.StoreInt32(&m.state, 0)
}

// Rlock == Lock
func (m *Mutex) RLock() {
	for !atomic.CompareAndSwapInt32(&m.state, 0, 1) {
		runtime.Gosched()
	}
}

// RUnlock == Unlock
func (m *Mutex) RUnlock() {
	atomic.StoreInt32(&m.state, 0)
}
