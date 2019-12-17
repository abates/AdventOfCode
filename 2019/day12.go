package main

import (
	"fmt"
)

func init() {
	d12 := &D12{}
	challenges[12] = &challenge{"Day 12", "input/day12.txt", d12}
}

type Moon struct {
	start    [3]int
	position [3]int
	velocity [3]int
}

func NewMoon(x, y, z int) *Moon {
	return &Moon{[3]int{x, y, z}, [3]int{x, y, z}, [3]int{0, 0, 0}}
}

func (m *Moon) Equal(other *Moon) bool {
	return m.start == other.start && m.velocity == other.velocity
}

func (m *Moon) UpdatePosition() {
	m.position[0] += m.velocity[0]
	m.position[1] += m.velocity[1]
	m.position[2] += m.velocity[2]
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (m *Moon) Energy() int {
	potential := abs(m.position[0]) + abs(m.position[1]) + abs(m.position[2])
	kinetic := abs(m.velocity[0]) + abs(m.velocity[1]) + abs(m.velocity[2])
	return potential * kinetic
}

func UpdateVelocity(m1, m2 *Moon) {
	for j := 0; j < 3; j++ {
		if m1.position[j] < m2.position[j] {
			m1.velocity[j] += 1
			m2.velocity[j] -= 1
		} else if m1.position[j] > m2.position[j] {
			m1.velocity[j] -= 1
			m2.velocity[j] += 1
		}
	}
}

func (m *Moon) String() string {
	return fmt.Sprintf("pos=%v vel=%v", m.position, m.velocity)
}

type OrbitalSystem []*Moon

func (os OrbitalSystem) Reset() {
	for _, moon := range os {
		moon.position = moon.start
		moon.velocity = [3]int{0, 0, 0}
	}
}

func (os OrbitalSystem) Step() {
	for i, m1 := range os {
		for _, m2 := range os[i:] {
			UpdateVelocity(m1, m2)
		}
	}

	for _, m1 := range os {
		m1.UpdatePosition()
	}
}

type D12 struct {
	moons OrbitalSystem
}

func (d12 *D12) parse(line string) error {
	x, y, z := 0, 0, 0
	_, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
	if err == nil {
		d12.moons = append(d12.moons, NewMoon(x, y, z))
	}
	return err
}

func (d12 *D12) runSystem(steps int) string {
	d12.moons.Reset()
	for i := 0; i < steps; i++ {
		d12.moons.Step()
	}
	energy := 0
	for _, moon := range d12.moons {
		energy += moon.Energy()
	}
	return fmt.Sprintf("Total Energy: %d", energy)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func lcm(values ...int) int {
	if len(values) > 2 {
		return lcm(values[0], lcm(values[1:]...))
	} else {
		ab := values[0] * values[1]
		if ab < 0 {
			ab *= -1
		}
		return ab / gcd(values[0], values[1])
	}
}

func (d12 *D12) runSystem2() string {
	steps := make([]int, 3)
	d12.moons.Reset()
	done := make(map[int]bool)
	period := 0
	for len(done) < 3 {
		period++
		d12.moons.Step()
		for i := 0; i < 3; i++ {
			if _, found := done[i]; found {
				continue
			}

			atStart := true
			for _, moon := range d12.moons {
				if moon.position[i] == moon.start[i] {
					if moon.velocity[i] != 0 {
						atStart = false
						break
					}
				} else {
					atStart = false
					break
				}
			}

			if atStart {
				done[i] = true
				steps[i] = period
			}
		}
	}
	return fmt.Sprintf("Steps to Equilibrium: %d", lcm(steps...))
}

func (d12 *D12) part1() (string, error) {
	return d12.runSystem(1000), nil
}

func (d12 *D12) part2() (string, error) {
	return d12.runSystem2(), nil
}
