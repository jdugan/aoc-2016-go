package day01

import (
	"fmt"
	"math"
)

// ========== DEFINITION ==================================

type Sprite struct {
	x int
	y int
	heading	string
}


// ========== RECEIVERS ===================================

func (s *Sprite) DistanceFromOrigin () int {
	x := float64(s.x)
	y := float64(s.y)
	return int(math.Abs(x) + math.Abs(y))
}

func (s *Sprite) Key () string {
	return fmt.Sprintf("%d,%d", s.x, s.y)
}

func (s *Sprite) Move (distance int) {
	switch s.heading {
	case "N":
		s.y -= distance
	case "S":
		s.y += distance
	case "E":
		s.x += distance
	case "W":
		s.x -= distance
	}
}

func (s *Sprite) Turn (direction string) {
	switch direction {
	case "L":
		switch s.heading {
		case "N":
			s.heading = "W"
		case "E":
			s.heading = "N"
		case "S":
			s.heading = "E"
		case "W":
			s.heading = "S"
		}
	case "R":
		switch s.heading {
		case "N":
			s.heading = "E"
		case "E":
			s.heading = "S"
		case "S":
			s.heading = "W"
		case "W":
			s.heading = "N"
		}
	}
}