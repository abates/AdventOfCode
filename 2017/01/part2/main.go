package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/abates/AdventOfCode/util"
)

func main() {
	list := util.CircularIntList{}
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	str := strings.TrimSpace(string(b))
	for _, s := range strings.Split(str, "") {
		i, _ := strconv.Atoi(s)
		list.Add(i)
	}

	sum := 0
	for i := 0; i < len(str); i++ {
		value := list.Next()
		if value == list.PeekN(len(str)/2) {
			sum += value
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
