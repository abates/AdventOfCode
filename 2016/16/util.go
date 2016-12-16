package main

import (
	"bytes"
	"fmt"
	"math/big"
)

func Generate(bitLen int, a *big.Int) *big.Int {
	b := &big.Int{}
	b.Set(a)
	b.Lsh(b, uint(bitLen+1))
	bytes := a.Bytes()
	for i := 0; i < bitLen; i++ {
		//println("Getting byte", len(bytes)-(i/8), "out of", len(bytes), "for bit", i)
		byt := byte(0x00)
		if (i / 8) < len(bytes) {
			byt = bytes[len(bytes)-(i/8)-1]
		}
		bit := uint8(i % 8)
		v := uint(0)
		if byt&(1<<bit) == 0 {
			v = 1
		}
		b.SetBit(b, bitLen-i-1, v)
		//b.SetBit(b, i, v)
	}

	return b
}

func Fill(length int, bitLen int, a *big.Int) *big.Int {
	add := 0
	if bitLen > a.BitLen() {
		add = bitLen - a.BitLen()
	}
	for ; a.BitLen()+add < length; a = Generate(a.BitLen()+add, a) {
		fmt.Printf("Length is now %d\n", a.BitLen())
	}
	return a.Rsh(a, uint(a.BitLen()+add-length))
}

func FillString(length int, input string) (*big.Int, error) {
	i := &big.Int{}
	_, err := fmt.Sscanf(input, "%b", i)
	bitLen := i.BitLen()
	if input[0] == '0' {
		bitLen++
	}
	return Fill(length, bitLen, i), err
}

func checksum(a string) string {
	checksum := &bytes.Buffer{}
	one := []byte("1")
	zero := []byte("0")
	for {
		for i := 0; i < len(a); i += 2 {
			if i+1 < len(a) {
				if a[i] == a[i+1] {
					checksum.Write(one)
				} else {
					checksum.Write(zero)
				}
			} else if a[i] == '0' {
				checksum.Write(one)
			} else {
				checksum.Write(zero)
			}
		}

		a = checksum.String()
		if len(a)%2 == 1 {
			break
		}
		checksum.Reset()
	}

	return a
}

func Checksum(leadingZero bool, a *big.Int) string {
	s := ""
	if leadingZero {
		s = fmt.Sprintf("0%b", a)
	} else {
		s = fmt.Sprintf("%b", a)
	}
	return checksum(s)
}
