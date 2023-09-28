package arraysandhashing

// 128. Longest Consecutive Sequence
// Medium
// 18.3K
// 829
// Companies
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

func (ArrayAlg) longestConsecutive(nums []int) int {
	checker := make(map[int]bool)
	for _, value := range nums {
		checker[value] = true
	}
	max := 0
	for _, value := range nums {
		if !checker[value-1] {
			length := 0
			for checker[value+length] {
				length++
			}
			if length > max {
				max = length
			}
		}
	}
	return max
}
