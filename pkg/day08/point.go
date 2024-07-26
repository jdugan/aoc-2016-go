package day08

// ========== DEFINITION ==================================

type Point struct {
	x int
	y int
	value string
}

// ========== RECEIVERS ===================================

func (p Point) Voltage () int {
	if p.value == "#" {
		return 1
	} else {
		return 0
	}
}