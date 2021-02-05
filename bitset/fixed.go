package bitset

// Non-resizable bitset.
type FixedBitSet struct {
	data    []byte
	maxBits int
}

func NewFixed(maxBits int) *FixedBitSet {
	var bytes = maxBits / 8
	if maxBits > bytes*8 {
		bytes++
	}
	return &FixedBitSet{data: make([]byte, bytes), maxBits: maxBits}
}

func (bs *FixedBitSet) Set(bit int) {
	bpos, shift := bs.position(bit)
	bs.data[bpos] |= 1 << shift
}

func (bs *FixedBitSet) Clear(bit int) {
	bpos, shift := bs.position(bit)
	bs.data[bpos] &^= 1 << shift
}

func (bs *FixedBitSet) Get(bit int) bool {
	bpos, shift := bs.position(bit)
	return bs.data[bpos]&(1<<shift) > 0
}

func (bs *FixedBitSet) Ones() []int {
	return bs.indexes(true)
}

func (bs *FixedBitSet) Zeros() []int {
	return bs.indexes(false)
}

func (bs *FixedBitSet) CountOnes() int {
	return bs.count(true)
}

func (bs *FixedBitSet) CountZeros() int {
	return bs.count(false)
}

func (bs *FixedBitSet) Reset() {
	bs.data = make([]byte, len(bs.data))
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
