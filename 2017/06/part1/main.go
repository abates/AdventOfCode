package main

import (
	"fmt"
	"strings"
)

type Bank int

func (b *Bank) Clear() int {
	value := int(*b)
	*b = 0
	return value
}

func (b *Bank) Add() {
	*b++
}

func (b *Bank) String() string { return fmt.Sprintf("%d", int(*b)) }

type Memory struct {
	banks []Bank
}

func NewMemory(blockSizes []int) *Memory {
	mem := &Memory{}
	for _, size := range blockSizes {
		mem.banks = append(mem.banks, Bank(size))
	}
	return mem
}

func (m *Memory) LargestBankIndex() int {
	index := 0
	for i, bank := range m.banks[1:] {
		if m.banks[index] < bank {
			index = i + 1
		}
	}
	return index
}

func (m *Memory) Redistribute() {
	index := m.LargestBankIndex()
	blocksize := m.banks[index].Clear()
	for ; blocksize > 0; blocksize-- {
		index++
		if index == len(m.banks) {
			index = 0
		}
		m.banks[index].Add()
	}
}

func (m *Memory) String() string {
	banks := make([]string, 0)
	for _, bank := range m.banks {
		banks = append(banks, fmt.Sprintf("%d", bank))
	}
	return strings.Join(banks, ",")
}

func (m *Memory) Cycle() int {
	steps := 0
	seen := make(map[string]bool)

	for {
		str := m.String()
		if _, found := seen[str]; found {
			break
		}
		steps++
		seen[str] = true
		m.Redistribute()
	}
	return steps
}

func main() {
	input := []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	mem := NewMemory(input)
	steps := mem.Cycle()

	fmt.Printf("Finished in %d steps\n", steps)

	steps = mem.Cycle()
	fmt.Printf("Next cycle after %d steps\n", steps)
}
