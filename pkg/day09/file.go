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

func (f *File) BruteForceDecompress () {
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

func (f File) DecompressedLineSize (line string, recursive bool) int {
	size := 0
	idx  := 0
	for idx < len(line) {
		substr := line[idx:]
		idx1   := strings.Index(substr, "(")
		if idx1 > -1 {
			idx2      := strings.Index(substr, ")")
			marker    := substr[idx1+1:idx2]
			parts     := strings.Split(marker, "x")
			length, _ := strconv.Atoi(parts[0])
			times, _  := strconv.Atoi(parts[1])
			idx       += idx2 + length + 1
			if recursive {
				start  := idx2 + 1
				finish := start + length
				inner  := substr[start:finish]
				length  = f.DecompressedLineSize(inner, recursive)
			}
			size += idx1 + (length * times)
		} else {
			size += len(substr)
			idx  += len(substr)
		}
	}
	return size
}

func (f File) DecompressedSize (recursive bool) int {
	sum := 0
	for _, line := range f.lines {
		sum += f.DecompressedLineSize(line, recursive)
	}
	return sum
}

func (f File) Size () int {
	sum := 0
	for _, line := range f.lines {
		sum += len(line)
	}
	return sum
}