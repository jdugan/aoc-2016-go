package day06

import (
	"fmt"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 6)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() string {
	decoder := Decoder{ data: data() }
	return decoder.DecryptByFrequency()
}

func Puzzle2() string {
	decoder := Decoder{ data: data() }
	return decoder.DecryptByScarcity()
}

// ========== PRIVATE FNS =================================

func data() []string {
	return reader.Lines("./data/day06/input.txt")
}
