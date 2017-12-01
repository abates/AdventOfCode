package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

var key = "yzbqklnj"

func compute(i int) {
	sum := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", key, i))))
	if "000000" == sum[0:6] {
		fmt.Printf("%d\n", i)
		os.Exit(0)
	}
}

func main() {
	for i := 0; ; i++ {
		compute(i)
	}
}
