package main

import (
	"fmt"
	"strings"
	"testing"
)

type challengeTest struct {
	name      string
	input     string
	wantPart1 string
	wantPart2 string
}

func testChallenge(challenge *challenge, test challengeTest) func(t *testing.T) {
	return func(t *testing.T) {
		r := strings.NewReader(test.input)
		w := &strings.Builder{}
		var err error
		var want string

		if test.wantPart1 != "" && test.wantPart2 != "" {
			want = fmt.Sprintf("%s Part 1: %s\n%s Part 2: %s\n", challenge.name, test.wantPart1, challenge.name, test.wantPart2)
			err = runChallenge(w, r, challenge, 3)
		} else if test.wantPart1 != "" {
			want = fmt.Sprintf("%s Part 1: %s\n", challenge.name, test.wantPart1)
			err = runChallenge(w, r, challenge, 1)
		} else if test.wantPart2 != "" {
			want = fmt.Sprintf("%s Part 2: %s\n", challenge.name, test.wantPart2)
			err = runChallenge(w, r, challenge, 2)
		}

		if err == nil {
			got := w.String()
			if want != got {
				t.Errorf("Wanted %q got %q", want, got)
			}
		} else {
			t.Errorf("Unexpected Error: %v", err)
		}
	}
}
