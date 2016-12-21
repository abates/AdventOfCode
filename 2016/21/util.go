package main

import (
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
)

type Password struct {
	str string
}

func NewPassword(input string) *Password {
	return &Password{input}
}

func sortXY(x *int, y *int) {
	if *x > *y {
		t := *x
		*x = *y
		*y = t
	}
}

func (p *Password) Swap(x, y int) {
	sortXY(&x, &y)
	writer := &util.StringWriter{}

	writer.Write(p.str[:x])
	writer.WriteByte(p.str[y])
	writer.Write(p.str[x+1 : y])
	writer.WriteByte(p.str[x])
	writer.Write(p.str[y+1:])

	p.str = writer.String()
}

func (p *Password) SwapLetter(x, y string) {
	writer := &util.StringWriter{}

	for _, l := range strings.Split(p.str, "") {
		if l == x {
			writer.Write(y)
		} else if l == y {
			writer.Write(x)
		} else {
			writer.Write(l)
		}
	}

	p.str = writer.String()
}

func (p *Password) RotateLeft(x int) {
	x = x % len(p.str)
	p.str = p.str[x:] + p.str[:x]
}

func (p *Password) RotateRight(x int) {
	x = x % len(p.str)
	p.str = p.str[len(p.str)-x:] + p.str[:len(p.str)-x]
}

func (p *Password) RotatePosition(x string) {
	index := strings.Index(p.str, x)
	if index < 4 {
		p.RotateRight(index + 1)
	} else {
		p.RotateRight(index + 2)
	}
}

func (p *Password) Reverse(x, y int) {
	sortXY(&x, &y)

	writer := &util.StringWriter{}
	writer.Write(p.str[:x])

	letters := strings.Split(p.str[x:y+1], "")
	for i := len(letters) - 1; i >= 0; i-- {
		writer.Write(letters[i])
	}
	writer.Write(p.str[y+1:])
	p.str = writer.String()
}

func (p *Password) Move(x, y int) {
	str := p.str[:x] + p.str[x+1:]
	writer := &util.StringWriter{}
	writer.Write(str[:y])
	writer.WriteByte(p.str[x])
	writer.Write(str[y:])
	p.str = writer.String()
}

func (p *Password) String() string {
	return p.str
}
