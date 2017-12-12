package main

import "fmt"

type List []int

func (l List) Get(index int) int {
	i := index % len(l)
	return l[i]
}

func (l List) Set(index, value int) {
	i := index % len(l)
	l[i] = value
}

func (l List) Slice(start, end int) []int {
	values := make([]int, 0)
	for i := start; i < end; i++ {
		values = append(values, l.Get(i))
	}
	return values
}

func (l List) Reverse(start, end int) {
	values := l.Slice(start, end)
	for i := start; i < end; i++ {
		value := values[end-i-1]
		l.Set(i, value)
	}
}

func compute(list List, lengths []int) {
	skip := 0
	start := 0
	for _, length := range lengths {
		list.Reverse(start, start+length)
		start += length + skip
		skip++
	}
}

func main() {
	lengths := []int{14, 58, 0, 116, 179, 16, 1, 104, 2, 254, 167, 86, 255, 55, 122, 244}
	list := make(List, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}
	compute(list, lengths)
	fmt.Printf("Value: %d\n", list[0]*list[1])
}
