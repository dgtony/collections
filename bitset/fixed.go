package bitset

type FixedBitSet struct {
	data    []byte
	maxBits int
}

// Non-resizable bitset with basic operations.
func NewFixed(maxBits int) *FixedBitSet {
	var bytes = maxBits / 8
	if maxBits > bytes*8 {
		bytes++
	}
	return &FixedBitSet{data: make([]byte, bytes), maxBits: maxBits}
}

// Write one at the given bit position.
func (bs *FixedBitSet) Set(bit int) {
	bpos, shift := bs.position(bit)
	bs.data[bpos] |= 1 << shift
}

// Write zero at the given bit position.
func (bs *FixedBitSet) Clear(bit int) {
	bpos, shift := bs.position(bit)
	bs.data[bpos] &^= 1 << shift
}

// Determine if given bit position was set.
func (bs *FixedBitSet) Get(bit int) bool {
	bpos, shift := bs.position(bit)
	return bs.data[bpos]&(1<<shift) > 0
}

// Return positions of all previously set bits.
func (bs *FixedBitSet) Ones() []int {
	return bs.indexes(true)
}

// Return positions of all zero-bits in the bitset.
func (bs *FixedBitSet) Zeros() []int {
	return bs.indexes(false)
}

// Return total number of ones in the bitset.
func (bs *FixedBitSet) CountOnes() int {
	return bs.count(true)
}

// Return total number of zeros in the bitset.
func (bs *FixedBitSet) CountZeros() int {
	return bs.count(false)
}

// Clear entire bitset.
func (bs *FixedBitSet) Reset() {
	bs.data = make([]byte, len(bs.data))
}

// Return size of underlying bytearray.
func (bs *FixedBitSet) Size() int {
	return len(bs.data)
}

func (bs *FixedBitSet) position(bit int) (bytePos int, shift int) {
	if bit >= bs.maxBits || bit < 0 {
		panic("bit index out of range")
	}
	bytePos, shift = bit/8, bit%8
	return
}

func (bs *FixedBitSet) indexes(ones bool) []int {
	var idxs = make([]int, 0, 1)
	for bit := 0; bit < bs.maxBits; bit++ {
		var bpos, shift = bit / 8, bit % 8
		if ones == (bs.data[bpos]&(1<<shift) > 0) {
			idxs = append(idxs, bit)
		}
	}
	return idxs
}

func (bs *FixedBitSet) count(ones bool) int {
	var total int
	for bit := 0; bit < bs.maxBits; bit++ {
		var bpos, shift = bit / 8, bit % 8
		if ones == (bs.data[bpos]&(1<<shift) > 0) {
			total++
		}
	}
	return total
}
