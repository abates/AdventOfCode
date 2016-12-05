package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
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

type Room struct {
	encryptedName string
	sectorId      int
	checksum      string
}

func NewRoom(encryptedName string, sectorId int, checksum string) *Room {
	return &Room{encryptedName, sectorId, checksum}
}

func (r *Room) Valid() bool {
	letters := NewLetters()
	for _, s := range r.encryptedName {
		if string(s) == "-" {
			continue
		}
		letters.Add(string(s))
	}

	sort.Sort(letters)
	top5 := strings.Join(letters.keys[0:5], "")
	return bytes.Equal([]byte(top5), []byte(r.checksum))
}

func (r *Room) DecryptedName() string {
	s := ""
	offset := r.sectorId % 26
	for _, l := range r.encryptedName {
		j := int(l) - 97
		if l == '-' {
			s = fmt.Sprintf("%s ", s)
		} else if j+offset < 26 {
			s = fmt.Sprintf("%s%c", s, 97+j+offset)
		} else {
			s = fmt.Sprintf("%s%c", s, 97+j+offset-26)
		}
	}

	return s
}

func parseRoom(line string) *Room {
	openBracket := strings.Index(line, "[")
	closeBracket := strings.Index(line, "]")
	lastDash := strings.LastIndex(line[:openBracket], "-")

	words := line[:lastDash]
	sectorId, _ := strconv.Atoi(line[lastDash+1 : openBracket])
	checksum := line[openBracket+1 : closeBracket]

	return NewRoom(words, sectorId, checksum)
}
