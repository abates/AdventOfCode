package main

type Slot struct {
	Value    int
	previous *Slot
	next     *Slot
}

func (s *Slot) Next() *Slot {
	if s.next == nil {
		s.next = &Slot{
			previous: s,
		}
	}
	return s.next
}

func (s *Slot) Previous() *Slot {
	if s.previous == nil {
		s.previous = &Slot{
			next: s,
		}
	}
	return s.previous
}

type Tape struct {
	current *Slot
}

func NewTape() *Tape {
	return &Tape{
		current: &Slot{},
	}
}

func (t *Tape) Advance() {
	t.current = t.current.Next()
}

func (t *Tape) Rewind() {
	t.current = t.current.Previous()
}

func (t *Tape) Value() int {
	return t.current.Value
}

func (t *Tape) Write(value int) {
	t.current.Value = value
}

func (t *Tape) Checksum() int {
	checksum := 0
	head := t.current
	for ; head.previous != nil; head = head.previous {
	}

	for current := head; current != nil; current = current.next {
		checksum += current.Value
	}
	return checksum
}
