package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	key := "yzbqklnj"
	for i := 0; ; i++ {
		sum := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", key, i))))
		if "00000" == sum[0:5] {
			fmt.Printf("%d\n", i)
			break
		}
	}

}
