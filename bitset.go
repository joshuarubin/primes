package primes

import "math"

// BitSet is an object that makes working with large numbers of boolean values
// simple and memory efficient
type BitSet []byte

// Len returns the length of the BitSet in bytes
func (s BitSet) Len() uint64 {
	return uint64(len(s))
}

// SetAll enables all bits in the BitSet
func (s BitSet) SetAll() BitSet {
	for i := 0; i < len(s); i++ {
		s[i] = math.MaxUint8
	}

	return s
}

// Unset the bit at index i
func (s BitSet) Unset(i uint64) BitSet {
	b, mask := s.byteFor(i)
	*b = *b & flipBits(mask)
	return s
}

// Set the bit at index i
func (s BitSet) Set(i uint64) BitSet {
	b, mask := s.byteFor(i)
	*b = *b | mask
	return s
}

// IsSet returns wheter the bit at index i is set
func (s BitSet) IsSet(i uint64) bool {
	b, mask := s.byteFor(i)
	return *b&mask > 0
}

func (s BitSet) byteFor(i uint64) (b *byte, mask byte) {
	return &s[i/8], byte(1) << ((i % 8) - 1)
}

// NewBitSet returns a new BitSet big enough to hold at least l values
func NewBitSet(l uint64) BitSet {
	numBlocks := uint64(math.Floor(float64(l)/8)) + 1
	return make(BitSet, int(numBlocks))
}
