package main

import (
	"fmt"
)

func main() {
	keys := GetKeys("yjdafjpo", Hash)
	fmt.Printf("Hash: %s at %d\n", keys[63].Hash, keys[63].Index)

	keys = GetKeys("yjdafjpo", HashRepeat(2016))
	fmt.Printf("Hash: %s at %d\n", keys[63].Hash, keys[63].Index)

}
