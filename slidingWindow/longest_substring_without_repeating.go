package slidingwindow

// 3. Longest Substring Without Repeating Characters
// Medium
// 37.4K
// 1.7K
// Companies
// Given a string s, find the length of the longest
// substring
//  without repeating characters.=
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

// Constraints:

// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
// // good solution but can't account for some edge cases
func (SlidingWindowAlg) lengthOfLongestSubstrinm(s string) int {
	max := 0
	counter := 0
	checker := make(map[rune]bool)
	for index, value := range s {
		if exist := checker[value]; exist {
			if counter > max {
				max = counter
			}
			counter = 0
			checker = make(map[rune]bool)
			if rune(s[index-1]) == value {
				checker[value] = true
				counter++
			}
		} else {
			counter++
			if index == len(s)-1 {
				if counter > max {
					max = counter
				}
			}
			checker[value] = true
		}

	}
	return max

}

func (SlidingWindowAlg) lengthOfLongestSubstring(s string) int {
	res := 0
	l := 0
	charSet := make(map[byte]bool)
	for r, _ := range s {
		for charSet[s[r]] {
			delete(charSet, s[l])
			l++
		}
		charSet[s[r]] = true
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res

}
