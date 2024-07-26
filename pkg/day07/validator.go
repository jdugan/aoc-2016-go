package day07

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// ========== DEFINITION ==================================

type Validator struct {
	addresses []string
}


// ========== RECEIVERS ===================================

func (v Validator) SslCount () int {
	count := 0
	for _, addr := range v.addresses {
		if v.SupportsSsl(addr) {
			count += 1
		}
	}
	return count
}

func (v Validator) TlsCount () int {
	count := 0
	for _, addr := range v.addresses {
		if v.SupportsTls(addr) {
			count += 1
		}
	}
	return count
}

// ---------- UTILITIES -----------------------------------

func (v Validator) ConvertAbasToBabs (abas []string) []string {
	babs := make([]string, 0)
	for _, aba := range abas {
		a   := string(aba[0])
		b   := string(aba[1])
		babs = append(babs, fmt.Sprintf("%s%s%s", b, a, b))
	}
	return babs
}

func (v Validator) GetAbaProtocolMatches (addrs []string) []string {
	matches := make([]string, 0)
	for _, addr := range addrs {
		parts := []string{ "-", "-", "-" }
		for _, r := range addr {
			parts = append(parts[1:], string(r))
			if parts[0] != parts[1] && parts[0] == parts[2] {
				matches = append(matches, strings.Join(parts, ""))
			}
		}
	}
	return matches
}

func (v Validator) GetHypernetAddressSequences (addr string) []string {
	re_start  := regexp.MustCompile(`^[a-z]+\[`)
	re_middle := regexp.MustCompile(`\][a-z]+\[`)
	re_end    := regexp.MustCompile(`\][a-z]+$`)

	modified := re_start.ReplaceAllString(addr, "")
	modified  = re_end.ReplaceAllString(modified, "")
	modified  = re_middle.ReplaceAllString(modified, "-")
	return strings.Split(modified, "-")
}

func (v Validator) GetSupernetAddressSequences (addr string) []string {
	re_start  := regexp.MustCompile(`^-`)
	re_middle := regexp.MustCompile(`\[[a-z]+\]`)
	re_end    := regexp.MustCompile(`-$`)

	modified := re_middle.ReplaceAllString(addr, "-")
	modified  = re_start.ReplaceAllString(modified, "")
	modified  = re_end.ReplaceAllString(modified, "")
	return strings.Split(modified, "-")
}

func (v Validator) HasAbbaProtocol (addrs []string) bool {
	found := false
	addr_loop:
	for _, addr := range addrs {
		parts := []string{ "-", "-", "-", "-" }
		for _, r := range addr {
			parts = append(parts[1:], string(r))
			if parts[0] != parts[1] && parts[0] == parts[3] && parts[1] == parts[2] {
				found = true
				break addr_loop
			}
		}

	}
	return found
}

func (v Validator) HasBabProtocol (addrs []string, babs [] string) bool {
	found := false
	addr_loop:
	for _, addr := range addrs {
		parts := []string{ "-", "-", "-" }
		for _, r := range addr {
			parts = append(parts[1:], string(r))
			if slices.Contains(babs, strings.Join(parts, "")) {
				found = true
				break addr_loop
			}
		}
	}
	return found
}

func (v Validator) SupportsSsl (addr string) bool {
	supernets := v.GetSupernetAddressSequences(addr)
	hypernets := v.GetHypernetAddressSequences(addr)
	abas      := v.GetAbaProtocolMatches(supernets)
	babs      := v.ConvertAbasToBabs(abas)
	secure    := v.HasBabProtocol(hypernets, babs)
	fmt.Println("-------------------------")
	fmt.Println(addr)
	fmt.Println(abas)
	fmt.Println(babs)
	fmt.Println(secure)
	return secure
}

func (v Validator) SupportsTls (addr string) bool {
	supernets := v.GetSupernetAddressSequences(addr)
	hypernets := v.GetHypernetAddressSequences(addr)
	secure    := v.HasAbbaProtocol(supernets) && !v.HasAbbaProtocol(hypernets)
	return secure
}