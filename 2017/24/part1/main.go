package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Component struct {
	port1      int
	port2      int
	unusedPort int
}

func (component *Component) Connectable(candidates map[*Component]bool) map[*Component]bool {
	result := make(map[*Component]bool)
	for candidate, _ := range candidates {
		if component.unusedPort == candidate.port1 || component.unusedPort == candidate.port2 {
			result[candidate] = true
		}
	}
	return result
}

func (component *Component) Connect(other *Component) {
	if component.unusedPort == other.port1 {
		other.unusedPort = other.port2
	}

	if component.unusedPort == other.port2 {
		other.unusedPort = other.port1
	}
}

func (component *Component) Strength() int {
	return component.port1 + component.port2
}

func (component *Component) String() string {
	return fmt.Sprintf("%d/%d", component.port1, component.port2)
}

type Bridge struct {
	components []*Component
}

func (bridge *Bridge) Strength() int {
	strength := 0
	for _, component := range bridge.components {
		strength += component.Strength()
	}
	return strength
}

func (bridge *Bridge) Len() int {
	return len(bridge.components)
}

func (bridge *Bridge) String() string {
	str := make([]string, len(bridge.components))
	for i, component := range bridge.components {
		str[i] = component.String()
	}
	return strings.Join(str, "--")
}

func connect(bridge []*Component, components map[*Component]bool) []*Bridge {
	bridges := make([]*Bridge, 0)

	end := bridge[len(bridge)-1]
	candidates := end.Connectable(components)
	if len(candidates) == 0 {
		newBridge := &Bridge{
			components: make([]*Component, len(bridge)),
		}
		copy(newBridge.components, bridge)
		return []*Bridge{newBridge}
	} else {
		for candidate, _ := range candidates {
			delete(components, candidate)
			end.Connect(candidate)
			bridge = append(bridge, candidate)
			bridges = append(bridges, connect(bridge, components)...)
			bridge = bridge[0 : len(bridge)-1]
			components[candidate] = true
		}
	}
	return bridges
}

func run(input string) {
	components := make(map[*Component]bool)
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		fields := strings.Split(strings.TrimSpace(line), "/")
		port1, _ := strconv.Atoi(fields[0])
		port2, _ := strconv.Atoi(fields[1])
		components[&Component{port1: port1, port2: port2}] = true
	}

	var maxLength *Bridge
	var maxStrength *Bridge
	for _, bridge := range connect([]*Component{&Component{port1: 0, port2: 0}}, components) {
		if maxStrength == nil || maxStrength.Strength() < bridge.Strength() {
			maxStrength = bridge
		}

		if maxLength == nil {
			maxLength = bridge
		} else if maxLength.Len() <= bridge.Len() {
			if maxLength.Len() == bridge.Len() {
				if maxLength.Strength() < bridge.Strength() {
					maxLength = bridge
				}
			} else {
				maxLength = bridge
			}
		}
	}
	fmt.Printf("Part 1: %s %d\n", maxStrength.String(), maxStrength.Strength())
	fmt.Printf("Part 2: %s %d\n", maxLength.String(), maxLength.Strength())
}

func main() {
	//run("0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10")
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	run(string(b))
}
