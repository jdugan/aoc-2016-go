package day04

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// ========== DEFINITION ==================================

type Room struct {
	checksum string
	name 	 string
	sector 	 int
}


// ========== RECEIVERS ===================================

func (r Room) CalculateChecksum () string {
	hash     := r.TranslateNameToHash()
	inverted := r.InvertNameHash(hash)
	keys     := r.SortInvertedKeys(inverted)
	letters  := make([]string, 0)
	for _, k := range keys {
		letters = append(letters, inverted[k]...)
	}
	if len(letters) > 5 {
		letters = letters[0:5]
	}
	return strings.Join(letters, "")
}

func (r Room) IsValid () bool {
	return r.checksum == r.CalculateChecksum()
}

func (r Room) Parse (key string) Room {
	words     := strings.Split(key,"-")
	final     := words[len(words)-1]
	parts     := strings.Split(final, "[")
	name      := strings.Join(words[0:len(words)-1], "-")
	checksum  := parts[1][:len(parts[1])-1]
	sector, _ := strconv.Atoi(parts[0])

	return Room{ name: name, checksum: checksum, sector: sector }
}

func (r Room) DecryptName () string {
	lookup := []string{ "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" }
	dname  := ""
	for _, rn := range r.name {
		s := string(rn)
		if s == "-" {
			dname = fmt.Sprintf("%s ", dname)
		} else {
			idx := ((r.sector % 26) + slices.Index(lookup, s)) % 26
			dname = fmt.Sprintf("%s%s", dname, lookup[idx])
		}
	}
	return dname
}


// ---------- UTILITIES -----------------------------------

func (r Room) InvertNameHash (hash map[string]int) map[int][]string {
	inverted := make(map[int][]string)
	for k, v := range hash {
		s, found := inverted[v]
		if found {
			s = append(s, k)
			slices.Sort(s)
		} else {
			s = []string{ k }
		}
		inverted[v] = s
	}
	return inverted
}

func (r Room) TranslateNameToHash () map[string]int {
	hash := make(map[string]int)
	for _, nr := range r.name {
		key := string(nr)
		if key != "-" {
			_, found := hash[key]
			if found {
				hash[key] += 1
			} else {
				hash[key]  = 1
			}
		}
	}
	return hash
}

func (r Room) SortInvertedKeys (inverted map[int][]string) []int {
	keys     := make([]int, 0)
	for k, _ := range inverted {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	slices.Reverse(keys)
	return keys
}