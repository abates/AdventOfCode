package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

func isTriangle(e1, e2, e3 int) bool {
	return e1+e2 > e3 && e1+e3 > e2 && e2+e3 > e1
}

func readInput() [][]int {
	rows := make([][]int, 0)
	for _, line := range util.ReadInput() {
		columns := make([]int, 3)
		fmt.Sscanf(line, "%d %d %d", &columns[0], &columns[1], &columns[2])
		rows = append(rows, columns)
	}
	return rows
}
