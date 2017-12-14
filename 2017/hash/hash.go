package hash

import (
	"bytes"
	"fmt"
)

type Hash []byte

func xor(slice []byte) byte {
	value := slice[0]
	for _, s := range slice[1:] {
		value ^= s
	}
	return value
}

func (hash Hash) get(index int) byte {
	i := index % len(hash)
	return hash[i]
}

func (hash Hash) set(index int, value byte) {
	i := index % len(hash)
	hash[i] = value
}

func (hash Hash) slice(start, end int) []byte {
	values := make([]byte, 0)
	for i := start; i < end; i++ {
		values = append(values, hash.get(i))
	}
	return values
}

func (hash Hash) reverse(start, end int) {
	values := hash.slice(start, end)
	for i := start; i < end; i++ {
		value := values[end-i-1]
		hash.set(i, value)
	}
}

func (hash Hash) String() string {
	var buf bytes.Buffer
	for i := 0; i < 255; i += 16 {
		buf.WriteString(fmt.Sprintf("%02x", xor(hash[i:i+16])))
	}
	return buf.String()
}

func ComputeString(input string) Hash {
	lengths := []byte(input)
	for _, b := range []byte{17, 31, 73, 47, 23} {
		lengths = append(lengths, b)
	}
	return Compute(64, lengths)
}

func Compute(iterations int, lengths []byte) Hash {
	hash := make(Hash, 256)
	for i := 0; i < 256; i++ {
		hash[i] = byte(i)
	}

	skip := 0
	start := 0
	for i := 0; i < iterations; i++ {
		for _, length := range lengths {
			hash.reverse(start, start+int(length))
			start += int(length) + skip
			skip++
		}
	}

	return hash
}
