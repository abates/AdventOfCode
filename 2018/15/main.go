package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type Move struct {
	x int
	y int
}

type Path struct {
	length int
	first  Move
	last   Move
}

func (p *Path) Length() int {
	return p.length
}

func (p *Path) Copy() *Path {
	newPath := &Path{
		length: p.length,
		first:  p.first,
		last:   p.last,
	}
	return newPath
}

func (p *Path) Last() Move {
	return p.last
}

func (p *Path) Append(x, y int) *Path {
	if p.length == 1 {
		p.first = Move{x, y}
	}
	p.length++
	p.last = Move{x, y}
	return p
}

type Thing interface {
	String() string
}

type LayoutThing rune

func (lt LayoutThing) String() string { return string([]rune{rune(lt)}) }

var (
	Wall  = LayoutThing('#')
	Empty = LayoutThing('.')
)

type Player struct {
	t     rune
	x     int
	y     int
	power int
	life  int
	gm    *GameBoard
}

func (p *Player) String() string { return string([]rune{p.t}) }

func (p *Player) WithinRange() (enemy *Player) {
	for _, c := range [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
		candidate := p.gm.Get(p.x+c[0], p.y+c[1])
		if candidate != nil && candidate.t == p.t {
			continue
		}

		if candidate != nil && (enemy == nil || candidate.life < enemy.life) {
			enemy = candidate
		}
	}
	return
}

func (p *Player) Attack() {
	enemy := p.WithinRange()
	if enemy != nil {
		enemy.life -= p.power
		if enemy.life <= 0 {
			p.gm.Kill(enemy)
		}
	}
}

func (p *Player) Move(enemies []*Player) error {
	if len(enemies) == 0 {
		return fmt.Errorf("No enemies to attack")
	}

	if enemy := p.WithinRange(); enemy != nil {
		p.Attack()
		return nil
	}

	// find enemies within range
	min := -1
	var minPath *Path
	for _, enemy := range enemies {
		for _, delta := range [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
			if p.gm.layout[enemy.y+delta[1]][enemy.x+delta[0]] != Empty {
				continue
			}

			path := p.gm.ShortestPath(p.x, p.y, enemy.x+delta[0], enemy.y+delta[1])
			//fmt.Printf("Candidate: %+v\n", path)
			distance := path.Length()
			if distance <= 1 {
				continue
			}

			if minPath == nil || distance < min {
				min = distance
				minPath = path
			}
		}
	}

	if minPath != nil {
		//fmt.Printf("Winning Path: %+v\n", minPath)
		p.gm.Move(p, minPath.first.x, minPath.first.y)
		p.x = minPath.first.x
		p.y = minPath.first.y
		p.Attack()
	}

	return nil
}

type GameBoard struct {
	layout          [][]Thing
	defaultElfPower int
}

func (gm *GameBoard) UnmarshalText(text []byte) error {
	gm.layout = nil
	for y, line := range bytes.Split(text, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		runes := []rune(string(line))
		row := []Thing{}
		for x, r := range runes {
			if r == 'E' || r == 'G' {
				player := &Player{t: r, x: x, y: y, gm: gm, power: 3, life: 200}
				if r == 'E' {
					player.power = gm.defaultElfPower
				}
				row = append(row, player)
			} else {
				row = append(row, LayoutThing(r))
			}
		}
		gm.layout = append(gm.layout, row)
	}
	return nil
}

func (gm *GameBoard) Kill(player *Player) {
	gm.layout[player.y][player.x] = Empty
}

func (gm *GameBoard) ShortestPath(fromX, fromY, toX, toY int) (path *Path) {
	visited := make([][]bool, len(gm.layout))
	for i, row := range gm.layout {
		visited[i] = make([]bool, len(row))
	}

	// stack of x,y coordinates
	path = &Path{last: Move{fromX, fromY}, length: 1}
	pathStack := []*Path{
		path.Copy().Append(fromX, fromY-1),
		path.Copy().Append(fromX-1, fromY),
		path.Copy().Append(fromX+1, fromY),
		path.Copy().Append(fromX, fromY+1),
	}

	for len(pathStack) > 0 {
		path := pathStack[0]
		pathStack = pathStack[1:]
		cur := path.Last()

		if cur.x == toX && cur.y == toY {
			return path
		}

		if gm.layout[cur.y][cur.x] != Empty {
			continue
		}

		if visited[cur.y][cur.x] {
			continue
		}
		visited[cur.y][cur.x] = true

		pathStack = append(pathStack, path.Copy().Append(cur.x, cur.y-1))
		pathStack = append(pathStack, path.Copy().Append(cur.x-1, cur.y))
		pathStack = append(pathStack, path.Copy().Append(cur.x+1, cur.y))
		pathStack = append(pathStack, path.Copy().Append(cur.x, cur.y+1))
	}
	return &Path{}
}

func (gm *GameBoard) Get(x, y int) *Player {
	thing := gm.layout[y][x]
	if player, ok := thing.(*Player); ok {
		return player
	}
	return nil
}

func (gm *GameBoard) Move(player *Player, x, y int) {
	gm.layout[y][x] = player
	gm.layout[player.y][player.x] = Empty
	return
}

func (gm *GameBoard) players(t rune) (players []*Player) {
	for _, row := range gm.layout {
		for _, thing := range row {
			if player, ok := thing.(*Player); ok && player.t == t {
				players = append(players, player)
			}
		}
	}
	return players
}

func (gm *GameBoard) Goblins() []*Player {
	return gm.players('G')
}

func (gm *GameBoard) Elves() []*Player {
	return gm.players('E')
}

func (gm *GameBoard) Advance() (elves, goblins []*Player, err error) {
	moved := make(map[*Player]struct{})
	for _, row := range gm.layout {
		for _, thing := range row {
			if player, ok := thing.(*Player); ok {
				// don't move twice in the same turn
				if _, found := moved[player]; found {
					continue
				}

				//fmt.Printf("Player (%d,%d)\n", x, y)
				if player.t == 'E' {
					err = player.Move(gm.Goblins())
				} else {
					err = player.Move(gm.Elves())
				}
				moved[player] = struct{}{}
				if err != nil {
					return gm.Elves(), gm.Goblins(), err
				}
			}
		}
	}
	return gm.Elves(), gm.Goblins(), nil
}

func hitPoints(players []*Player) int {
	sum := 0
	for _, player := range players {
		if player.life > 0 {
			sum += player.life
		}
	}
	return sum
}

func (gm *GameBoard) Play() (round, points int, winner string) {
	for round = 1; ; round++ {
		//fmt.Printf("%s\n", gm.String())
		elves, goblins, err := gm.Advance()
		if err != nil {
			round--
		}

		if len(elves) == 0 {
			winner = "Goblins"
			points = hitPoints(goblins)
			break
		} else if len(goblins) == 0 {
			winner = "Elves"
			points = hitPoints(elves)
			break
		}
	}
	return
}

func (gm *GameBoard) String() string {
	runes := []rune{}
	for _, row := range gm.layout {
		for _, thing := range row {
			if player, ok := thing.(*Player); ok {
				runes = append(runes, player.t)
			} else if lt, ok := thing.(LayoutThing); ok {
				runes = append(runes, rune(lt))
			}
		}
		runes = append(runes, '\n')
	}
	return string(runes)
}

func part1(input []byte) error {
	gm := &GameBoard{defaultElfPower: 3}
	err := gm.UnmarshalText(input)
	if err == nil {
		round, points, winner := gm.Play()
		fmt.Printf("%v\n", gm.String())
		fmt.Printf("Combat ends after %d full rounds\n", round)
		fmt.Printf("%s win with %d total hit points left\n", winner, points)
		fmt.Printf("Outcome: %d x %d = %d\n", round, points, round*points)
	}
	return err
}

func part2(input []byte) error {
	for power := 4; ; power++ {
		fmt.Printf("Power %d\n", power)
		gm := &GameBoard{defaultElfPower: power}
		err := gm.UnmarshalText(input)
		startElves := len(gm.Elves())
		if err == nil {
			round, points, winner := gm.Play()
			if winner == "Elves" && len(gm.Elves()) == startElves {
				fmt.Printf("%v\n", gm.String())
				fmt.Printf("Combat ends after %d full rounds\n", round)
				fmt.Printf("%s win with %d total hit points left\n", winner, points)
				fmt.Printf("Outcome: %d x %d = %d\n", round, points, round*points)
				return nil
			}
		} else {
			return err
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file>\n", os.Args[0])
		os.Exit(-1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(-1)
	}
	err = part1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to perform part 1: %v\n", err)
		os.Exit(-1)
	}

	err = part2(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to perform part 2: %v\n", err)
		os.Exit(-1)
	}
}
