package main

import (
	"io/ioutil"
	"os"
)

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	count := V1EncodedString(input).Decode()
	println(count)
	count = V2EncodedString(input).Decode()
	println(count)
}
