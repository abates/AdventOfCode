package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Set map[string]struct{}

func (set Set) Add(str string) { set[str] = struct{}{} }

func (set Set) Intersection(other Set) Set {
	intersection := make(Set)
	for k := range set {
		if _, found := other[k]; found {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func (set Set) Delete(str string) {
	delete(set, str)
}

func (set Set) Len() int { return len(set) }

func (set Set) ToA() []string {
	values := []string{}
	for k := range set {
		values = append(values, k)
	}
	sort.Strings(values)
	return values
}

func (set Set) String() string {
	return fmt.Sprintf("[%s]", strings.Join(set.ToA(), ", "))
}

type Registers [4]int

func (r *Registers) UnmarshalText(text string) error {
	_, err := fmt.Sscanf(text, "[%d, %d, %d, %d]", &((*r)[0]), &((*r)[1]), &((*r)[2]), &((*r)[3]))
	return err
}

func (r Registers) String() string {
	return fmt.Sprintf("[%d, %d, %d, %d]", r[0], r[1], r[2], r[3])
}

type Alg struct {
	registers Registers
}

func (alg *Alg) addr(a, b, c int) {
	alg.registers[c] = alg.registers[a] + alg.registers[b]
}

func (alg *Alg) addi(a, b, c int) {
	alg.registers[c] = alg.registers[a] + b
}

func (alg *Alg) mulr(a, b, c int) {
	alg.registers[c] = alg.registers[a] * alg.registers[b]
}

func (alg *Alg) muli(a, b, c int) {
	alg.registers[c] = alg.registers[a] * b
}

func (alg *Alg) banr(a, b, c int) {
	alg.registers[c] = alg.registers[a] & alg.registers[b]
}

func (alg *Alg) bani(a, b, c int) {
	alg.registers[c] = alg.registers[a] & b
}

func (alg *Alg) borr(a, b, c int) {
	alg.registers[c] = alg.registers[a] | alg.registers[b]
}

func (alg *Alg) bori(a, b, c int) {
	alg.registers[c] = alg.registers[a] | b
}

func (alg *Alg) setr(a, b, c int) {
	alg.registers[c] = alg.registers[a]
}

func (alg *Alg) seti(a, b, c int) {
	alg.registers[c] = a
}

func (alg *Alg) gtir(a, b, c int) {
	if a > alg.registers[b] {
		alg.registers[c] = 1
	} else {
		alg.registers[c] = 0
	}
}

func (alg *Alg) gtri(a, b, c int) {
	if alg.registers[a] > b {
		alg.registers[c] = 1
	} else {
		alg.registers[c] = 0
	}
}

func (alg *Alg) gtrr(a, b, c int) {
	if alg.registers[a] > alg.registers[b] {
		alg.registers[c] = 1
	} else {
		alg.registers[c] = 0
	}
}

func (alg *Alg) eqir(a, b, c int) {
	if a == alg.registers[b] {
		alg.registers[c] = 1
	} else {
		alg.registers[c] = 0
	}
}

func (alg *Alg) eqri(a, b, c int) {
	if alg.registers[a] == b {
		alg.registers[c] = 1
	} else {
		alg.registers[c] = 0
	}
}

func (alg *Alg) eqrr(a, b, c int) {
	if alg.registers[a] == alg.registers[b] {
		alg.registers[c] = 1
	} else {
		alg.registers[c] = 0
	}
}

func (alg *Alg) execute(op string, a, b, c int) {
	switch op {
	case "addr":
		alg.addr(a, b, c)
	case "addi":
		alg.addi(a, b, c)
	case "mulr":
		alg.mulr(a, b, c)
	case "muli":
		alg.muli(a, b, c)
	case "banr":
		alg.banr(a, b, c)
	case "bani":
		alg.bani(a, b, c)
	case "borr":
		alg.borr(a, b, c)
	case "bori":
		alg.bori(a, b, c)
	case "setr":
		alg.setr(a, b, c)
	case "seti":
		alg.seti(a, b, c)
	case "gtir":
		alg.gtir(a, b, c)
	case "gtri":
		alg.gtri(a, b, c)
	case "gtrr":
		alg.gtrr(a, b, c)
	case "eqir":
		alg.eqir(a, b, c)
	case "eqri":
		alg.eqri(a, b, c)
	case "eqrr":
		alg.eqrr(a, b, c)
	}
}

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(input), "\n")
	numGt3 := 0
	i := 0
	candidates := make(map[int]Set)

	for i = 0; i < len(lines); {
		if !strings.HasPrefix(lines[i], "Before:") {
			break
		}

		var before Registers
		err := before.UnmarshalText(lines[i][8:])
		if err != nil {
			panic(err.Error())
		}
		i++

		opcode := 0
		a := 0
		b := 0
		c := 0
		_, err = fmt.Sscanf(lines[i], "%d %d %d %d", &opcode, &a, &b, &c)
		if err != nil {
			panic(err.Error())
		}
		i++

		var after Registers
		err = after.UnmarshalText(lines[i][8:])
		if err != nil {
			panic(err.Error())
		}
		i += 2

		matches := 0

		set := make(Set)
		for _, op := range []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtir", "gtri", "gtrr", "eqir", "eqri", "eqrr"} {
			alg := &Alg{
				registers: before,
			}
			alg.execute(op, a, b, c)
			if alg.registers == after {
				set.Add(op)
				matches++
			}
		}

		if s, found := candidates[opcode]; found {
			candidates[opcode] = set.Intersection(s)
		} else {
			candidates[opcode] = set
		}

		if matches >= 3 {
			numGt3++
		}
	}
	fmt.Printf("Part 1: %d\n", numGt3)

	opIndex := make(map[int]string)
	for len(candidates) > 0 {
		for opcode, set := range candidates {
			for _, oc := range opIndex {
				set.Delete(oc)
			}

			if set.Len() == 1 {
				opIndex[opcode] = set.ToA()[0]
				delete(candidates, opcode)
				break
			}
		}
	}

	alg := &Alg{}
	for ; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		opcode, a, b, c := 0, 0, 0, 0
		fmt.Sscanf(lines[i], "%d %d %d %d", &opcode, &a, &b, &c)
		alg.execute(opIndex[opcode], a, b, c)
	}
	fmt.Printf("Part 2: %d\n", alg.registers[0])
}
