package slidingwindow

// 424. Longest Repeating Character Replacement
// Medium
// 9.6K
// 414
// Companies
// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.
// Example 1:
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.
// Example 2:
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achive this answer too.
// Constraints:
// 1 <= s.length <= 105
// s consists of only uppercase English letters.
// 0 <= k <= s.length

func (sl SlidingWindowAlg) characterReplacement(s string, k int) int {
	maxf, l, res := 0, 0, 0
	charSet := make(map[byte]int)
	for r, _ := range s {
		charSet[s[r]] += 1
		if charSet[s[r]] > maxf {
			maxf = charSet[s[r]]
		}
		if r-l+1-maxf > k {
			charSet[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
	}
	return res

}
