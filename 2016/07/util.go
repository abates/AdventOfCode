package main

import (
	"strings"
)

type Sequence string

func (s Sequence) HasAbba() bool {
	letters := strings.Split(string(s), "")
	for i, l := range letters {
		if i > 2 {
			if letters[i-3] == l && letters[i-2] == letters[i-1] && l != letters[i-1] {
				return true
			}
		}
	}
	return false
}

func (s Sequence) Abas() []string {
	abas := make([]string, 0)
	letters := strings.Split(string(s), "")
	for i := 2; i < len(letters); i++ {
		if letters[i-2] == letters[i] && letters[i] != letters[i-1] {
			abas = append(abas, letters[i-2]+letters[i-1]+letters[i])
		}
	}
	return abas
}

type IP struct {
	supernets []Sequence
	hypernets []Sequence
}

func ParseIP(s string) IP {
	ip := IP{}
	letters := strings.Split(s, "")
	sequence := Sequence("")
	for _, s := range letters {
		if s == "[" {
			ip.supernets = append(ip.supernets, sequence)
			sequence = Sequence("")
		} else if s == "]" {
			ip.hypernets = append(ip.hypernets, sequence)
			sequence = Sequence("")
		} else {
			sequence += Sequence(s)
		}
	}

	if string(sequence) != "" {
		ip.supernets = append(ip.supernets, sequence)
	}
	return ip
}

func (ip IP) SupportsTLS() bool {
	hypernetHasAbba := false
	for _, hypernet := range ip.hypernets {
		if hypernet.HasAbba() {
			hypernetHasAbba = true
			break
		}
	}

	supernetHasAbba := false
	for _, supernet := range ip.supernets {
		if supernet.HasAbba() {
			supernetHasAbba = true
			break
		}
	}

	return supernetHasAbba && !hypernetHasAbba
}

func (ip IP) SupportsSSL() bool {
	abas := []string{}
	for _, supernet := range ip.supernets {
		abas = append(abas, supernet.Abas()...)
	}

	if len(abas) == 0 {
		return false
	}

	babs := []string{}
	for _, hypernet := range ip.hypernets {
		babs = append(babs, hypernet.Abas()...)
	}

	if len(babs) == 0 {
		return false
	}

	for _, aba := range abas {
		for _, bab := range babs {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}

	return false
}
