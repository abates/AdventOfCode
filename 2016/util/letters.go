package util

import (
	"fmt"
	"sort"
	"strings"
)

type Letters struct {
	counts map[string]int
	keys   []string
}

func NewLetters() *Letters {
	return &Letters{
		counts: make(map[string]int),
	}
}

func (l *Letters) Add(s string) {
	if l.counts[s] == 0 {
		l.keys = append(l.keys, s)
	}
	l.counts[s]++
}

func (l *Letters) Len() int { return len(l.counts) }

func (l *Letters) Less(i, j int) bool {
	if l.counts[l.keys[i]] > l.counts[l.keys[j]] {
		return true
	}

	if l.counts[l.keys[i]] == l.counts[l.keys[j]] {
		return strings.Compare(l.keys[i], l.keys[j]) == -1
	}

	return false
}

func (l *Letters) Swap(i, j int) {
	temp := l.keys[i]
	l.keys[i] = l.keys[j]
	l.keys[j] = temp
}

func (l *Letters) String() string {
	s := ""

	for _, k := range l.keys {
		s += fmt.Sprintf("%s:%d ", k, l.counts[k])
	}

	return s
}

func (l *Letters) Sorted() []string {
	sort.Sort(l)
	return l.keys
}

func (l *Letters) Count(key string) int {
	return l.counts[key]
}
