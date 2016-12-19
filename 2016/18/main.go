package main

import (
	"fmt"
)

func main() {
	startRow := NewRowFromString("......^.^^.....^^^^^^^^^...^.^..^^.^^^..^.^..^.^^^.^^^^..^^.^.^.....^^^^^..^..^^^..^^.^.^..^^..^^^..")
	row := startRow
	count := 0
	for i := 0; i < 40; i++ {
		count += row.NumSafe()
		row = row.NextRow()
	}

	fmt.Printf("Num safe: %d\n", count)

	row = startRow
	count = 0
	for i := 0; i < 400000; i++ {
		count += row.NumSafe()
		row = row.NextRow()
	}

	fmt.Printf("Num safe: %d\n", count)
}
