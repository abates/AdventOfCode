package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	registers map[string]int
	lastSound int
}

func (p *Player) getValue(input string) int {
	if value, err := strconv.Atoi(input); err == nil {
		return value
	}
	return p.registers[input]
}

func (p *Player) Snd(x int) {
	p.lastSound = x
}

func (p *Player) Set(x string, y int) {
	p.registers[x] = y
}

func (p *Player) Add(x string, y int) {
	p.registers[x] += y
}

func (p *Player) Mul(x string, y int) {
	p.registers[x] *= y
}

func (p *Player) Mod(x string, y int) {
	p.registers[x] = p.registers[x] % y
}

func (p *Player) Rcv(x int) bool {
	if x != 0 {
		fmt.Printf("Recover %d\n", p.lastSound)
	}
	return x != 0
}

func (p *Player) play(song []string) {
	pc := 0
loop:
	for 0 <= pc && pc < len(song) {
		fields := strings.Fields(song[pc])
		switch fields[0] {
		case "snd":
			p.Snd(p.getValue(fields[1]))
		case "set":
			p.Set(fields[1], p.getValue(fields[2]))
		case "add":
			p.Add(fields[1], p.getValue(fields[2]))
		case "mul":
			p.Mul(fields[1], p.getValue(fields[2]))
		case "mod":
			p.Mod(fields[1], p.getValue(fields[2]))
		case "rcv":
			if p.Rcv(p.getValue(fields[1])) {
				break loop
			}
		case "jgz":
			if p.getValue(fields[1]) > 0 {
				pc += p.getValue(fields[2])
				continue
			}
		}
		pc++
	}
}

func play(input string) {
	song := []string{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		song = append(song, strings.TrimSpace(line))
	}

	player := &Player{registers: make(map[string]int)}
	player.play(song)
}

func main() {
	//test := "set a 1\nadd a 2\nmul a a\nmod a 5\nsnd a\nset a 0\nrcv a\njgz a -1\nset a 1\njgz a -2"
	//play(test)
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	play(string(b))
}
