package day09

import (
	"fmt"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 9)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	file := file()
	file.Decompress()
	return file.Size()
}

func Puzzle2() int {
	return -2
}

// ========== PRIVATE FNS =================================

func file() File {
	lines := reader.Lines("./data/day09/input.txt")
	return File{ lines }
}
