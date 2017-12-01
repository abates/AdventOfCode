package util

type circularIntList struct {
	ints []int
	pos  int
}

type CircularIntList interface {
	Add(int)
	Peek() int
	PeekN(int) int
	Next() int
}

func NewCircularIntList() CircularIntList {
	return &circularIntList{pos: -1}
}

func (cil *circularIntList) Add(i int) {
	cil.ints = append(cil.ints, i)
}

func (cil *circularIntList) PeekN(n int) int {
	pos := cil.pos + n
	if pos >= len(cil.ints) {
		pos = pos % len(cil.ints)
	}
	return cil.ints[pos]
}

func (cil *circularIntList) Peek() int {
	/*pos := cil.pos + 1
	if pos >= len(cil.ints) {
		pos = 0
	}
	return cil.ints[pos]*/
	return cil.PeekN(1)
}

func (cil *circularIntList) Next() int {
	cil.pos++
	if cil.pos >= len(cil.ints) {
		cil.pos = 0
	}
	i := cil.ints[cil.pos]
	return i
}
