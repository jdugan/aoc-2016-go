package day01

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 1)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	sprite := Sprite{ x: 0, y: 0, heading: "N" }
	for _, cmd := range commands() {
		sprite.Turn(cmd.direction)
		sprite.Move(cmd.distance)
	}
	return sprite.DistanceFromOrigin()
}

func Puzzle2() int {
	commands := commands()
	sprite   := Sprite{ x: 0, y: 0, heading: "N" }
	visited  := map[string]bool{"0,0": true}
	found    := false
	if !found {
		cmd_loop:
		for _, cmd := range commands {
			sprite.Turn(cmd.direction)
			for i := 0; i < cmd.distance; i++ {
				sprite.Move(1)
				_, found = visited[sprite.Key()]
				if found {
					break cmd_loop
				} else {
					visited[sprite.Key()] = true
				}
			}
		}
	}
	return sprite.DistanceFromOrigin()
}

// ========== PRIVATE FNS =================================

func commands() []Command {
	lines := reader.Lines("./data/day01/input.txt")
	raws  := strings.Split(lines[0], ", ")
	cmds  := make([]Command, 0)
	for _, raw := range raws {
		direction   := string(raw[0])
		tmp         := string(raw[1:])
		distance, _ := strconv.Atoi(tmp)

		cmds = append(cmds, Command{direction: direction, distance: distance})
	}
	return cmds
}
