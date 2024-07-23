package day02

import (
	"fmt"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 2)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() string {
	sprite := Sprite{ keypad: assumed_keypad(), position: "5" }
	code  := ""
	for _, cmd := range commands() {
		code = fmt.Sprintf("%s%s", code, sprite.Move(cmd))
	}
	return code
}

func Puzzle2() string {
	sprite := Sprite{ keypad: real_keypad(), position: "5" }
	code  := ""
	for _, cmd := range commands() {
		code = fmt.Sprintf("%s%s", code, sprite.Move(cmd))
	}
	return code
}


// ========== PRIVATE FNS =================================

func assumed_keypad() map[string]map[string]string {
	return map[string]map[string]string{
		"1": { "D": "4", "R": "2" },
		"2": { "D": "5", "L": "1", "R": "3" },
		"3": { "D": "6", "L": "2"},
		"4": { "D": "7", "R": "5", "U": "1" },
		"5": { "D": "8", "L": "4", "R": "6", "U": "2" },
		"6": { "D": "9", "L": "5", "U": "3" },
		"7": { "R": "8", "U": "4" },
		"8": { "L": "7", "R": "9", "U": "5" },
		"9": { "L": "8", "U": "6" },
	}
}

func real_keypad() map[string]map[string]string {
	return map[string]map[string]string{
		"1": { "D": "3" },
		"2": { "D": "6", "R": "3" },
		"3": { "D": "7", "L": "2", "R": "4", "U": "1" },
		"4": { "D": "8", "L": "3" },
		"5": { "R": "6" },
		"6": { "D": "A", "L": "5", "R": "7", "U": "2" },
		"7": { "D": "B", "L": "6", "R": "8", "U": "3" },
		"8": { "D": "C", "L": "7", "R": "9", "U": "4" },
		"9": { "L": "8" },
		"A": { "R": "B", "U": "6" },
		"B": { "D": "D", "L": "A", "R": "C", "U": "7" },
		"C": { "L": "B", "U": "8" },
		"D": { "U": "B" },
	}
}

// ---------- DATA HELPERS --------------------------------

func commands() []string {
	return reader.Lines("./data/day02/input.txt")
}
