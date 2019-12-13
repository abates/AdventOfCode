package main

import "testing"

func TestD7TryCombo(t *testing.T) {
	tests := []struct {
		name     string
		feedback bool
		program  string
		input    []int
		want     int
	}{
		{"test 1", false, "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", []int{4, 3, 2, 1, 0}, 43210},
		{"test 2", true, "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", []int{9, 8, 7, 6, 5}, 139629729},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d7 := &D7{}
			d7.mem, _ = ParseComputerMemory([]string{test.program})
			got, err := d7.tryCombo(test.feedback, test.input)
			if err == nil {
				if test.want != got {
					t.Errorf("Wanted %d got %d", test.want, got)
				}
			} else {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestD7Parts(t *testing.T) {
	tests := []challengeTest{
		{"test 1", "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", "Max Thruster Signal: 43210", ""},
		{"test 2", "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", "Max Thruster Signal: 54321", ""},
		{"test 3", "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", "Max Thruster Signal: 65210", ""},
		{"test 4", "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", "", "Max Thruster Signal: 139629729"},
		{"test 5", "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10", "", "Max Thruster Signal: 18216"},
	}

	for _, test := range tests {
		d7 := &D7{}
		challenge := &challenge{"Test Day 07", "", nil, d7.parseFile, d7.part1, d7.part2}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
