package benchmark_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/antlinker/go-cmap"
	"github.com/cornelk/hashmap"
	"github.com/lemon-mint/slowtable"
	"github.com/snowmerak/concurrent"
)

const Size = 8192

var keys = func() []string {
	tmp := make([]string, Size)
	for i := 0; i < Size; i++ {
		tmp[i] = strconv.Itoa(i)
	}
	return tmp
}()

var keysBytes = func() [][]byte {
	tmp := make([][]byte, Size)
	for i := 0; i < Size; i++ {
		tmp[i] = []byte(strconv.Itoa(i))
	}
	return tmp
}()

func BenchmarkSlowTablGetS(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	for i := 0; i < Size; i++ {
		table.SetS(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				table.GetS(keys[i])
			}
		}
	})
}

func BenchmarkSlowTableSetS(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				table.SetS(keys[i], nil)
			}
		}
	})
}

func BenchmarkSlowTableGetSetS(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				table.SetS(keys[i], nil)
				table.GetS(keys[i])
			}
		}
	})
}

func BenchmarkSlowTablGet(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	for i := 0; i < Size; i++ {
		table.Set(keysBytes[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				table.Get(keysBytes[i])
			}
		}
	})
}

func BenchmarkSlowTableSet(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				table.Set(keysBytes[i], nil)
			}
		}
	})
}

func BenchmarkSlowTableGetSet(b *testing.B) {
	table := slowtable.NewTable(nil, 65535)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				table.Set(keysBytes[i], nil)
				table.Get(keysBytes[i])
			}
		}
	})
}

func BenchmarkSyncMapGetS(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < Size; i++ {
		m.Store(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Load(keys[i])
			}
		}
	})
}

func BenchmarkSyncMapSetS(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Store(keys[i], nil)
			}
		}
	})
}

func BenchmarkSyncMapGetSetS(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Store(keys[i], nil)
				m.Load(keys[i])
			}
		}
	})
}

func BenchmarkConcurrentGet(b *testing.B) {
	m := concurrent.NewMap()
	for i := 0; i < Size; i++ {
		m.Set(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Get(keys[i])
			}
		}
	})
}

func BenchmarkConcurrentSet(b *testing.B) {
	m := concurrent.NewMap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Set(keys[i], nil)
			}
		}
	})
}

func BenchmarkConcurrentGetSet(b *testing.B) {
	m := concurrent.NewMap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Set(keys[i], nil)
				m.Get(keys[i])
			}
		}
	})
}

func BenchmarkHashMapGet(b *testing.B) {
	m := hashmap.New(hashmap.DefaultSize)
	for i := 0; i < Size; i++ {
		m.Set(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Get(keys[i])
			}
		}
	})
}

func BenchmarkHashMapSet(b *testing.B) {
	m := hashmap.New(hashmap.DefaultSize)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Set(keys[i], nil)
			}
		}
	})
}
func BenchmarkHashMapGetSet(b *testing.B) {
	m := hashmap.New(hashmap.DefaultSize)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Set(keys[i], nil)
				m.Get(keys[i])
			}
		}
	})
}

func BenchmarkCMapGet(b *testing.B) {
	m := cmap.NewConcurrencyMap()
	for i := 0; i < Size; i++ {
		m.Set(keys[i], nil)
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Get(keys[i])
			}
		}
	})
}

func BenchmarkCMapSet(b *testing.B) {
	m := cmap.NewConcurrencyMap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Set(keys[i], nil)
			}
		}
	})
}

func BenchmarkCMapGetSet(b *testing.B) {
	m := cmap.NewConcurrencyMap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for i := 0; i < Size; i++ {
				m.Set(keys[i], nil)
				m.Get(keys[i])
			}
		}
	})
}
