package main

type Wheel struct {
	position  int
	divisions int
}

func (w *Wheel) PeekAdvance(n int) int {
	position := (w.position + n) % w.divisions
	return position
}

func (w *Wheel) Advance() {
	w.position = w.PeekAdvance(1)
}

type Pachinko struct {
	wheels []*Wheel
}

func NewPachinko() *Pachinko {
	return &Pachinko{
		wheels: make([]*Wheel, 0),
	}
}

func (p *Pachinko) AddWheel(position, divisions int) {
	p.wheels = append(p.wheels, &Wheel{position, divisions})
}

func (p *Pachinko) LinedUp() bool {
	for i, wheel := range p.wheels {
		if wheel.PeekAdvance(i+1) != 0 {
			return false
		}
	}
	return true
}

func (p *Pachinko) Advance() {
	for _, wheel := range p.wheels {
		wheel.Advance()
	}
}

func (p *Pachinko) Run() int {
	t := 0
	for ; !p.LinedUp(); t++ {
		p.Advance()
	}

	return t
}
