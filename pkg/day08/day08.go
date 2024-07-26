package day08

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
	lines  := data()
	screen := screen(lines)
	card   := card(lines)
	screen.Swipe(card)
	return screen.Voltage()
}

func Puzzle2() string {
	lines  := data()
	screen := screen(lines)
	card   := card(lines)
	screen.Swipe(card)
	screen.Print()
	return "UPOJFLBCEZ"
}

// ========== PRIVATE FNS =================================

func data() []string {
	return reader.Lines("./data/day08/input.txt")
}

func card(lines []string) Card {
	re_rect  := regexp.MustCompile(`rect (\d+)x(\d+)`)
	re_row   := regexp.MustCompile(`rotate row y=(\d+) by (\d+)`)
	re_col   := regexp.MustCompile(`rotate column x=(\d+) by (\d+)`)
	commands := make([]Command, 0)
	for _, line := range lines[1:] {
		action  := ""
		factor1 := 0
		factor2 := 0
		elements := re_rect.FindStringSubmatch(line)
		if len(elements) > 0 {
			action     = "rect"
			factor1, _ = strconv.Atoi(elements[1])
			factor2, _ = strconv.Atoi(elements[2])
		} else {
			elements = re_row.FindStringSubmatch(line)
			if len(elements) > 0 {
				action     = "rotate_row"
				factor1, _ = strconv.Atoi(elements[1])
				factor2, _ = strconv.Atoi(elements[2])
			} else {
				elements = re_col.FindStringSubmatch(line)
				action     = "rotate_col"
				factor1, _ = strconv.Atoi(elements[1])
				factor2, _ = strconv.Atoi(elements[2])
			}
		}
		commands = append(commands, Command{ action, factor1, factor2 })
	}
	return Card{ commands }
}

func screen(lines []string) Screen {
	dims      := strings.Split(lines[0], "x")
	width, _  := strconv.Atoi(dims[0])
	height, _ := strconv.Atoi(dims[1])
	points    := make(map[string]Point)
	screen    := Screen{ width, height, points }
	screen.Initialize()
	return screen
}
