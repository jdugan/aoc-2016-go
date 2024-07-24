package day05

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// ========== DEFINITION ==================================

type Decoder struct {
	seed string
}


// ========== RECEIVERS ===================================

func (d Decoder) FindComplexPassword () string {
	pwd_hash := d.DefaultPasswordHash()
	iter     := 0
	for !d.IsPasswordHashComplete(pwd_hash) {
		candidate := d.FindNextCandidate(iter)
		val, ok   := pwd_hash[candidate.value1]
		if ok && val == "x" {
			pwd_hash[candidate.value1] = candidate.value2
		}
		iter = candidate.iter
	}
	return d.GeneratePasswordFromHash(pwd_hash)
}

func (d Decoder) FindSimplePassword () string {
	password := ""
	iter     := 0
	for len(password) < 8 {
		candidate := d.FindNextCandidate(iter)
		password   = fmt.Sprintf("%s%s", password, candidate.value1)
		iter       = candidate.iter
	}
	return password
}


// ---------- UTILITIES -----------------------------------

func (d Decoder) DefaultPasswordHash () map[string]string {
	phash := make(map[string]string)
	for i := range 8 {
		key := strconv.Itoa(i)
		phash[key] = "x"
	}
	return phash
}

func (d Decoder) FindNextCandidate (iter int) Candidate {
	hash := "xxxxxxxxxx"
	for hash[:5] != "00000" {
		iter += 1
		data := []byte(fmt.Sprintf("%s%d", d.seed, iter))
		hash  = fmt.Sprintf("%x", md5.Sum(data))
	}
	return Candidate{ iter: iter, value1: string(hash[5]), value2: string(hash[6])}
}

func (d Decoder) GeneratePasswordFromHash (pwd_hash map[string]string) string {
	password := ""
	for i := range 8 {
		key := strconv.Itoa(i)
		password = fmt.Sprintf("%s%s", password, pwd_hash[key])
	}
	return password
}

func (d Decoder) IsPasswordHashComplete (pwd_hash map[string]string) bool {
	complete := true
	for _, v := range pwd_hash {
		if v == "x" {
			complete = false
			break
		}
	}
	return complete
}