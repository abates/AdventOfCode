package main

import "fmt"

type Hash []byte

func (h Hash) Get(index int) byte {
	i := index % len(h)
	return h[i]
}

func (h Hash) Set(index int, value byte) {
	i := index % len(h)
	h[i] = value
}

func (h Hash) Slice(start, end int) []byte {
	values := make([]byte, 0)
	for i := start; i < end; i++ {
		values = append(values, h.Get(i))
	}
	return values
}

func (h Hash) Reverse(start, end int) {
	values := h.Slice(start, end)
	for i := start; i < end; i++ {
		value := values[end-i-1]
		h.Set(i, value)
	}
}

func xor(slice []byte) byte {
	value := slice[0]
	for _, s := range slice[1:] {
		value ^= s
	}
	return value
}

func hash(input string) string {
	lengths := []byte(input)
	for _, b := range []byte{17, 31, 73, 47, 23} {
		lengths = append(lengths, b)
	}

	hash := make(Hash, 256)
	for i := 0; i < 256; i++ {
		hash[i] = byte(i)
	}

	skip := 0
	start := 0
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			hash.Reverse(start, start+int(length))
			start += int(length) + skip
			skip++
		}
	}

	str := ""
	for i := 0; i < 255; i += 16 {
		str += fmt.Sprintf("%02x", xor(hash[i:i+16]))
	}

	return str
}

func main() {
	fmt.Printf("Hash: %q\n", hash("14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244"))
}
