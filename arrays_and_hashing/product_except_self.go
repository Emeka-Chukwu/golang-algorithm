package arraysandhashing

// 238. Product of Array Except Self
// Medium
// 20.1K
// 1.1K
// Companies
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
// Example 2:

// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:

// 2 <= nums.length <= 105
// -30 <= nums[i] <= 30
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

func (ArrayAlg) productExceptSelf(nums []int) []int {
	var total int
	var result []int
	var hasZero bool
	var zeroCounter int
	for _, value := range nums {
		if value != 0 {
			if total == 0 {
				total = value
			} else {
				total = value * total
			}

		} else {
			hasZero = true
			zeroCounter++
		}
	}
	for _, value := range nums {
		if value != 0 {
			val := int(total / value)
			if hasZero {
				val = 0
			}
			result = append(result, val)
		} else {
			if zeroCounter == 1 {
				result = append(result, total)
			} else {
				result = append(result, 0)
			}

		}
	}
	return result
}

func (ArrayAlg) productExceptSelfWithoutDivider(nums []int) []int {
	result := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix = prefix * nums[i]
	}
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * postfix
		postfix = postfix * nums[i]
	}
	return result
}
