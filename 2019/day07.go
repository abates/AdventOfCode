package main

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

func init() {
	d7 := &D7{}
	challenges = append(challenges, &challenge{"Day 07", "input/day07.txt", nil, d7.parseFile, d7.part1, d7.part2})
}

type D7 struct {
	mem []*Int
}

func (d7 *D7) parseFile(lines []string) (err error) {
	d7.mem, err = ParseComputerMemory(lines)
	return err
}

func getCombos(length, l, h int, values chan<- []int) {
	var f func(int, int, int, []int)
	index := make(map[int]bool)
	f = func(pos, l, h int, accum []int) {
		accum = append(accum, 0)
		for i := l; i <= h; i++ {
			if _, found := index[i]; found {
				continue
			}
			index[i] = true
			accum[len(accum)-1] = i
			if pos == 0 {
				values <- accum
			} else {
				f(pos-1, l, h, accum)
			}
			delete(index, i)
		}
	}

	f(length, l, h, []int{})
	close(values)
}

func (d7 *D7) tryCombo(feedback bool, combo []int) (int, error) {
	amps := make([]*Computer, len(combo))
	outputs := make([]io.Reader, len(combo))
	var wg sync.WaitGroup

	for j := 0; j < len(amps); j++ {
		amps[j] = NewComputer(d7.mem)
		amps[j].name = fmt.Sprintf("Amp %d", j)
		if j == 0 {
			amps[j].SetInput(strings.NewReader(fmt.Sprintf("%d\n0\n", combo[j])))
		} else {
			amps[j].SetInput(io.MultiReader(strings.NewReader(fmt.Sprintf("%d\n", combo[j])), outputs[j-1]))
		}

		pr, pw := io.Pipe()
		outputs[j] = pr
		amps[j].SetOutput(pw)
	}

	if feedback {
		amps[0].SetInput(io.MultiReader(amps[0].input, outputs[len(outputs)-1]))
	}

	for i, amp := range amps {
		if i < len(amps)-1 {
			wg.Add(1)
		}
		go func(i int, amp *Computer) {
			err := amp.Run()
			if err != nil {
				panic(err.Error())
			}
			if i < len(amps)-1 {
				wg.Done()
			}
		}(i, amp)
	}

	wg.Wait()
	output := 0
	_, err := fmt.Fscanf(outputs[len(outputs)-1], "%d", &output)
	return output, err
}

func (d7 *D7) run(low, high int, feedback bool) (string, error) {
	max := 0
	ch := make(chan []int)
	go getCombos(4, low, high, ch)
	for combo := range ch {
		output, err := d7.tryCombo(feedback, combo)
		if err != nil {
			panic(err.Error())
		}
		if max < output {
			max = output
		}
	}
	return fmt.Sprintf("Max Thruster Signal: %d", max), nil
}

func (d7 *D7) part1() (string, error) {
	return d7.run(0, 4, false)
}

func (d7 *D7) part2() (string, error) {
	return d7.run(5, 9, true)
}
