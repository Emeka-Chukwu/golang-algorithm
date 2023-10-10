package slidingwindow

// 567. Permutation in String
// Medium
// 10.7K
// 353
// Companies
// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.
// Example 1:
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false
// Constraints:
// 1 <= s1.length, s2.length <= 104
// s1 and s2 consist of lowercase English letters.

func (SlidingWindowAlg) checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count, s2Count := [26]int{}, [26]int{}
	for index, _ := range s1 {
		s1Count[s1[index]-'a']++
		s2Count[s2[index]-'a']++
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}
	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		index := s2[r] - 'a'
		s2Count[index]++
		if s2Count[index] == s1Count[index] {
			matches++
		} else if s1Count[index]+1 == s2Count[index] {
			matches--
		}

		index = s2[l] - 'a'
		s1Count[index]++
		if s2Count[index] == s1Count[index] {
			matches++
		} else if s1Count[index]-1 == s2Count[index] {
			matches--
		}
		l++
	}
	return matches == 26
}
