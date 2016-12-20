package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strconv"
	"strings"
)

func main() {
	ranges := &Ranges{}
	for _, line := range util.ReadInput() {
		values := strings.Split(line, "-")
		low, _ := strconv.Atoi(values[0])
		high, _ := strconv.Atoi(values[1])

		if high < low {
			l := high
			high = low
			low = l
		}
		ranges.Insert(&Range{low, high})
	}

	fmt.Printf("Lowest: %d\n", ranges.LowestException())
	count := 0
	for i := 0; i < len(ranges.ranges)-1; i++ {
		rng1 := ranges.ranges[i]
		rng2 := ranges.ranges[i+1]

		rng3 := &Range{rng1.high + 1, rng2.low - 1}
		if rng3.Count() < 0 {
			panic(fmt.Sprintf("Panic at %d: %d\n%v - %v", i, rng3.Count(), rng1, rng2))
		}
		count += rng3.Count()
	}
	fmt.Printf("Count: %d\n", count)
}
