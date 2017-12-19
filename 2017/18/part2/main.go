package main

// This is a terrible solution :-/
// not proud

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Queue struct {
	queue []int
	mutex *sync.Mutex
	cond  *sync.Cond
}

func NewQueue() *Queue {
	q := &Queue{mutex: &sync.Mutex{}}
	cond := sync.NewCond(q.mutex)
	q.cond = cond
	return q
}

func (q *Queue) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.queue)
}

func (q *Queue) Write(i int) {
	q.mutex.Lock()
	q.queue = append(q.queue, i)
	q.mutex.Unlock()
	q.cond.Broadcast()
}

func (q *Queue) Read() int {
	q.mutex.Lock()
	if len(q.queue) == 0 {
		q.cond.Wait()
	}
	i := q.queue[0]
	q.queue = q.queue[1:]
	q.mutex.Unlock()
	return i
}

type Player struct {
	name       string
	registers  map[string]int
	readQueue  *Queue
	writeQueue *Queue
	sndCount   int
}

func (p *Player) getValue(input string) int {
	if value, err := strconv.Atoi(input); err == nil {
		return value
	}
	return p.registers[input]
}

func (p *Player) Snd(x int) bool {
	p.writeQueue.Write(x)
	p.sndCount++
	fmt.Printf("%s %d\n", p.name, p.sndCount)
	return true
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

func (p *Player) Rcv(x string) bool {
	p.registers[x] = p.readQueue.Read()
	return true
}

func (p *Player) play(song []string) {
	pc := 0
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
			p.Rcv(fields[1])
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
	wg := &sync.WaitGroup{}
	song := []string{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		song = append(song, strings.TrimSpace(line))
	}

	q1 := NewQueue()
	q2 := NewQueue()

	p1 := &Player{name: "p0", registers: map[string]int{"p": 0}, readQueue: q1, writeQueue: q2}
	p2 := &Player{name: "p1", registers: map[string]int{"p": 1}, readQueue: q2, writeQueue: q1}

	go p1.play(song)
	go p2.play(song)
	wg.Add(2)
	wg.Wait()

	fmt.Printf("P1 sent %d times\n", p1.sndCount)
}

func main() {
	//test := "snd 1\nsnd 2\nsnd p\nrcv a\nrcv b\nrcv c\nrcv d"
	//play(test)

	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	play(string(b))
}
