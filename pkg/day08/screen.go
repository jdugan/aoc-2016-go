package day08

import "fmt"

// ========== DEFINITION ==================================

type Screen struct {
	width int
	height int
	points map[string]Point
}


// ========== RECEIVERS ===================================

func (s *Screen) Initialize () {
	for y := range s.height {
		for x := range s.width {
			key   := s.KeyFromCoords(x, y)
			value := "."
			s.points[key] = Point{ x, y, value }
		}
	}
}

func (s *Screen) Swipe (card Card) {
	for _, cmd := range card.commands {
		switch cmd.action {
		case "rect":
			for y := range cmd.factor2 {
				for x := range cmd.factor1 {
					key          := s.KeyFromCoords(x, y)
					point        := s.points[key]
					s.points[key] = Point{ x: point.x, y: point.y, value: "#" }
				}
			}
		case "rotate_col":
			x    := cmd.factor1
			vals := s.ColValues(x)
			for y0, value := range vals {
				y   := (y0 + cmd.factor2) % s.height
				key := s.KeyFromCoords(x, y)
				s.points[key] = Point{ x, y, value }
			}
		case "rotate_row":
			y    := cmd.factor1
			vals := s.RowValues(y)
			for x0, value := range vals {
				x   := (x0 + cmd.factor2) % s.width
				key := s.KeyFromCoords(x, y)
				s.points[key] = Point{ x, y, value }
			}
		}
	}
}

func (s Screen) Voltage () int {
	sum := 0
	for _, p := range s.points {
		sum += p.Voltage()
	}
	return sum
}


// ---------- UTILITIES -----------------------------------

func (s Screen) ColValues (x int) []string {
	values := make([]string, 0)
	for y := range s.height {
		key := s.KeyFromCoords(x, y)
		values = append(values, s.points[key].value)
	}
	return values
}

func (s Screen) RowValues (y int) []string {
	values := make([]string, 0)
	for x := range s.width {
		key := s.KeyFromCoords(x, y)
		values = append(values, s.points[key].value)
	}
	return values
}

func (s Screen) KeyFromCoords (x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (s Screen) Print () {
	fmt.Println("")
	for y := range s.height {
		row := ""
		for x := range s.width {
			key := s.KeyFromCoords(x, y)
			row  = fmt.Sprintf("%s%s", row, s.points[key].value)
		}
		fmt.Println(row)
	}
	fmt.Println("")
}