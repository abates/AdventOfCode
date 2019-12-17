package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/abates/AdventOfCode/coordinate"
)

func TestD10Parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  asteroids
	}{
		{"test 1", ".#..#", asteroids{coordinate.New(1, 0), coordinate.New(4, 0)}},
		{"test 2", ".#..#\n.....\n#####\n", asteroids{coordinate.New(1, 0), coordinate.New(4, 0), coordinate.New(0, 2), coordinate.New(1, 2), coordinate.New(2, 2), coordinate.New(3, 2), coordinate.New(4, 2)}},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d10 := &D10{}
			challenge := &challenge{"Test Day 10", "", d10}
			err := parseFile(test.input, challenge)
			if err == nil {
				got := d10.asteroids
				if !reflect.DeepEqual(test.want, got) {
					t.Errorf("Wanted %v got %v", test.want, got)
				}
			} else {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestD10Sort(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantCenter  coordinate.Coordinate
		wantVisible asteroids
	}{
		{"test 1", ".#....#####...#..\n##...##.#####..##\n##...#...#.#####.\n..#.....#...###..\n..#.#.....#....##\n",
			coordinate.New(8, 3),
			asteroids{
				coordinate.New(8, 1),
				coordinate.New(9, 0),
				coordinate.New(9, 1),
				coordinate.New(10, 0),
				coordinate.New(9, 2),
				coordinate.New(11, 1),
				coordinate.New(12, 1),
				coordinate.New(11, 2),
				coordinate.New(15, 1),

				coordinate.New(12, 2),
				coordinate.New(13, 2),
				coordinate.New(14, 2),
				coordinate.New(15, 2),
				coordinate.New(12, 3),
				coordinate.New(16, 4),
				coordinate.New(15, 4),
				coordinate.New(10, 4),
				coordinate.New(4, 4),

				coordinate.New(2, 4),
				coordinate.New(2, 3),
				coordinate.New(0, 2),
				coordinate.New(1, 2),
				coordinate.New(0, 1),
				coordinate.New(1, 1),
				coordinate.New(5, 2),
				coordinate.New(1, 0),
				coordinate.New(5, 1),

				coordinate.New(6, 1),
				coordinate.New(6, 0),
				coordinate.New(7, 0),
			},
		},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d10 := &D10{}
			for _, line := range strings.Split(test.input, "\n") {
				if line == "" {
					continue
				}
				err := d10.parse(line)
				if err != nil {
					t.Errorf("Unexpected Error: %v", err)
				}
			}
			gotCenter, gotVisible := d10.asteroids.best()
			if !test.wantCenter.Equal(gotCenter) {
				t.Errorf("Wanted center %v got %v", test.wantCenter, gotCenter)
			}

			if !reflect.DeepEqual(test.wantVisible, gotVisible) {
				t.Errorf("Wanted visible %v got %v", test.wantVisible, gotVisible)
			}
		})
	}
}

func TestD10Parts(t *testing.T) {
	tests := []challengeTest{
		{"test 1", ".#..#\n.....\n#####\n....#\n...##\n", "Best Location: (3, 4): 8 asteroids visible", ""},
		{"test 2", ".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##\n", "", "Value for 200th asteroid destroyed: 802"},
	}

	for _, test := range tests {
		d10 := &D10{}
		challenge := &challenge{"Test Day 10", "", d10}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
