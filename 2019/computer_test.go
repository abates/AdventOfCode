package main

import (
	"testing"
)

func TestComputerRun(t *testing.T) {
	tests := []struct {
		name    string
		program string
		input   string
		wantMem string
		wantOut string
	}{
		{"test 1", "1,0,0,0,99", "", "2,0,0,0,99", ""},
		{"test 2", "2,3,0,3,99", "", "2,3,0,6,99", ""},
		{"test 3", "2,4,4,5,99,0", "", "2,4,4,5,99,9801", ""},
		{"test 4", "1,1,1,4,99,5,6,0,99", "", "30,1,1,4,2,5,6,0,99", ""},
		{"test 5", "1002,4,3,4,33", "", "1002,4,3,4,99", ""},
		{"test 6", "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", "", "", "109\n1\n204\n-1\n1001\n100\n1\n100\n1008\n100\n16\n101\n1006\n101\n0\n99"},
		{"test 7", "1102,34915192,34915192,7,4,7,99,0", "", "", "1219070632396864"},
		{"test 8", "104,1125899906842624,99", "", "", "1125899906842624"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			program, _ := ParseComputerMemory([]string{test.program})
			c := NewComputer(program)
			gotOut, err := c.RunWithInput(test.input)
			if err == nil {
				gotMem := c.Dump()
				if test.wantMem != "" && test.wantMem != gotMem {
					t.Errorf("Wanted %q got %q", test.wantMem, gotMem)
				}

				if test.wantOut != "" {
					if test.wantOut != gotOut {
						t.Errorf("Wanted output %q got %q", test.wantOut, gotOut)
					}
				}
			} else {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
