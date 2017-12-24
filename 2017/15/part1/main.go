package main

import "fmt"

type Generator struct {
	factor   int
	multiple int
	divisor  int
	valueCh  chan int
}

func NewGenerator(seed, factor, multiple int) *Generator {
	gen := &Generator{
		valueCh: make(chan int),
	}

	go gen.generate(seed, factor, multiple)
	return gen
}

func (gen *Generator) generate(seed, factor, multiple int) {
	divisor := 2147483647
	for {
		seed = (seed * factor) % divisor
		if seed%multiple == 0 {
			gen.valueCh <- seed
		}
	}
}

func (gen *Generator) Next() int {
	return <-gen.valueCh
}

func Match(value1, value2 int) bool {
	return (value1 & 0xffff) == (value2 & 0xffff)
}

func CountMatches(genA, genB *Generator, limit int) int {
	count := 0
	for i := 0; i < limit; i++ {
		value1 := genA.Next()
		value2 := genB.Next()
		if Match(value1, value2) {
			count++
		}
	}
	return count
}

func main() {
	// part 1
	/*genA := NewGenerator(699, 16807, 1)
	genB := NewGenerator(124, 48271, 1)
	count := CountMatches(genA, genB, 40000000)
	fmt.Printf("Count: %d\n", count)*/

	// part 2
	genA := NewGenerator(699, 16807, 4)
	genB := NewGenerator(124, 48271, 8)
	count := CountMatches(genA, genB, 5000000)
	fmt.Printf("Count: %d\n", count)
}
