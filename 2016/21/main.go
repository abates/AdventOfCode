package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"github.com/cznic/mathutil"
	"sort"
	"strconv"
	"strings"
)

func scramble(passwd string, instructions []string) string {
	password := NewPassword(passwd)

	for _, line := range instructions {
		fields := strings.Fields(line)
		switch fields[0] {
		case "swap":
			if fields[1] == "position" {
				x, _ := strconv.Atoi(fields[2])
				y, _ := strconv.Atoi(fields[5])
				password.Swap(x, y)
			} else {
				password.SwapLetter(fields[2], fields[5])
			}
		case "rotate":
			if fields[1] == "left" {
				x, _ := strconv.Atoi(fields[2])
				password.RotateLeft(x)
			} else if fields[1] == "right" {
				x, _ := strconv.Atoi(fields[2])
				password.RotateRight(x)
			} else {
				password.RotatePosition(fields[6])
			}
		case "reverse":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[4])
			password.Reverse(x, y)
		case "move":
			x, _ := strconv.Atoi(fields[2])
			y, _ := strconv.Atoi(fields[5])
			password.Move(x, y)
		}
	}

	return password.String()
}

func main() {
	instructions := util.ReadInput()
	password := scramble("abcdefgh", instructions)
	fmt.Printf("Password: %s\n", password)

	letters := strings.Split("fbgdceah", "")
	str := sort.StringSlice(letters)
	mathutil.PermutationFirst(str)
	for mathutil.PermutationFirst(str); mathutil.PermutationNext(str); {
		//println(strings.Join(str, ""))
		password = scramble(strings.Join(str, ""), instructions)
		if password == "fbgdceah" {
			fmt.Printf("Password for %s is %s\n", password, strings.Join(str, ""))
			break
		}
	}
}
