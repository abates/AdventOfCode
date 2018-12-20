package graph

import (
	"bufio"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSPF(t *testing.T) {
	files, err := filepath.Glob("testdata/spf_test*.txt")
	if err != nil {
		panic(err.Error())
	}

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			panic(err.Error())
		}

		reader := bufio.NewReader(file)
		graph := NewBasicGraph()
		distances := make(map[string]map[string]int)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}

			fields := strings.Fields(string(line))
			if len(fields) == 0 {
				continue
			}

			switch fields[0] {
			case "e":
				node := fields[1]
				neighbor := fields[2]
				weight, _ := strconv.Atoi(fields[3])
				if _, found := distances[node]; !found {
					distances[node] = make(map[string]int)
					distances[node][node] = 0
				}

				graph.AddDirectedEdge(node, neighbor, weight)
			case "d":
				node := fields[1]
				neighbor := fields[2]
				distances[node][neighbor], _ = strconv.Atoi(fields[3])
			default:
				panic("Uknown type " + fields[0])
			}
		}

		spf := SPFAll(graph)
		got := make(map[string]map[string]int)
		for node1, spf := range spf {
			if _, found := got[node1.(*BasicNode).id]; !found {
				got[node1.(*BasicNode).id] = make(map[string]int)
			}
			for node2, distance := range spf {
				got[node1.(*BasicNode).id][node2.(*BasicNode).id] = distance
			}
		}

		if !reflect.DeepEqual(got, distances) {
			t.Errorf("tests[%s] expected %v got %v", filename, distances, got)
		}
	}
}
