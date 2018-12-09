package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Marble struct {
	previous *Marble
	next     *Marble
	value    int
	current  bool
}

func (m *Marble) Delete() {
	m.previous.next = m.next
	m.next.previous = m.previous
}

func (m *Marble) Rotate(delta int) *Marble {
	if delta < 0 {
		for ; delta < 0; delta++ {
			m = m.previous
		}
	} else {
		for ; delta > 0; delta-- {
			m = m.next
		}
	}
	return m
}

func (m *Marble) Add(marble int) (current *Marble, score int) {
	m.current = false
	if marble%23 == 0 {
		score += marble
		m = m.Rotate(-7)
		score += m.value
		m.Delete()
		m = m.next
		m.current = true
	} else {
		m = m.Rotate(1)
		newM := &Marble{
			previous: m,
			next:     m.next,
			value:    marble,
			current:  true,
		}
		m.next.previous = newM
		m.next = newM
		m = newM
	}

	return m, score
}

func (m *Marble) String() string {
	fields := []string{}
	start := m
	for {
		if m.current {
			fields = append(fields, fmt.Sprintf("(%d)", m.value))
		} else {
			fields = append(fields, fmt.Sprintf("%d", m.value))
		}
		m = m.next
		if m == start {
			break
		}
	}
	return strings.Join(fields, " ")
}

func aToi(str, errMsg string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", errMsg)
		os.Exit(-1)
	}
	return i
}

type Player struct {
	score int
}

func (p *Player) Less(other *Player) bool {
	return p.score < other.score
}

type Players struct {
	players []*Player
}

func (p *Players) Len() int           { return len(p.players) }
func (p *Players) Less(i, j int) bool { return p.players[i].Less(p.players[j]) }
func (p *Players) Swap(i, j int) {
	tmp := p.players[i]
	p.players[i] = p.players[j]
	p.players[j] = tmp
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <num players> <num marbles>\n", os.Args[0])
		os.Exit(-1)
	}
	numPlayers := aToi(os.Args[1], "num players must be a valid positive integer")
	numMarbles := aToi(os.Args[2], "num marbles must be a valid positive integer")

	score := 0
	current := &Marble{
		current: true,
	}
	current.next = current
	current.previous = current
	players := &Players{make([]*Player, numPlayers)}

	for numMarble := 1; numMarble < numMarbles; {
		for i, player := range players.players {
			if player == nil {
				players.players[i] = &Player{}
				player = players.players[i]
			}
			current, score = current.Add(numMarble)
			player.score += score
			numMarble++
			if numMarble >= numMarbles {
				break
			}
		}
	}
	sort.Sort(players)

	fmt.Printf("High Score: %d\n", players.players[len(players.players)-1].score)
}
