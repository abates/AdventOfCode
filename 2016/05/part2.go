package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	key := "ojvtpuvg"
	p := make([]byte, 8)
	keys := map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true}
	for i := 0; ; i++ {
		sum := md5.Sum([]byte(fmt.Sprintf("%s%d", key, i)))
		if sum[0]&0xff == 0 && sum[1]&0xff == 0 && sum[2]&0xf0 == 0 {
			position := int(sum[2] & 0x0f)
			if position >= len(p) {
				continue
			}
			if _, found := keys[position]; found {
				p[position] = sum[3] >> 4
				delete(keys, position)
				if len(keys) == 0 {
					break
				}
			}
		}
	}

	for _, b := range p {
		fmt.Printf("%x", b)
	}
	fmt.Printf("\n")
}
