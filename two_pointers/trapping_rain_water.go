package twopointers

// 42. Trapping Rain Water
// Hard
// 29.3K
// 418
// Companies
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
// Example 2:
// Input: height = [4,2,0,3,2,5]
// Output: 9
// Constraints:
// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105

func (al AlTwoPointers) trap(height []int) int {
	l, r := 0, len(height)-1
	lMax, rMax := height[0], height[r]
	total := 0
	for l < r {
		isLLesser := al.getMinOfMaxandLisLess(lMax, rMax)
		if isLLesser {
			total += al.getDiff(lMax, height[l])
			l++
			lMax = al.getMax(lMax, height[l])
		} else {
			total += al.getDiff(rMax, height[r])
			r--
			rMax = al.getMax(rMax, height[r])
		}

	}
	return total

}

func (al AlTwoPointers) getMinOfMaxandLisLess(l, r int) bool {
	if r < l {
		return false
	}
	return true
}
func (al AlTwoPointers) getMax(l, r int) int {
	if r > l {
		return r
	}
	return l
}

func (AlTwoPointers) getDiff(max, value int) int {
	if max-value <= 0 {
		return 0
	}
	return max - value
}
