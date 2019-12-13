package main

func init() {
	d9 := &D9{}
	challenges = append(challenges, &challenge{"Day 09", "input/day09.txt", nil, d9.parseFile, d9.part1, d9.part2})
}

type D9 struct {
	mem []*Int
}

func (d9 *D9) parseFile(lines []string) (err error) {
	d9.mem, err = ParseComputerMemory(lines)
	return err
}

func (d9 *D9) part1() (string, error) {
	return NewComputer(d9.mem).RunWithInput("1")
}

func (d9 *D9) part2() (string, error) {
	return NewComputer(d9.mem).RunWithInput("2")
}
