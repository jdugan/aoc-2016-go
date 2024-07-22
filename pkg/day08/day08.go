package day08

import (
	"fmt"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 8)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	return -1
}

func Puzzle2() int {
	return -2
}

// ========== PRIVATE FNS =================================

func data() []string {
	lines := reader.Lines("./data/day08/input.txt")

	return lines
}
