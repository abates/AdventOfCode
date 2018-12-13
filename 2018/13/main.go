package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Turn int

const (
	TURN_RIGHT Turn = iota
	TURN_LEFT
	STRAIGHT
)

type Direction rune

const (
	DIRECTION_UP    Direction = '^'
	DIRECTION_RIGHT           = '>'
	DIRECTION_DOWN            = 'v'
	DIRECTION_LEFT            = '<'
)

type Cart struct {
	x         int
	y         int
	layout    *Layout
	lastTurn  Turn
	direction Direction
}

func (cart *Cart) turnLeft() {
	switch cart.direction {
	case DIRECTION_UP:
		cart.direction = DIRECTION_LEFT
	case DIRECTION_RIGHT:
		cart.direction = DIRECTION_UP
	case DIRECTION_DOWN:
		cart.direction = DIRECTION_RIGHT
	case DIRECTION_LEFT:
		cart.direction = DIRECTION_DOWN
	}
}

func (cart *Cart) turnRight() {
	switch cart.direction {
	case DIRECTION_UP:
		cart.direction = DIRECTION_RIGHT
	case DIRECTION_RIGHT:
		cart.direction = DIRECTION_DOWN
	case DIRECTION_DOWN:
		cart.direction = DIRECTION_LEFT
	case DIRECTION_LEFT:
		cart.direction = DIRECTION_UP
	}
}

func (cart *Cart) advance() error {
	x := cart.x
	y := cart.y
	switch cart.direction {
	case DIRECTION_UP:
		y -= 1
	case DIRECTION_DOWN:
		y += 1
	case DIRECTION_LEFT:
		x -= 1
	case DIRECTION_RIGHT:
		x += 1
	}
	err := cart.layout.MoveTo(cart, x, y)
	if err == nil {
		cart.x = x
		cart.y = y
	}
	return err
}

func (cart *Cart) Advance() error {
	position := cart.layout.Position(cart.x, cart.y)
	switch {
	case position.track == '|' || position.track == '-':
	case position.track == '/':
		if cart.direction == DIRECTION_UP || cart.direction == DIRECTION_DOWN {
			cart.turnRight()
		} else {
			cart.turnLeft()
		}
	case position.track == '\\':
		if cart.direction == DIRECTION_UP || cart.direction == DIRECTION_DOWN {
			cart.turnLeft()
		} else {
			cart.turnRight()
		}
	case position.track == '+':
		switch cart.lastTurn {
		case TURN_LEFT:
			cart.lastTurn = STRAIGHT
		case STRAIGHT:
			cart.turnRight()
			cart.lastTurn = TURN_RIGHT
		case TURN_RIGHT:
			cart.turnLeft()
			cart.lastTurn = TURN_LEFT
		}
	}
	return cart.advance()
}

func (cart *Cart) String() string {
	return string([]rune{rune(cart.direction)})
}

type Position struct {
	cart  *Cart
	track rune
}

func (p *Position) String() string {
	if p.cart == nil {
		return string([]rune{p.track})
	}
	return p.cart.String()
}

type Layout struct {
	positions [][]*Position
}

func (l *Layout) UnmarshalText(text []byte) (err error) {
	for y, line := range bytes.Split(text, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		positions := []*Position{}
		runes := []rune(string(line))
		for x, r := range runes {
			position := &Position{}

			switch {
			case r == 'v' || r == '^':
				position.track = '|'
				position.cart = &Cart{layout: l, direction: Direction(r), x: x, y: y}
			case r == '<' || r == '>':
				position.track = '-'
				position.cart = &Cart{layout: l, direction: Direction(r), x: x, y: y}
			default:
				position.track = r
			}
			positions = append(positions, position)
		}
		l.positions = append(l.positions, positions)
	}
	return err
}

func (l *Layout) Position(x, y int) *Position {
	return l.positions[y][x]
}

type ErrCrash struct {
	x     int
	y     int
	cart1 *Cart
	cart2 *Cart
}

func (err ErrCrash) Error() string {
	return fmt.Sprintf("Crash at %d,%d", err.x, err.y)
}

func (l *Layout) MoveTo(cart *Cart, x, y int) (err error) {
	l.positions[cart.y][cart.x].cart = nil
	if l.positions[y][x].cart != nil {
		// CRASH!
		err = &ErrCrash{
			cart1: l.positions[y][x].cart,
			cart2: cart,
			x:     x,
			y:     y,
		}
		l.positions[y][x].cart = nil
	} else {
		l.positions[y][x].cart = cart
	}
	return err
}

func (l *Layout) Advance(errOnCrash bool) (carts []*Cart, err error) {
	moves := make(map[*Cart]struct{})
	for _, row := range l.positions {
		for _, position := range row {
			cart := position.cart
			if _, found := moves[cart]; cart != nil && !found {
				if err = position.cart.Advance(); err != nil {
					if errOnCrash {
						return
					}
					crashErr := err.(*ErrCrash)
					delete(moves, crashErr.cart1)
					delete(moves, crashErr.cart2)
				} else {
					moves[cart] = struct{}{}
				}
			}
		}
	}

	for cart := range moves {
		carts = append(carts, cart)
	}
	return
}

func (l *Layout) String() string {
	var builder strings.Builder
	for _, row := range l.positions {
		for _, position := range row {
			builder.WriteString(position.String())
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func part1(input []byte) {
	layout := &Layout{}
	err := layout.UnmarshalText(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse input: %v", err)
		os.Exit(-1)
	}

	for {
		_, err = layout.Advance(true)
		if err != nil {
			fmt.Printf("Part1: %v\n", err)
			return
		}
	}
}

func part2(input []byte) {
	layout := &Layout{}
	err := layout.UnmarshalText(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse input: %v", err)
		os.Exit(-1)
	}

	for {
		carts, _ := layout.Advance(false)
		if len(carts) == 1 {
			fmt.Printf("Part 2: %d,%d\n", carts[0].x, carts[0].y)
			return
		}
	}
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
	part1(input)
	part2(input)
}
