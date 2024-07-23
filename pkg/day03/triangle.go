package day03

// ========== DEFINITION ==================================

type Triangle struct {
	a int
	b int
	c int
}


// ========== RECEIVERS ===================================

func (t *Triangle) IsValid () bool {
	return t.a + t.b > t.c
}