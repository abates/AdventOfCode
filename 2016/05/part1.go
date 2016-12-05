package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	key := "ojvtpuvg"
	p := make([]byte, 0)
	for i := 0; ; i++ {
		sum := md5.Sum([]byte(fmt.Sprintf("%s%d", key, i)))
		if sum[0]&0xff == 0 && sum[1]&0xff == 0 && sum[2]&0xf0 == 0 {
			p = append(p, 0x0f&sum[2])
			if len(p) == 8 {
				break
			}
		}
	}

	for _, b := range p {
		fmt.Printf("%x", b)
	}
	fmt.Printf("\n")
}
