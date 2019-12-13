package main

import "testing"

func TestDay02Run(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"test 1", "1,0,0,0,99", "2,0,0,0,99"},
		{"test 2", "2,3,0,3,99", "2,3,0,6,99"},
		{"test 3", "2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"test 4", "1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d2 := &D2{}
			err := d2.parse([]string{test.input})
			if err == nil {
				err = d2.runProgram(d2.mem)
				if err == nil {
					got := d2.dump(d2.mem)
					if test.want != got {
						t.Errorf("Wanted %q got %q", test.want, got)
					}
				} else {
					t.Errorf("Unexpected error: %v", err)
				}
			} else {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
