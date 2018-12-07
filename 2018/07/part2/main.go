package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

type Step struct {
	id           string
	dependencies map[string]struct{}
}

func (s *Step) canEvaluate(previous map[string]struct{}) bool {
	for k := range s.dependencies {
		if _, found := previous[k]; !found {
			return false
		}
	}
	return true
}

func (s *Step) AddDependency(dependency *Step) {
	s.dependencies[dependency.id] = struct{}{}
}

type Instructions struct {
	index map[string]*Step
}

func (i *Instructions) Lookup(id string) *Step {
	if i.index == nil {
		i.index = make(map[string]*Step)
	}

	step, found := i.index[id]
	if !found {
		step = &Step{id: id, dependencies: make(map[string]struct{})}
		i.index[id] = step
	}
	return step
}

type Worker struct {
	delay   int
	seconds int
	step    string
}

func (w *Worker) Increment() { w.seconds++ }

func (w *Worker) Done() bool {
	needed := int([]rune(w.step)[0]) - 64 + w.delay
	return w.seconds == needed-1
}

func (i *Instructions) Evaluate(numWorkers, delay int) int {
	steps := []string{}
	for step := range i.index {
		steps = append(steps, step)
	}

	previous := make(map[string]struct{})
	workers := []*Worker{}
	seconds := 0
	for len(steps) > 0 || len(workers) > 0 {
		sort.Strings(steps)
		for j := 0; j < len(steps); j++ {
			step := steps[j]
			if len(workers) < numWorkers {
				if i.index[step].canEvaluate(previous) {
					worker := &Worker{
						step:  step,
						delay: delay,
					}
					workers = append(workers, worker)
					steps = append(steps[0:j], steps[j+1:]...)
					j--
				}
			} else {
				break
			}
		}

		for j := 0; j < len(workers); j++ {
			worker := workers[j]
			if worker.Done() {
				previous[worker.step] = struct{}{}
				workers = append(workers[0:j], workers[j+1:]...)
				j--
			}
			worker.Increment()
		}
		seconds++
	}
	return seconds
}

func (i *Instructions) AddText(text []byte) error {
	id := ""
	dependent := ""
	_, err := fmt.Sscanf(string(text), "Step %s must be finished before step %s can begin.", &id, &dependent)
	if err == nil {
		step := i.Lookup(id)
		d := i.Lookup(dependent)
		d.AddDependency(step)
	}
	return err
}

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Error: " + err.Error())
	}

	instructions := &Instructions{}
	for _, line := range bytes.Split(input, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		err := instructions.AddText(line)
		if err != nil {
			panic("Error: " + err.Error())
		}
	}
	fmt.Printf("Seconds: %v\n", instructions.Evaluate(5, 60))
}
