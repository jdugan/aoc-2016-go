package day10

import "slices"

// ========== DEFINITION ==================================

type Bot struct {
	id int
	chips []int
	high_receiver_id int
	high_receiver_type string
	low_receiver_id int
	low_receiver_type string
}


// ========== RECEIVERS ===================================

func (b *Bot) AddChip (chip int) {
	b.chips = append(b.chips, chip)
	slices.Sort(b.chips)
}

func (b *Bot) ClearChips () {
	b.chips = make([]int, 0)
}

func (b Bot) IsDoubleFisted () bool {
	return len(b.chips) == 2
}