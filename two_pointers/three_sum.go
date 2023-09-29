package twopointers

import (
	"math"
	"sort"
)

// 15. 3Sum
// Medium
// 28.6K
// 2.6K
// Companies
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// Example 2:

// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.
// Example 3:

// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

// Constraints:

// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105
func (AlTwoPointers) threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)
	noRepetitionChecker := math.MinInt
	for index, value := range nums {
		if noRepetitionChecker == value {
			continue
		}
		l, r := index+1, len(nums)-1
		for l < r {
			currentSum := nums[l] + nums[r] + value

			if currentSum > 0 {
				r--
			} else if currentSum < 0 {
				l++
			} else {
				res = append(res, []int{value, nums[l], nums[r]})
				l += 1
				for nums[l] == nums[l-1] && l < r {
					l += 1
				}
			}
			noRepetitionChecker = value
		}
	}
	return res

}
