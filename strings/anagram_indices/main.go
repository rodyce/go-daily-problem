package anagram_indices

import (
	"reflect"
)

type Counter map[rune]int

func makeCounter(s string) Counter {
	result := make(map[rune]int)
	for _, c := range s {
		if _, ok := result[c]; ok {
			result[c]++
		} else {
			result[c] = 1
		}
	}
	return result
}

func counterEquals(c1, c2 *Counter) bool {
	return reflect.DeepEqual(c1, c2)
}

func isAnagram(s1, s2 string) bool {
	c1 := makeCounter(s1)
	c2 := makeCounter(s2)
	return counterEquals(&c1, &c2)
}

func findAnagramIndicesNaive(word, s string) []int {
	result := []int{}
	word_len := len(word)
	s_len := len(s)
	if word_len > s_len {
		return result
	}
	for i := 0; i <= s_len-word_len; i++ {
		sub := s[i : i+word_len]
		if isAnagram(word, sub) {
			result = append(result, i)
		}
	}
	return result
}

func findAnagramIndices(word, s string) []int {
	// Sample closure :-)
	del_if_zero := func(m map[rune]int, key rune) {
		value, ok := m[key]
		if ok && value == 0 {
			delete(m, key)
		}
	}

	result := []int{}
	word_len := len(word)
	s_len := len(s)
	if word_len > s_len {
		return result
	}

	freq := makeCounter(word)

	for _, c := range s[:len(word)] {
		freq[c] -= 1
		del_if_zero(freq, c)
	}

	if len(freq) == 0 {
		result = append(result, 0)
	}

	for i := word_len; i < s_len; i++ {
		start_char, end_char := s[i-word_len], s[i]

		freq[rune(start_char)]++
		del_if_zero(freq, rune(start_char))

		freq[rune(end_char)]--
		del_if_zero(freq, rune(end_char))

		if len(freq) == 0 {
			beginning_index := i - len(word) + 1
			result = append(result, beginning_index)
		}
	}

	return result
}
