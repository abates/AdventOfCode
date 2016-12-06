package main

import (
	"bytes"
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strconv"
	"strings"
)

type Room struct {
	encryptedName string
	sectorId      int
	checksum      string
}

func NewRoom(encryptedName string, sectorId int, checksum string) *Room {
	return &Room{encryptedName, sectorId, checksum}
}

func (r *Room) Valid() bool {
	letters := util.NewLetters()
	for _, s := range r.encryptedName {
		if string(s) == "-" {
			continue
		}
		letters.Add(string(s))
	}

	top5 := strings.Join(letters.Sorted()[0:5], "")
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
