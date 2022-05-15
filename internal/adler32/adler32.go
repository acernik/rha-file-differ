package adler32

// https://www.youtube.com/watch?v=BWqH4O7OuyY

// mod is the largest prime number that fits in the 2^16.
// The reason why the division is made with primes is that when you divide by regular numbers you get consistent patterns.
// mod is used to modulo both A and B sums.
// The reason why the mod is done is that the numbers are quite large and when mod is applied they fit within 2^16.
const (
	mod = 65521
)

type Adler32 struct {
	window []byte // Slice of bytes that is currently being processed.
	count  int    // Current number of bytes processed.
	old    uint8  // Latest byte processed.
	// A and B are accumulative sums of the characters up until this point in the stream.
	a uint16 // A is a sum of accumulative sums of the characters up until this point in the stream.
	b uint16
}

// New returns a value of type Adler32.
func New() Adler32 {
	return Adler32{
		window: []byte{},
		count:  0,
		a:      0,
		b:      0,
	}
}

// Write calculates initial checksum from byte slice.
func (adl Adler32) Write(data []byte) Adler32 {
	for index, char := range data {
		adl.a += uint16(char)
		adl.b += uint16(len(data)-index) * uint16(char)
		adl.count++
	}

	adl.a %= mod
	adl.b %= mod

	return adl
}

// Sum returns hash for current a and b values.
// Two 16-bit checksums A and B are calculated and their bits are concatenated into a 32-bit integer.
func (adl Adler32) Sum() uint32 {
	return uint32(adl.b)<<16 | uint32(adl.a)&0xFFFFF
}

// Window returns current slice of bytes (chunk) that is being processed.
func (adl Adler32) Window() []byte {
	return adl.window
}

// Count returns number of bytes that have been processed.
func (adl Adler32) Count() int {
	return adl.count
}

// Removed returns the latest byte processed.
func (adl Adler32) Removed() uint8 {
	return adl.old
}

// RollIn adds byte to rolling checksum.
func (adl Adler32) RollIn(input byte) Adler32 {
	adl.a = (adl.a + uint16(input)) % mod
	adl.b = (adl.b + adl.a) % mod
	adl.window = append(adl.window, input)
	adl.count++

	return adl
}

// RollOut removes leftmost byte from the checksum.
func (adl Adler32) RollOut() Adler32 {
	if len(adl.window) == 0 {
		adl.count = 0

		return adl
	}

	adl.old = adl.window[0]
	adl.a = (adl.a - uint16(adl.old)) % mod
	adl.b = (adl.b - (uint16(len(adl.window)) * uint16(adl.old))) % mod
	adl.window = adl.window[1:]
	adl.count--

	return adl
}
