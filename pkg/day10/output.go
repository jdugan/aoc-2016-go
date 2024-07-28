package day10

import "slices"

// ========== DEFINITION ==================================

type Output struct {
	id int
	chips []int
}

// ========== RECEIVERS ===================================

func (o *Output) AddChip (chip int) {
	o.chips = append(o.chips, chip)
	slices.Sort(o.chips)
}