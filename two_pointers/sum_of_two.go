package twopointers

// 167. Two Sum II - Input Array Is Sorted
// Medium
// 10.8K
// 1.3K
// Companies
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.

// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

// The tests are generated such that there is exactly one solution. You may not use the same element twice.

// Your solution must use only constant extra space.

// Example 1:

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
// Example 2:

// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
// Example 3:

// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

// Constraints:

// 2 <= numbers.length <= 3 * 104
// -1000 <= numbers[i] <= 1000
// numbers is sorted in non-decreasing order.
// -1000 <= target <= 1000
// The tests are generated such that there is exactly one solution.

// this is good for unsorted arra
func (AlTwoPointers) twoSum(numbers []int, target int) []int {
	checkedSet := make(map[int]int)
	for index, value := range numbers {
		diff := target - value
		if val, ok := checkedSet[diff]; ok {
			return []int{val + 1, index + 1}
		}
		checkedSet[value] = index
	}
	return []int{0, 0}
}

// for sorted arrays this will be better

func (AlTwoPointers) twoSumSort(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		curSum := numbers[l] + numbers[r]
		if curSum > target {
			r--
		} else if curSum < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{l + 1, r + 1}
}
