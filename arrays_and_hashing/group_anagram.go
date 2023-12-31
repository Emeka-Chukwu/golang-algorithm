package arraysandhashing

import (
	"sort"
)

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

// m x n log n
func (ArrayAlg) groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	var results [][]string
	for _, value := range strs {
		data := []rune(value)
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
		datastr := string(data)
		anagramMap[datastr] = append(anagramMap[datastr], value)
	}

	for _, value := range anagramMap {
		results = append(results, value)
	}
	return results

}

func (ArrayAlg) groupAnagramsWithoutsort(strs []string) [][]string {
	var hastTable = make(map[int][]string)
	var result [][]string
	for _, value := range strs {
		var total int
		for _, char := range value {
			total += int(char - 'a')
		}
		hastTable[total] = append(hastTable[total], value)
	}

	for _, value := range hastTable {
		result = append(result, value)
	}
	return result
}
