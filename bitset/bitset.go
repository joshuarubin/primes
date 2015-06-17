package bitset

import "math"

const (
	// Byte is 8 bits
	Byte = 8
)

// Bitset is an object that makes working with large numbers of boolean values
// simple and memory efficient
type Bitset []byte

// Len returns the size of the Bitset in bytes
func (s Bitset) Len() uint64 {
	return uint64(len(s))
}

// SetAll enables all bits in the Bitset
func (s Bitset) SetAll() Bitset {
	for i := 0; i < len(s); i++ {
		s[i] = math.MaxUint8
	}

	return s
}

// Unset the bit at index i
func (s Bitset) Unset(i uint64) Bitset {
	b, mask := s.byteFor(i)
	*b = *b & flipBits(mask)
	return s
}

// Set the bit at index i
func (s Bitset) Set(i uint64) Bitset {
	b, mask := s.byteFor(i)
	*b = *b | mask
	return s
}

// IsSet returns whether the bit at index i is set
func (s Bitset) IsSet(i uint64) bool {
	b, mask := s.byteFor(i)
	return *b&mask > 0
}

// Flip toggles the bit at index i and returns the new value
func (s Bitset) Flip(i uint64) bool {
	b, mask := s.byteFor(i)

	if *b&mask == 0 {
		// bit is not set
		*b = *b | mask
		return true
	}

	// bit is set
	*b = *b & flipBits(mask)
	return false
}

func (s Bitset) byteFor(i uint64) (b *byte, mask byte) {
	b = &s[i/Byte]
	mask = byte(1) << (i % Byte)
	return
}

// Size returns the number of bytes required for the Bitset to contain at
// least n values
func Size(n uint64) uint64 {
	return uint64(math.Floor(float64(n)/Byte)) + 1
}

// New returns a new Bitset big enough to hold at least n values
func New(n uint64) Bitset {
	return make(Bitset, int(Size(n)))
}

// Max returns the highest value that can be set
func (s Bitset) Max() uint64 {
	return s.Len()*Byte - 1
}

// ListSet returns a set of all enabled indexes
func (s Bitset) ListSet() []uint64 {
	ret := make([]uint64, 0, s.Max()+1)

	for i := uint64(0); i <= s.Max(); i++ {
		if s.IsSet(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

// ListUnset returns a set of all unset indexes
func (s Bitset) ListUnset() []uint64 {
	ret := make([]uint64, 0, s.Max()+1)

	for i := uint64(0); i <= s.Max(); i++ {
		if !s.IsSet(i) {
			ret = append(ret, i)
		}
	}

	return ret
}

func flipBits(b byte) byte {
	return math.MaxUint8 &^ b
}
