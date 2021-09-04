package benchmark_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/cornelk/hashmap"
	"github.com/lemon-mint/slowtable"
	"github.com/snowmerak/concurrent"
)

var keys = func() []string {
	tmp := make([]string, 8192)
	for i := 0; i < 8192; i++ {
		tmp[i] = strconv.Itoa(i)
	}
	return tmp
}()

var keysBytes = func() [][]byte {
	tmp := make([][]byte, 8192)
	for i := 0; i < 8192; i++ {
		tmp[i] = []byte(strconv.Itoa(i))
	}
	return tmp
}()

func BenchmarkSlowTablGetS(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	table.PoolPreload()
	for i := 0; i < 8192; i++ {
		table.SetS(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				table.GetS(keys[i])
			}
		}
	})
}

func BenchmarkSlowTableSetS(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	table.PoolPreload()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				table.SetS(keys[i], nil)
			}
		}
	})
}

func BenchmarkSlowTablGet(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	table.PoolPreload()
	for i := 0; i < 8192; i++ {
		table.Set(keysBytes[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				table.Get(keysBytes[i])
			}
		}
	})
}

func BenchmarkSlowTableSet(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	table.PoolPreload()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				table.Set(keysBytes[i], nil)
			}
		}
	})
}

func BenchmarkSyncMapGetS(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < 8192; i++ {
		m.Store(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				m.Load(keys[i])
			}
		}
	})
}

func BenchmarkSyncMapSetS(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				m.Store(keys[i], nil)
			}
		}
	})
}

func BenchmarkConcurrentGet(b *testing.B) {
	m := concurrent.NewMap()
	for i := 0; i < 8192; i++ {
		m.Set(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				m.Get(keys[i])
			}
		}
	})
}

func BenchmarkConcurrentSet(b *testing.B) {
	m := concurrent.NewMap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				m.Set(keys[i], nil)
			}
		}
	})
}

func BenchmarkHashMapGet(b *testing.B) {
	m := hashmap.New(hashmap.DefaultSize)
	for i := 0; i < 8192; i++ {
		m.Set(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				m.Get(keys[i])
			}
		}
	})
}

func BenchmarkHashMapSet(b *testing.B) {
	m := hashmap.New(hashmap.DefaultSize)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < 8192; i++ {
				m.Set(keys[i], nil)
			}
		}
	})
}
