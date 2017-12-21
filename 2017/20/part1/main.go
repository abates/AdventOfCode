package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/abates/AdventOfCode/coordinate"
)

var ORIGIN = coordinate.New(0, 0, 0)

type Particle struct {
	id           int
	position     *coordinate.Coordinate
	velocity     *coordinate.Coordinate
	acceleration *coordinate.Coordinate
}

func NewParticle(id int, position, velocity, acceleration *coordinate.Coordinate) *Particle {
	return &Particle{id, position, velocity, acceleration}
}

func (p *Particle) Move() {
	p.velocity = p.velocity.Add(p.acceleration)
	p.position = p.position.Add(p.velocity)
}

func (p *Particle) Distance() int {
	return coordinate.ManhattanDistance(ORIGIN, p.position)
}

func move(particles map[int]*Particle) *Particle {
	var closest *Particle
	distance := 0

	for _, particle := range particles {
		particle.Move()
		newDistance := particle.Distance()
		if closest == nil || newDistance < distance {
			distance = newDistance
			closest = particle
		}
	}
	return closest
}

func removeCollisions(particles map[int]*Particle) {
	for id1, p1 := range particles {
		collision := false
		for id2, p2 := range particles {
			if id1 == id2 {
				continue
			}

			if p1.position.Equal(p2.position) {
				delete(particles, id2)
				collision = true
			}
		}

		if collision {
			delete(particles, id1)
		}
	}
}

func main() {
	particles := make(map[int]*Particle, 0)
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)

	for i, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		var px, py, pz int
		var vx, vy, vz int
		var ax, ay, az int
		fmt.Sscanf(line, "p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>", &px, &py, &pz, &vx, &vy, &vz, &ax, &ay, &az)
		particles[i] = NewParticle(i, coordinate.New(px, py, pz), coordinate.New(vx, vy, vz), coordinate.New(ax, ay, az))
	}

	//for {
	//fmt.Printf("Closest is %d\n", move(particles).id)
	//}

	for {
		move(particles)
		removeCollisions(particles)
		fmt.Printf("%d particles remaining\n", len(particles))
	}
}
