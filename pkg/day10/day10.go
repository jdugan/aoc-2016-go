package day10

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
	fmt.Println("DAY", 10)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	factory   := factory()
	checksums := checksums()
	return factory.Process(checksums)
}

func Puzzle2() int {
	factory   := factory()
	checksums := checksums()
	factory.Process(checksums)
	return factory.Result()
}

// ========== PRIVATE FNS =================================

func checksums () []int {
	chips   := make([]int, 0)
	line    := reader.Lines("./data/day10/input.txt")[0]
	parts   := strings.Split(line, ",")
	low, _  := strconv.Atoi(parts[0])
	high, _ := strconv.Atoi(parts[1])
	chips  = append(chips, low)
	chips  = append(chips, high)
	return chips
}

func factory() Factory {
	lines   := reader.Lines("./data/day10/input.txt")
	bots    := make(map[int]Bot)
	outputs := make(map[int]Output)

	re_bot  := regexp.MustCompile(`^bot (\d+) gives low to (\w+) (\d+) and high to (\w+) (\d+)$`)
	for _, line := range lines {
		elements := re_bot.FindStringSubmatch(line)
		if len(elements) > 0 {
			id, _      := strconv.Atoi(elements[1])
			low_type   := elements[2]
			low_id, _  := strconv.Atoi(elements[3])
			high_type  := elements[4]
			high_id, _ := strconv.Atoi(elements[5])
			bots[id] = Bot{
				id: 			 	id,
				chips: 			 	make([]int, 0),
				low_receiver_id: 	low_id,
				low_receiver_type: 	low_type,
				high_receiver_id: 	high_id,
				high_receiver_type: high_type,
			}
			if low_type == "output" {
				_, ok := outputs[low_id]
				if !ok {
					outputs[low_id] = Output{ id: low_id, chips: make([]int, 0) }
				}
			}
			if high_type == "output" {
				_, ok := outputs[high_id]
				if !ok {
					outputs[high_id] = Output{ id: high_id, chips: make([]int, 0) }
				}
			}
		}
	}

	re_chip := regexp.MustCompile(`^value (\d+) goes to bot (\d+)$`)
	for _, line := range lines {
		elements := re_chip.FindStringSubmatch(line)
		if len(elements) > 0 {
			chip, _ := strconv.Atoi(elements[1])
			id, _   := strconv.Atoi(elements[2])
			bot     := bots[id]
			bot.AddChip(chip)
			bots[id] = bot
		}
	}

	return Factory{ bots, outputs }
}
