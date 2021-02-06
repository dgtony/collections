package bitset

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixedBitSet_BadIndex(t *testing.T) {
	var bs = NewFixed(123)
	assert.Panics(t, func() { bs.Get(123) })
	assert.Panics(t, func() { bs.Set(1 << 20) })
	assert.Panics(t, func() { bs.Clear(-2) })
}

func TestFixedBitSet_Basic(t *testing.T) {
	var (
		maxBits = 123
		ones    = []int{0, 21, 96, 113, 122}
		bs      = NewFixed(maxBits)
	)

	for _, bit := range ones {
		bs.Set(bit)
	}

	for i := 0; i < maxBits; i++ {
		if bs.Get(i) {
			assert.Contains(t, ones, i)
		} else {
			assert.NotContains(t, ones, i)
		}
	}

	for _, bit := range ones {
		bs.Clear(bit)
	}

	for i := 0; i < maxBits; i++ {
		assert.False(t, bs.Get(i))
	}
}

func TestFixedBitSet_Idempotence(t *testing.T) {
	var (
		maxBits = 123
		ones    = []int{0, 7, 8, 21, 64, 96, 113, 122}
		bs      = NewFixed(maxBits)
	)

	for _, bit := range append(ones, ones...) {
		bs.Set(bit)
	}

	for i := 0; i < maxBits; i++ {
		if bs.Get(i) {
			assert.Contains(t, ones, i)
		} else {
			assert.NotContains(t, ones, i)
		}
	}

	for _, bit := range append(ones, ones...) {
		bs.Clear(bit)
	}

	for i := 0; i < maxBits; i++ {
		assert.False(t, bs.Get(i))
	}
}

func TestFixedBitSet_Group(t *testing.T) {
	var (
		maxBits   = 1 << 10
		onesRatio = 0.5
		ones      = make(map[int]struct{}, maxBits/2)
		bs        = NewFixed(maxBits)
	)

	// set bits randomly
	for i := 0; i < maxBits; i++ {
		if rand.Float64() < onesRatio {
			ones[i] = struct{}{}
		}
	}

	for bit := range ones {
		bs.Set(bit)
	}

	var (
		numOnes  = bs.CountOnes()
		numZeros = bs.CountZeros()
	)

	assert.Equal(t, len(ones), numOnes)
	assert.Equal(t, maxBits-len(ones), numZeros)

	for _, bit := range bs.Ones() {
		assert.Contains(t, ones, bit)
	}

	for _, bit := range bs.Zeros() {
		assert.NotContains(t, ones, bit)
	}
}

func BenchmarkFixedBitSet_Basic(b *testing.B) {
	for _, tt := range []struct {
		name    string
		maxBits int
	}{
		{"small", 1 << 6},
		{"medium", 1 << 12},
		{"large", 1 << 18},
	} {
		b.Run(tt.name, func(b *testing.B) {
			var bs = NewFixed(tt.maxBits)
			for i := 0; i < b.N; i++ {
				// basic operations share the same time complexity
				bs.Set((tt.maxBits/2 + i) % tt.maxBits)
			}
		})
	}
}

func BenchmarkFixedBitSet_Group(b *testing.B) {
	for _, op := range []struct {
		name string
		op   func(bs *FixedBitSet)
	}{
		{"count", func(bs *FixedBitSet) { _ = bs.CountOnes() }},
		{"get-indexes", func(bs *FixedBitSet) { _ = bs.Zeros() }},
	} {
		for _, set := range []struct {
			name    string
			maxBits int
		}{
			{"small", 1 << 6},
			{"medium", 1 << 12},
			{"large", 1 << 18},
		} {
			b.Run(op.name+":"+set.name, func(b *testing.B) {
				var bs = NewFixed(set.maxBits)
				for i := 0; i < b.N; i++ {
					op.op(bs)
				}
			})
		}
	}
}

func BenchmarkFixedBitSet_BasicTS(b *testing.B) {
	var (
		maxBits = 1 << 12
		bs      = NewFixedThreadSafe(maxBits)
	)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := rand.Int() % maxBits
			if i%2 == 0 {
				bs.Set(i)
			} else {
				bs.Clear(i)
			}
		}
	})
}
