package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

type HashFunc func(salt string, index int) string

func HashRepeat(num int) HashFunc {
	return func(salt string, index int) string {
		hash := Hash(salt, index)
		for i := 0; i < num; i++ {
			hash = MD5(hash)
		}
		return hash
	}
}

func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Hash(salt string, index int) string {
	return MD5(fmt.Sprintf("%s%d", salt, index))
}

func FindSequence(str string, length int) string {
	for i := 0; i <= len(str)-length; i++ {
		if strings.Count(str[i:i+length], str[i:i+1]) == length {
			return str[i : i+1]
		}
	}
	return ""
}

func FindKey(salt string, start int, hashFunc HashFunc) (string, int) {
	for i := start; ; i++ {
		key := hashFunc(salt, i)
		s := FindSequence(key, 3)
		if s != "" {
			for j := i + 1; j <= i+1000; j++ {
				hash := hashFunc(salt, j)
				index := strings.Index(hash, strings.Repeat(s, 5))
				if index >= 0 {
					return key, i
				}
			}
		}
	}
	return "", 0
}

type Key struct {
	Index int
	Hash  string
}

func GetKeys(salt string, hashFunc HashFunc) []Key {
	cache := make([]string, 0)
	cacheLookup := func(salt string, index int) string {
		for len(cache) <= index {
			cache = append(cache, hashFunc(salt, len(cache)))
		}
		return cache[index]
	}

	keys := make([]Key, 0)
	key := ""
	for index := 0; len(keys) < 64; index++ {
		key, index = FindKey(salt, index, cacheLookup)
		keys = append(keys, Key{index, key})
	}
	return keys
}
