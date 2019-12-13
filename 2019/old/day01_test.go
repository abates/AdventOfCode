package main

import "testing"

func TestD1Parse(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"12", 12},
		{"14", 14},
		{"1969", 1969},
		{"100756", 100756},
	}

	for _, test := range tests {
		t.Run("Parsing Mass"+test.input, func(t *testing.T) {
			d1p1 := &D1{}
			err := d1p1.parse(test.input)
			if err == nil {
				if len(d1p1.modules) != 1 {
					t.Errorf("Expected 1 module to be added, got %d", len(d1p1.modules))
				} else if d1p1.modules[0] != test.want {
					t.Errorf("Wanted module mass %d got %d", test.want, d1p1.modules[0])
				}
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestMFR(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Mass 12", 12, 2},
		{"Mass 14", 14, 2},
		{"Mass 1969", 1969, 654},
		{"Mass 100756", 100756, 33583},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MFR(test.input)
			if test.want != got {
				t.Errorf("Wanted fuel requirement %d got %d", test.want, got)
			}
		})
	}
}

func TestTotalMFR(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Mass 14", 14, 2},
		{"Mass 1969", 1969, 966},
		{"Mass 100756", 100756, 50346},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TotalMFR(test.input)
			if test.want != got {
				t.Errorf("Wanted fuel requirement %d got %d", test.want, got)
			}
		})
	}
}
