package twopointers

import (
	"strings"
	"unicode"
)

// 125. Valid Palindrome
// Easy
// 8.2K
// 7.9K
// Companies
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// Constraints:

// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters.

func (pointer AlTwoPointers) isPalindrome(s string) bool {

	cleanStr := ""
	for _, value := range s {
		if unicode.IsLetter(value) || unicode.IsDigit(value) {
			cleanStr += string(unicode.ToLower(value))
		}
	}
	return cleanStr == pointer.reverseStr(cleanStr)

}

func (AlTwoPointers) reverseStr(s string) string {
	reversedStr := ""
	for _, value := range s {
		reversedStr = string(value) + reversedStr
	}

	return reversedStr
}

func (alp AlTwoPointers) isPalindrome2(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !alp.isAlphaNumeric(rune(s[left])) {
			left++
		}
		for left < right && !alp.isAlphaNumeric(rune(s[right])) {
			right--
		}
		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func (AlTwoPointers) isAlphaNumeric(char rune) bool {
	return int(char) >= int('A') && int(char) <= int('Z') ||
		int(char) >= int('a') && int(char) <= int('z') ||
		int(char) >= int('0') && int(char) <= int('9')
}
