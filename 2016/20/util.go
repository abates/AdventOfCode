package main

import (
	"fmt"
)

type Range struct {
	low  int
	high int
}

func (r *Range) Merge(other *Range) {
	if other.low < r.low {
		r.low = other.low
	}

	if other.high > r.high {
		r.high = other.high
	}
}

func (r *Range) Adjacent(other *Range) bool {
	if r.Compare(other) == 0 {
		return true
	}

	return other.high == r.low-1 || other.low == r.high+1
}

func (r *Range) Compare(other *Range) int {
	if other.low < r.low && other.high < r.low {
		return -1
	}

	if other.low > r.high && other.high > r.high {
		return 1
	}

	return 0
}

func (r *Range) Contains(other *Range) bool {
	return other.low >= r.low && other.high <= r.high
}

func (r *Range) Count() int {
	return r.high - r.low + 1
}

func (r *Range) String() string {
	return fmt.Sprintf("(%d,%d)", r.low, r.high)
}

type Ranges struct {
	ranges []*Range
}

func (r *Ranges) Insert(other *Range) {
	if len(r.ranges) == 0 {
		r.ranges = []*Range{other}
		return
	}

	left := 0
	right := len(r.ranges) - 1
	cmp := 0
	index := 0
	for left <= right {
		index = (left + right) / 2
		search := r.ranges[index]
		cmp = search.Compare(other)
		if cmp < 0 {
			right = index - 1
		} else if cmp > 0 {
			left = index + 1
		} else {
			break
		}
	}

	if cmp < 0 {
		r.ranges = append(r.ranges[:index], append([]*Range{other}, r.ranges[index:]...)...)
	} else if cmp > 0 {
		r.ranges = append(r.ranges[:index+1], append([]*Range{other}, r.ranges[index+1:]...)...)
		index++
	} else {
		r.ranges[index].Merge(other)
	}

	if index > 0 && r.ranges[index].Adjacent(r.ranges[index-1]) {
		r.ranges[index].Merge(r.ranges[index-1])
		r.ranges = append(r.ranges[:index-1], r.ranges[index:]...)
		index--
	}

	if index < len(r.ranges)-1 && r.ranges[index].Adjacent(r.ranges[index+1]) {
		r.ranges[index].Merge(r.ranges[index+1])
		r.ranges = append(r.ranges[:index+1], r.ranges[index+2:]...)
	}
}

func (r *Ranges) LowestException() int {
	if r.ranges[0].low > 0 {
		return 0
	}

	return r.ranges[0].high + 1
}
