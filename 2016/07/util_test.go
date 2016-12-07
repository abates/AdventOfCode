package main

import (
	"testing"
)

type testValue struct {
	result bool
	value  string
}

func TestSupportsTLS(t *testing.T) {
	tests := []testValue{
		{true, "abba"},
		{true, "abba[mnop]qrst"},
		{false, "abcd[bddb]xyyx"},
		{false, "aaaa[qwer]tyui"},
		{true, "ioxxoj[asdfgh]zxcvbn"},
	}

	for i, test := range tests {
		ip := ParseIP(test.value)
		result := ip.SupportsTLS()
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %v got %v", i, test.result, result)
		}
	}
}

func TestSupportSSL(t *testing.T) {
	tests := []testValue{
		{true, "aba[bab]xyz"},
		{false, "xyx[xyx]xyx"},
		{true, "aaa[kek]eke"},
		{true, "aaa[kek]ekek"},
		{true, "zazbz[bzb]cdb"},
	}

	for i, test := range tests {
		ip := ParseIP(test.value)
		result := ip.SupportsSSL()
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %v got %v", i, test.result, result)
		}
	}
}
