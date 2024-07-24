package day05

import (
	"fmt"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 5)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() string {
	decoder := Decoder{ seed: seed() }
	return decoder.FindSimplePassword()
}

func Puzzle2() string {
	decoder := Decoder{ seed: seed() }
	return decoder.FindComplexPassword()
}

// ========== PRIVATE FNS =================================

func seed() string {
	lines := reader.Lines("./data/day05/input.txt")
	return lines[0]
}
