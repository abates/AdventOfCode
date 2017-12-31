package main

type StateFn func(*Tape) StateFn

func Run(tape *Tape, state StateFn, iterations int) {
	for i := 0; i < iterations; i++ {
		state = state(tape)
	}
}

func StateA(tape *Tape) StateFn {
	if tape.Value() == 0 {
		tape.Write(1)
		tape.Advance()
		return StateB
	}

	tape.Write(0)
	tape.Rewind()
	return StateD
}

func StateB(tape *Tape) StateFn {
	if tape.Value() == 0 {
		tape.Write(1)
		tape.Advance()
		return StateC
	}
	tape.Write(0)
	tape.Advance()
	return StateF
}

func StateC(tape *Tape) StateFn {
	if tape.Value() == 0 {
		tape.Write(1)
		tape.Rewind()
		return StateC
	}
	tape.Write(1)
	tape.Rewind()
	return StateA
}

func StateD(tape *Tape) StateFn {
	if tape.Value() == 0 {
		tape.Write(0)
		tape.Rewind()
		return StateE
	}
	tape.Write(1)
	tape.Advance()
	return StateA
}

func StateE(tape *Tape) StateFn {
	if tape.Value() == 0 {
		tape.Write(1)
		tape.Rewind()
		return StateA
	}
	tape.Write(0)
	tape.Advance()
	return StateB
}

func StateF(tape *Tape) StateFn {
	if tape.Value() == 0 {
		tape.Write(0)
		tape.Advance()
		return StateC
	}
	tape.Write(0)
	tape.Advance()
	return StateE
}
