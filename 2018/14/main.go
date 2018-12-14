package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Recipe struct {
	value     int
	previous  *Recipe
	next      *Recipe
	available bool
}

func (recipe *Recipe) Append(newRecipe *Recipe) {
	recipe.next.previous = newRecipe
	newRecipe.next = recipe.next
	newRecipe.previous = recipe
	recipe.next = newRecipe
}

type Elf struct {
	id            int
	currentRecipe *Recipe
}

func (elf *Elf) Advance() {
	elf.currentRecipe.available = true
	num := elf.currentRecipe.value + 1
	for i := 0; i < num; i++ {
		elf.currentRecipe = elf.currentRecipe.next
	}
	if !elf.currentRecipe.available {
		elf.currentRecipe = elf.currentRecipe.next
	}
	elf.currentRecipe.available = false
}

type RecipeList struct {
	first  *Recipe
	last   *Recipe
	elf1   *Elf
	elf2   *Elf
	length int
}

func (rl *RecipeList) Push(r int) {
	recipe := &Recipe{value: r, available: true}
	if rl.first == nil {
		recipe.next = recipe
		recipe.previous = recipe
		rl.first = recipe
		rl.last = recipe
		rl.elf1 = &Elf{id: 1, currentRecipe: recipe}
	} else if rl.elf2 == nil {
		rl.elf2 = &Elf{id: 2, currentRecipe: recipe}
	}
	rl.last.Append(recipe)
	rl.last = recipe
	rl.length++
}

func (rl *RecipeList) AddRecipes(recipes int) {
	if recipes == 0 {
		rl.Push(0)
	} else {
		values := []int{}
		for ; recipes > 0; recipes /= 10 {
			values = append(values, recipes%10)
		}

		for i := len(values) - 1; i >= 0; i-- {
			rl.Push(values[i])
		}
	}
}

func (rl *RecipeList) Run(length int) {
	for rl.length < length {
		rl.Advance()
	}
}

func (rl *RecipeList) Advance() {
	sum := rl.elf1.currentRecipe.value + rl.elf2.currentRecipe.value
	rl.AddRecipes(sum)
	rl.elf1.Advance()
	rl.elf2.Advance()
}

func (rl *RecipeList) Score(offset, length int) string {
	scores := []rune{}
	recipe := rl.last
	for i := 0; i < offset; i++ {
		scores = append([]rune{
			[]rune(strconv.Itoa(recipe.value))[0],
		}, scores...)
		recipe = recipe.previous
		if recipe == rl.first {
			break
		}
	}
	if len(scores) >= length {
		return string(scores[0:length])
	}
	return string(scores)
}

func (rl *RecipeList) String() string {
	values := []string{}
	for recipe := rl.first; ; recipe = recipe.next {
		if rl.elf1.currentRecipe == recipe {
			values = append(values, fmt.Sprintf("(%d)", recipe.value))
		} else if rl.elf2.currentRecipe == recipe {
			values = append(values, fmt.Sprintf("[%d]", recipe.value))
		} else {
			values = append(values, strconv.Itoa(recipe.value))
		}

		if recipe == rl.last {
			break
		}
	}
	return strings.Join(values, " ")
}

func part1(input string) {
	iterations, err := strconv.Atoi(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read recipes input: %v\n", err)
		os.Exit(-1)
	}
	list := &RecipeList{}
	list.AddRecipes(37)
	list.Run(iterations + 10)
	fmt.Printf("Part 1: %s\n", list.Score(10, 10))
}

func part2(input string) {
	l := len([]rune(input))

	list := &RecipeList{}
	list.AddRecipes(37)
	for {
		list.Advance()
		if list.Score(l, l) == input {
			fmt.Printf("Part 2: %d\n", list.length-l)
			break
		} else if list.Score(l+1, l) == input {
			fmt.Printf("Part 2: %d\n", list.length-l-1)
			break
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input>\n", os.Args[0])
		os.Exit(-1)
	}
	input := os.Args[1]

	part1(input)
	part2(input)
}
