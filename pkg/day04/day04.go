package day04

import (
	"fmt"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 4)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

// 749 too low
func Puzzle1() int {
	sum := 0
	for _, room := range rooms() {
		if room.IsValid() {
			sum += room.sector
		}
	}
	return sum
}

func Puzzle2() int {
	secret_name := "northpole object storage"
	sector      := 0
	for sector == 0 {
		for _, room := range rooms() {
			if room.IsValid() && room.DecryptName() == secret_name {
				sector = room.sector
			}
		}
	}
	return sector
}

// ========== PRIVATE FNS =================================

func rooms() []Room {
	rooms := make([]Room, 0)
	lines := reader.Lines("./data/day04/input.txt")
	for _, line := range lines {
		room := Room{}.Parse(line)
		rooms = append(rooms, room)
	}
	return rooms
}
