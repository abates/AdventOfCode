package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	Name string
}

type Pipeline struct {
	Programs []*Program
	index    map[string]int
}

func (p *Pipeline) Spin(x int) {
	p.Programs = append(p.Programs[len(p.Programs)-x:], p.Programs[0:len(p.Programs)-x]...)
	for i, program := range p.Programs {
		p.index[program.Name] = i
	}
}

func (p *Pipeline) Exchange(a, b int) {
	ap := p.Programs[a]
	bp := p.Programs[b]
	p.Programs[a] = bp
	p.Programs[b] = ap
	p.index[ap.Name] = b
	p.index[bp.Name] = a
}

func (p *Pipeline) Partner(a, b string) {
	ai := p.index[a]
	bi := p.index[b]
	ap := p.Programs[ai]
	bp := p.Programs[bi]
	p.Programs[ai] = bp
	p.Programs[bi] = ap
	p.index[a] = bi
	p.index[b] = ai
}

func (p *Pipeline) String() string {
	var buf bytes.Buffer
	for _, p := range p.Programs {
		buf.WriteString(p.Name)
	}
	return buf.String()
}

func buildPipeline(length int) *Pipeline {
	p := &Pipeline{
		Programs: make([]*Program, length),
		index:    make(map[string]int),
	}

	for i := 0; i < length; i++ {
		p.Programs[i] = &Program{Name: fmt.Sprintf("%c", 97+i)}
		p.index[p.Programs[i].Name] = i
	}

	return p
}

func dance(input string, pipeline *Pipeline, iterations int) {
	seen := make(map[string][]int)
	moves := strings.Split(strings.TrimSpace(input), ",")
	for i := 0; i < iterations; i++ {
		if position, found := seen[pipeline.String()]; found {
			if len(position) == 1 {
				position = []int{position[0], i}
				seen[pipeline.String()] = position
			}

			delta := position[1] - position[0]
			delta = delta * ((iterations - i) / delta)
			if delta > 0 && i+delta < iterations {
				i += delta
				continue
			}
		} else {
			seen[pipeline.String()] = []int{i}
		}

		for _, move := range moves {
			switch move[0] {
			case 's':
				x, _ := strconv.Atoi(move[1:])
				pipeline.Spin(x)
			case 'x':
				fields := strings.Split(move[1:], "/")
				a, _ := strconv.Atoi(fields[0])
				b, _ := strconv.Atoi(fields[1])
				pipeline.Exchange(a, b)
			case 'p':
				fields := strings.Split(move[1:], "/")
				pipeline.Partner(fields[0], fields[1])
			}
		}
	}
	fmt.Println()
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)

	pipeline := buildPipeline(16)
	dance(string(b), pipeline, 1)

	fmt.Printf("Part 1: %s\n", pipeline)

	dance(string(b), pipeline, 1000000000)
	fmt.Printf("Part 2: %s\n", pipeline)
}
