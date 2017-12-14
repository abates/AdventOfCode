package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/abates/AdventOfCode/util"
)

type Packet int

func (p *Packet) Penalize(penalty int) {
	*p += Packet(penalty)
}

type Layer struct {
	depth int
	rng   int
	step  int
}

func (l *Layer) Reset() {
	l.step = 0
}

func (l *Layer) Step() {
	l.step++
}

func (l *Layer) Position() int {
	index := l.step % ((l.rng - 1) * 2)
	midpoint := l.rng - 1
	//println("Index", index, "midpoint", midpoint, "midpoint - index", (midpoint - index))
	return midpoint - util.Abs(midpoint-index)
}

func (l *Layer) Penalty() int {
	return l.depth * l.rng
}

func (l *Layer) String() string {
	if l == nil {
		return "..."
	}

	var buffer bytes.Buffer
	for i := 0; i < l.rng; i++ {
		if i == l.Position() {
			buffer.WriteString("[S]")
		} else {
			buffer.WriteString("[ ]")
		}
	}
	return buffer.String()
}

type Firewall struct {
	layers []*Layer
}

func (fw *Firewall) String() string {
	var buffer bytes.Buffer
	for i, layer := range fw.layers {
		buffer.WriteString(fmt.Sprintf("%-2d %s\n", i, layer.String()))
	}
	return buffer.String()
}

func (fw *Firewall) Step() {
	for _, layer := range fw.layers {
		if layer != nil {
			layer.Step()
		}
	}
}

func (fw *Firewall) Traverse(packet *Packet) bool {
	permit := true
	for i := 0; i < len(fw.layers); i++ {
		if fw.layers[i] != nil && fw.layers[i].Position() == 0 {
			packet.Penalize(fw.layers[i].Penalty())
			permit = false
		}
		fw.Step()
	}
	return permit
}

func (fw *Firewall) Delay(delay int) {
	for _, layer := range fw.layers {
		if layer != nil {
			layer.step = delay
		}
	}
}

func (fw *Firewall) TraversalDelay() int {
	for i := 0; ; i++ {
		fw.Reset()
		fw.Delay(i)
		pkt := Packet(0)
		if fw.Traverse(&pkt) {
			return i
		}
	}
	return 0
}

func (fw *Firewall) Reset() {
	for _, layer := range fw.layers {
		if layer != nil {
			layer.Reset()
		}
	}
}

func (fw *Firewall) AddLayer(depth, rng int) {
	if depth > len(fw.layers)-1 {
		fw.layers = append(fw.layers, make([]*Layer, depth-len(fw.layers)+1)...)
	}
	fw.layers[depth] = &Layer{depth: depth, rng: rng}
}

func initializeFirewall(input string) *Firewall {
	firewall := &Firewall{}
	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")
		depth, _ := strconv.Atoi(strings.TrimSpace(fields[0]))
		rng, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
		firewall.AddLayer(depth, rng)
	}
	return firewall
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	fw := initializeFirewall(string(b))

	pkt := Packet(0)
	fw.Traverse(&pkt)

	fmt.Printf("Total penalty: %d\n", int(pkt))
	fmt.Printf("Traversal delay: %d\n", fw.TraversalDelay())
}
