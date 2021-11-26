package benchmark_test

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash/fnv"
	"hash/maphash"
	"testing"

	"github.com/bytedance/gopkg/util/xxhash3"
	"github.com/cespare/xxhash/v2"
	"github.com/dchest/siphash"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/zeebo/blake3"
	zeebo_xxh "github.com/zeebo/xxh3"
	"golang.org/x/crypto/blake2b"
)

var TestData = [][]byte{
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
	[]byte("Pack my box with five dozen liquor jugs."),
	[]byte("The quick brown fox jumps over the lazy dog."),
	[]byte("Waltz, nymph, for quick jigs vex Bud."),
	[]byte("Sphinx of black quartz, judge my vow."),
}

func BenchmarkSHA256(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = sha256.Sum256(data)
			}
		}
	})
}

func BenchmarkSHA512(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = sha512.Sum512(data)
			}
		}
	})
}

func BenchmarkXXH64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = xxhash.Sum64(data)
			}
		}
	})
}
func BenchmarkBLAKE3_256(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = blake3.Sum256(data)
			}
		}
	})
}

func BenchmarkBLAKE3_512(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = blake3.Sum512(data)
			}
		}
	})
}

func BenchmarkBLAKE2B_256(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = blake2b.Sum256(data)
			}
		}
	})
}

func BenchmarkBLAKE2B_512(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = blake2b.Sum512(data)
			}
		}
	})
}

func BenchmarkXXH3(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = xxhash3.Hash(data)
			}
		}
	})
}

func BenchmarkZeeboXXH3(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				_ = zeebo_xxh.Hash(data)
			}
		}
	})
}

func BenchmarkMaphash(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var h maphash.Hash
		for pb.Next() {
			for _, data := range TestData {
				h.Write(data)
				h.Sum(nil)
				h.Reset()
			}
		}
	})
}

func BenchmarkSiphash(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var h = siphash.New([]byte("0123456789101112"))
		for pb.Next() {
			for _, data := range TestData {
				h.Write(data)
				h.Sum64()
				h.Reset()
			}
		}
	})
}

func BenchmarkStdFnv32(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var h = fnv.New32()
		for pb.Next() {
			for _, data := range TestData {
				h.Write(data)
				h.Sum32()
				h.Reset()
			}
		}
	})
}

func BenchmarkStdFnv32a(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var h = fnv.New32a()
		for pb.Next() {
			for _, data := range TestData {
				h.Write(data)
				h.Sum32()
				h.Reset()
			}
		}
	})
}

func BenchmarkStdFnv64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var h = fnv.New64()
		for pb.Next() {
			for _, data := range TestData {
				h.Write(data)
				h.Sum64()
				h.Reset()
			}
		}
	})
}

func BenchmarkStdFnv64a(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var h = fnv.New64a()
		for pb.Next() {
			for _, data := range TestData {
				h.Write(data)
				h.Sum64()
				h.Reset()
			}
		}
	})
}

func BenchmarkFnv1a(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, data := range TestData {
				fnv1a.HashBytes64(data)
			}
		}
	})
}
