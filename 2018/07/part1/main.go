package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
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

func (i *Instructions) Evaluate() string {
	steps := []string{}
	for k := range i.index {
		steps = append(steps, k)
	}

	builder := &strings.Builder{}
	previous := make(map[string]struct{})
	for len(steps) > 0 {
		sort.Strings(steps)
		for j, step := range steps {
			if i.index[step].canEvaluate(previous) {
				previous[step] = struct{}{}
				builder.WriteString(step)
				steps = append(steps[0:j], steps[j+1:]...)
				break
			}
		}
	}
	return builder.String()
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
	fmt.Printf("Instructions: %v\n", instructions.Evaluate())
}
