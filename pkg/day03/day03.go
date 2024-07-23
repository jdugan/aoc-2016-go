package day03

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc/2016/pkg/reader"
)

// ========== PUBLIC FNS ==================================

func Both() {
	fmt.Println(" ")
	fmt.Println("DAY", 3)
	fmt.Println("  Puzzle 1", "=>", Puzzle1())
	fmt.Println("  Puzzle 2", "=>", Puzzle2())
	fmt.Println(" ")
}

func Puzzle1() int {
	count:= 0
	for _, triangle := range horizontal_triangles() {
		if triangle.IsValid() {
			count += 1
		}
	}
	return count
}

func Puzzle2() int {
	count:= 0
	for _, triangle := range vertical_triangles() {
		if triangle.IsValid() {
			count += 1
		}
	}
	return count
}

// ========== PRIVATE FNS =================================

func sides() []int {
	sides := make([]int, 0)
	lines := reader.Lines("./data/day03/input.txt")
	for _, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			val, _ := strconv.Atoi(word)
			sides   = append(sides, val)
		}
	}
	return sides
}

func horizontal_triangles() []Triangle {
	triangles := make([]Triangle, 0)
	chunk     := make([]int, 0)
	for _, side := range sides() {
		chunk = append(chunk, side)
		if len(chunk) == 3 {
			slices.Sort(chunk)
			triangles = append(triangles, Triangle{ a: chunk[0], b: chunk[1], c: chunk[2] })
			chunk     = make([]int, 0)
		}
	}
	return triangles
}

func vertical_triangles() []Triangle {
	triangles := make([]Triangle, 0)
	chunk     := make([]int, 0)
	for _, side := range sides() {
		chunk = append(chunk, side)
		if len(chunk) == 9 {
			chunk0 := []int{ chunk[0], chunk[3], chunk[6] }
			chunk1 := []int{ chunk[1], chunk[4], chunk[7] }
			chunk2 := []int{ chunk[2], chunk[5], chunk[8] }
			slices.Sort(chunk0)
			slices.Sort(chunk1)
			slices.Sort(chunk2)
			triangles = append(triangles, Triangle{ a: chunk0[0], b: chunk0[1], c: chunk0[2] })
			triangles = append(triangles, Triangle{ a: chunk1[0], b: chunk1[1], c: chunk1[2] })
			triangles = append(triangles, Triangle{ a: chunk2[0], b: chunk2[1], c: chunk2[2] })
			chunk     = make([]int, 0)
		}
	}
	return triangles
}