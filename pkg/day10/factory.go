package day10

import (
	"slices"
)

// ========== DEFINITION ==================================

type Factory struct {
	bots map[int]Bot
	outputs map[int]Output
}

// ========== RECEIVERS ===================================

func (f Factory) Result () int {
	a := f.outputs[0].chips[0]
	b := f.outputs[1].chips[0]
	c := f.outputs[2].chips[0]
	return a * b * c
}

func (f *Factory) Process (checksums []int) int {
	checksum := -1
	ids      := f.GetDoubleFistedBotIds()
	for len(ids) > 0 {
		for _, id := range ids {
			bot  := f.bots[id]
			low  := bot.chips[0]
			high := bot.chips[1]
			if bot.low_receiver_type == "bot" {
				low_receiver := f.bots[bot.low_receiver_id]
				low_receiver.AddChip(low)
				f.bots[bot.low_receiver_id] = low_receiver
			} else {
				low_receiver := f.outputs[bot.low_receiver_id]
				low_receiver.AddChip(low)
				f.outputs[bot.low_receiver_id] = low_receiver
			}
			if bot.high_receiver_type == "bot" {
				high_receiver := f.bots[bot.high_receiver_id]
				high_receiver.AddChip(high)
				f.bots[bot.high_receiver_id] = high_receiver
			} else {
				high_receiver := f.outputs[bot.high_receiver_id]
				high_receiver.AddChip(high)
				f.outputs[bot.high_receiver_id] = high_receiver
			}
			if slices.Contains(bot.chips, checksums[0]) && slices.Contains(bot.chips, checksums[1]) {
				checksum = bot.id
			}
			bot.ClearChips()
			f.bots[id] = bot
		}
		ids = f.GetDoubleFistedBotIds()
	}
	return checksum
}

// ---------- UTILITIES -----------------------------------

func (f Factory) GetDoubleFistedBotIds () []int {
	ids := make([]int, 0)
	for id, bot := range f.bots {
		if bot.IsDoubleFisted() {
			ids = append(ids, id)
		}
	}
	return ids
}