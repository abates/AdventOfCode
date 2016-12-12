package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
	"sync"
)

var wg = &sync.WaitGroup{}

type Droppable interface {
	Drop(value int)
}

type Instruction struct {
	LowDrop  Droppable
	HighDrop Droppable
}

type Bot struct {
	id           int
	done         chan bool
	chipCh       chan (int)
	instCh       chan (Instruction)
	chips        []int
	instructions []Instruction
	mutex        *sync.Mutex
}

func NewBot(id int) *Bot {
	return &Bot{
		id:     id,
		done:   make(chan bool, 1),
		chipCh: make(chan int, 2),
		instCh: make(chan Instruction, 1),
		mutex:  &sync.Mutex{},
	}
}

func (b *Bot) Drop(value int) {
	b.chipCh <- value
}

func (b *Bot) GiveInstruction(i Instruction) {
	b.instCh <- i
}

func (b *Bot) Idle() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return len(b.chips) == 0 && len(b.instructions) == 0
}

func (b *Bot) executeInstruction() {
	if len(b.instructions) > 0 && len(b.chips) >= 2 {
		if (b.chips[0] == 61 && b.chips[1] == 17) || (b.chips[0] == 17 && b.chips[1] == 61) {
			fmt.Printf("********* %s received 17 and 61\n", b)
		}

		instruction := b.instructions[0]
		if b.chips[0] < b.chips[1] {
			instruction.LowDrop.Drop(b.chips[0])
			instruction.HighDrop.Drop(b.chips[1])
		} else {
			instruction.LowDrop.Drop(b.chips[1])
			instruction.HighDrop.Drop(b.chips[0])
		}
		b.instructions = b.instructions[1:]
		b.chips = b.chips[2:]
	}
}

func (b *Bot) Run() {
	done := false
	for !done {
		select {
		case <-b.done:
			done = true
		case instruction := <-b.instCh:
			b.mutex.Lock()
			b.instructions = append(b.instructions, instruction)
			b.executeInstruction()
			b.mutex.Unlock()
		case v := <-b.chipCh:
			b.mutex.Lock()
			b.chips = append(b.chips, v)
			b.executeInstruction()
			b.mutex.Unlock()
		}
	}
	wg.Done()
}

func (b *Bot) Shutdown() {
	b.done <- true
}

func (b *Bot) String() string { return fmt.Sprintf("Bot %d", b.id) }

type Bin struct {
	id     int
	chipCh chan int
	chips  []int
	done   chan bool
}

func NewBin(id int) *Bin {
	return &Bin{
		id:     id,
		chipCh: make(chan int, 1),
		chips:  make([]int, 0),
		done:   make(chan bool, 1),
	}
}

func (b *Bin) Drop(value int) { b.chipCh <- value }

func (b *Bin) Run() {
	done := false
	for !done {
		select {
		case <-b.done:
			done = true
		case v := <-b.chipCh:
			b.chips = append(b.chips, v)
		}
	}
	wg.Done()
}

func (b *Bin) Shutdown()      { b.done <- true }
func (b *Bin) String() string { return fmt.Sprintf("Bin %d", b.id) }

type Factory struct {
	bots map[int]*Bot
	bins map[int]*Bin
}

func (f *Factory) GetBot(id int) *Bot {
	if _, found := f.bots[id]; !found {
		f.bots[id] = NewBot(id)
		go f.bots[id].Run()
		wg.Add(1)
	}
	return f.bots[id]
}

func (f *Factory) GetBin(id int) *Bin {
	if _, found := f.bins[id]; !found {
		f.bins[id] = NewBin(id)
		go f.bins[id].Run()
		wg.Add(1)
	}
	return f.bins[id]
}

func (f *Factory) GetDroppable(id int, devType string) Droppable {
	switch devType {
	case "bot":
		return f.GetBot(id)
	case "output":
		return f.GetBin(id)
	default:
		panic(fmt.Sprintf("Can't handle devtype %s", devType))
	}
}

func (f *Factory) Idle() bool {
	for _, bot := range f.bots {
		if !bot.Idle() {
			return false
		}
	}
	return true
}

func (f *Factory) Shutdown() {
	for _, bot := range f.bots {
		bot.Shutdown()
	}

	for _, bin := range f.bins {
		bin.Shutdown()
	}
}

func readInput() {
	f := &Factory{
		bots: make(map[int]*Bot, 0),
		bins: make(map[int]*Bin, 0),
	}

	for _, line := range util.ReadInput() {
		if strings.HasPrefix(line, "value") {
			var value, bot int
			fmt.Sscanf(line, "value %d goes to bot %d", &value, &bot)
			f.GetBot(bot).Drop(value)
		} else if strings.HasPrefix(line, "bot") {
			var bot, dev1Id, dev2Id int
			var dev1Type, dev2Type string
			fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &bot, &dev1Type, &dev1Id, &dev2Type, &dev2Id)

			f.GetBot(bot).GiveInstruction(Instruction{
				LowDrop:  f.GetDroppable(dev1Id, dev1Type),
				HighDrop: f.GetDroppable(dev2Id, dev2Type),
			})
		} else {
			panic("Can't handle " + line)
		}
	}

	for !f.Idle() {
	}

	f.Shutdown()
	wg.Wait()
	fmt.Printf("Bin 0: %v\n", f.GetBin(0).chips)
	fmt.Printf("Bin 1: %v\n", f.GetBin(1).chips)
	fmt.Printf("Bin 2: %v\n", f.GetBin(2).chips)
}
