package day02

// ========== DEFINITION ==================================

type Sprite struct {
	keypad map[string]map[string]string
	position string
}


// ========== RECEIVERS ===================================

func (s *Sprite) Move (cmd string) string {
	for _, direction := range cmd {
		kp_entry := s.keypad[s.position]
		pos, found := kp_entry[string(direction)]
		if found {
			s.position = pos
		}
	}
	return s.position
}