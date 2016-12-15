package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		salt   string
		index  int
		result string
	}{
		{"abc", 18, "0034e0923cc38887a57bd7b1d4f953df"},
		{"abc", 39, "347dac6ee8eeea4652c7476d0f97bee5"},
	}

	for i, test := range tests {
		result := Hash(test.salt, test.index)
		if result != test.result {
			t.Errorf("Test %d failed. Expected %s got %s", i, test.result, result)
		}
	}
}

func TestSequence(t *testing.T) {
	tests := []struct {
		salt   string
		index  int
		result string
	}{
		{"abc", 18, "8"},
		{"abc", 39, "e"},
		{"abc", 7858, "0"},
	}

	for i, test := range tests {
		result := FindSequence(Hash(test.salt, test.index), 3)
		if result != test.result {
			t.Errorf("Test %d failed. Expected %s got %s", i, test.result, result)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		salt     string
		start    int
		result   int
		hashFunc HashFunc
	}{
		{"abc", 0, 39, Hash},
		{"abc", 40, 92, Hash},
		{"abc", 0, 10, HashRepeat(2016)},
	}

	for i, test := range tests {
		_, result := FindKey(test.salt, test.start, test.hashFunc)
		if result != test.result {
			t.Errorf("Test %d failed. Expected %d got %d", i, test.result, result)
		}
	}
}

func TestPart1GetKeys(t *testing.T) {
	indices := []int{
		39, 92, 110, 184, 291, 314, 343, 385, 459, 461, 489, 771, 781, 887,
		955, 1144, 5742, 5781, 5783, 6016, 6093, 6219, 7833, 7858, 7918, 7937,
		8042, 8045, 8183, 8189, 8205, 8232, 8375, 8407, 8431, 8503, 8517, 8626,
		8672, 8730, 8811, 9497, 9536, 13268, 13439, 13479, 13560, 13663, 15758,
		15883, 16187, 16342, 16479, 20087, 20371, 20582, 20635, 20669, 21908,
		21927, 21978, 22023, 22193, 22728,
	}

	keys := GetKeys("abc", Hash)
	for i, v := range indices {
		if keys[i].Index != v {
			t.Errorf("Expected index %d to be %d but got %d", i, v, keys[i].Index)
		}
	}
}

func TestPart2GetKeys(t *testing.T) {
	keys := GetKeys("abc", HashRepeat(2016))
	if keys[0].Index != 10 {
		t.Errorf("Expected %d Got %d", 10, keys[0].Index)
	}
}
