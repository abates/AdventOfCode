package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

func main() {
	for _, line := range util.ReadInput() {
		room := parseRoom(line)
		if room.Valid() {
			fmt.Printf("%d:%v\n", room.sectorId, room.DecryptedName())
		}
	}
}
