package bitset

import "sync"

type FixedThreadSafeBitSet struct {
	bs *FixedBitSet
	mx sync.RWMutex
}

// Thread-safe wrapper around fixed-size bitset.
func NewFixedThreadSafe(maxBits int) *FixedThreadSafeBitSet {
	return &FixedThreadSafeBitSet{bs: NewFixed(maxBits)}
}

func (bs *FixedThreadSafeBitSet) Set(bit int) {
	bs.mx.Lock()
	defer bs.mx.Unlock()
	bs.bs.Set(bit)
}

func (bs *FixedThreadSafeBitSet) Clear(bit int) {
	bs.mx.Lock()
	defer bs.mx.Unlock()
	bs.bs.Clear(bit)
}

func (bs *FixedThreadSafeBitSet) Get(bit int) bool {
	bs.mx.RLock()
	defer bs.mx.RUnlock()
	return bs.bs.Get(bit)
}

func (bs *FixedThreadSafeBitSet) Ones() []int {
	bs.mx.RLock()
	defer bs.mx.RUnlock()
	return bs.bs.Ones()
}

func (bs *FixedThreadSafeBitSet) Zeros() []int {
	bs.mx.RLock()
	defer bs.mx.RUnlock()
	return bs.bs.Zeros()
}

func (bs *FixedThreadSafeBitSet) CountOnes() int {
	bs.mx.RLock()
	defer bs.mx.RUnlock()
	return bs.bs.CountOnes()
}

func (bs *FixedThreadSafeBitSet) CountZeros() int {
	bs.mx.RLock()
	defer bs.mx.RUnlock()
	return bs.bs.CountZeros()
}

func (bs *FixedThreadSafeBitSet) Reset() {
	bs.mx.Lock()
	defer bs.mx.Unlock()
	bs.bs.Reset()
}
