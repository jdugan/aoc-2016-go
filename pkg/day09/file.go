package day09

import (
	"fmt"
	"strconv"
	"strings"
)

// ========== DEFINITION ==================================

type File struct {
	lines []string
}


// ========== RECEIVERS ===================================

func (f *File) Decompress () {
	new_lines := make([]string, 0)
	for _, line := range f.lines {
		new_line   := ""
		collecting := false
		marker     := ""
		max_idx    := len(line)
		idx        := 0
		for idx < max_idx {
			s := string(line[idx])
			switch s {
			case "(":
				collecting = true
				idx += 1
			case ")":
				collecting = false
				idx       += 1
				parts     := strings.Split(marker, "x")
				length, _ := strconv.Atoi(parts[0])
				times, _  := strconv.Atoi(parts[1])
				substr    := line[idx:idx+length]
				for i := 0; i < times; i++ {
					new_line = fmt.Sprintf("%s%s", new_line, substr)
				}
				marker  = ""
				idx    += length
			default:
				if collecting {
					marker = fmt.Sprintf("%s%s", marker, s)
				} else {
					new_line = fmt.Sprintf("%s%s", new_line, s)
				}
				idx += 1
			}
		}
		new_lines = append(new_lines, new_line)
	}
	f.lines = new_lines
}

func (f File) Size () int {
	sum := 0
	for _, line := range f.lines {
		fmt.Println(line)
		sum += len(line)
	}
	return sum
}