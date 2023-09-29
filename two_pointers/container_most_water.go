package twopointers

import (
	"math"
)

// 11. Container With Most Water
// Medium
// 26.8K
// 1.5K
// Companies
// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

// Example 1:

// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
// Example 2:

// Input: height = [1,1]
// Output: 1

// Constraints:

// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104

// /// Brute force
func (AlTwoPointers) maxArea(height []int) int {
	max := math.MinInt
	min := math.MaxInt
	for index, value := range height {
		for i := index; i < len(height); i++ {
			min = height[i]
			if min > value {
				min = value
			}
			if min*(i-index) > max {
				max = min * (i - index)
			}
		}
	}
	return max
}

func (AlTwoPointers) maxAreaPointers(height []int) int {
	max := math.MinInt
	l, r := 0, len(height)-1
	total := math.MinInt
	for l < r {

		if height[l] < height[r] {
			total = height[l] * (r - l)
			l++
		} else {
			total = height[r] * (r - l)
			r--
		}

		if total > max {
			max = total
		}
	}
	return max
}
