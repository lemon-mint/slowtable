package benchmark_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/lemon-mint/slowtable"
)

var keys = func() []string {
	tmp := make([]string, 8192)
	for i := 0; i < 8192; i++ {
		tmp[i] = strconv.Itoa(i)
	}
	return tmp
}()

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
