package main

type TakeFn func(elf *Elf, ex *WhiteElephantExchange)

type Elf struct {
	id       int
	numGifts int
	next     *Elf
}

func TakeGiftsLeft(e *Elf, ex *WhiteElephantExchange) {
	e.numGifts += e.next.numGifts
	e.next = e.next.next
}

func TakeGiftsAcross(e *Elf, ex *WhiteElephantExchange) {
	e.numGifts += ex.middleElf.next.numGifts
	ex.middleElf.next = ex.middleElf.next.next
}

type WhiteElephantExchange struct {
	startElf  *Elf
	middleElf *Elf
	size      int
	exchange  TakeFn
}

func NewWhiteElephantExchange(numElves int, exchange TakeFn) *WhiteElephantExchange {
	ex := &WhiteElephantExchange{
		startElf: &Elf{1, 1, nil},
		size:     numElves,
		exchange: exchange,
	}

	elf := ex.startElf
	for i := 2; i <= numElves; i++ {
		elf.next = &Elf{i, 1, nil}
		elf = elf.next
		if i == numElves/2 {
			ex.middleElf = elf
		}
	}
	elf.next = ex.startElf
	return ex
}

func (ex *WhiteElephantExchange) Exchange() int {
	elf := ex.startElf
	ex.startElf = nil
	for ; elf != elf.next; elf = elf.next {
		ex.exchange(elf, ex)
		if ex.size%2 != 0 {
			ex.middleElf = ex.middleElf.next
		}
		ex.size--
	}
	return elf.id
}

func (ex *WhiteElephantExchange) Size() int {
	return ex.size
}
