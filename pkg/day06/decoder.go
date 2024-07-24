package day06

import (
	"fmt"
	"slices"
)

// ========== DEFINITION ==================================

type Decoder struct {
	data []string
}


// ========== RECEIVERS ===================================

func (d Decoder) DecryptByFrequency () string {
	count_hash := d.GeneratePositionalHash()
	max_hash   := d.GeneratePositionalHashByFrequency(count_hash)
	secret     := d.GenerateSecretWord((max_hash))
	return secret
}

func (d Decoder) DecryptByScarcity () string {
	count_hash := d.GeneratePositionalHash()
	max_hash   := d.GeneratePositionalHashByScarcity(count_hash)
	secret     := d.GenerateSecretWord((max_hash))
	return secret
}


// ---------- UTILITIES -----------------------------------

func (d Decoder) FindLeastFrequentLetter (hash map[string]int) string {
	result := ""
	min    := 1000000
	for k, v := range hash {
		if v < min {
			min    = v
			result = k
		}
	}
	return result
}

func (d Decoder) FindMostFrequentLetter (hash map[string]int) string {
	result := ""
	max    := 0
	for k, v := range hash {
		if v > max {
			max    = v
			result = k
		}
	}
	return result
}

func (d Decoder) GenerateDefaultPositionalHash (size int) map[int]map[string]int {
	hash := make(map[int]map[string]int)
	for i := range size {
		hash[i] = make(map[string]int)
	}
	return hash
}

func (d Decoder) GeneratePositionalHash () map[int]map[string]int {
	words := data()
	hash  := d.GenerateDefaultPositionalHash(len(words[0]))
	for _, word := range words {
		for i, r := range word {
			key   := string(r)
			_, ok := hash[i][key]
			if ok {
				hash[i][key] += 1
			} else {
				hash[i][key]  = 1
			}
		}
	}
	return hash
}

func (d Decoder) GeneratePositionalHashByFrequency (hash map[int]map[string]int) map[int]string {
	max_hash := make(map[int]string)
	for k, v := range hash {
		max_hash[k] = d.FindMostFrequentLetter(v)
	}
	return max_hash
}

func (d Decoder) GeneratePositionalHashByScarcity (hash map[int]map[string]int) map[int]string {
	max_hash := make(map[int]string)
	for k, v := range hash {
		max_hash[k] = d.FindLeastFrequentLetter(v)
	}
	return max_hash
}

func (d Decoder) GenerateSecretWord (max_hash map[int]string) string {
	keys   := d.GenerateSortedKeys(max_hash)
	result := ""
	for _, key := range keys {
		result = fmt.Sprintf("%s%s", result, max_hash[key])
	}
	return result
}

func (d Decoder) GenerateSortedKeys (hash map[int]string) []int {
	keys := make([]int, 0)
	for key, _ := range hash {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}