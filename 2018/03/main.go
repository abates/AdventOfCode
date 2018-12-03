package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Claim struct {
	id     int
	startX int
	startY int
	endX   int
	endY   int
}

func NewClaim(def string) *Claim {
	c := &Claim{}
	width := 0
	height := 0
	fmt.Sscanf(def, "#%d @ %d,%d: %dx%d", &c.id, &c.startX, &c.startY, &width, &height)
	c.endX = c.startX + width
	c.endY = c.startY + height
	return c
}

type Fabric struct {
	grid   [][][]*Claim
	claims map[*Claim]struct{}
}

func NewFabric() *Fabric {
	f := &Fabric{
		grid:   make([][][]*Claim, 1000),
		claims: make(map[*Claim]struct{}),
	}

	for i := range f.grid {
		f.grid[i] = make([][]*Claim, 1000)
	}

	return f
}

func (f *Fabric) AddClaim(claim *Claim) {
	f.claims[claim] = struct{}{}
	for y := claim.startY; y < claim.endY; y++ {
		for x := claim.startX; x < claim.endX; x++ {
			f.grid[y][x] = append(f.grid[y][x], claim)
		}
	}
}

func (f *Fabric) CountOverlap() (int, *Claim) {
	var nonOverlapping *Claim

	// copy the map
	claims := make(map[*Claim]struct{})
	for claim := range f.claims {
		claims[claim] = struct{}{}
	}

	overlap := 0
	for y := 0; y < len(f.grid); y++ {
		for x := 0; x < len(f.grid[y]); x++ {
			if len(f.grid[y][x]) > 1 {
				overlap++
				for _, claim := range f.grid[y][x] {
					delete(claims, claim)
				}
			}
		}
	}

	if len(claims) != 1 {
		panic("Fatal error, should have only one non-overlapping claim")
	}

	for claim := range claims {
		nonOverlapping = claim
	}

	return overlap, nonOverlapping
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	fabric := NewFabric()

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		claim := NewClaim(line)
		fabric.AddClaim(claim)
	}

	overlapping, nonOverlapping := fabric.CountOverlap()
	fmt.Printf("Overlapping: %d NonOverlapping %d\n", overlapping, nonOverlapping.id)
}
