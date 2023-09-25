package arraysandhashing

// Example 1:

// Input: nums = [1,2,3,1]
// Output: true
// Example 2:

// Input: nums = [1,2,3,4]
// Output: false
// Example 3:

// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

// Constraints:

// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109

func (ArrayAlg) ContainsDuplicate(nums []int) bool {
	set := make(map[int]int)
	for _, num := range nums {
		if _, exist := set[num]; exist {
			return true
		}
		set[num] = 1
	}
	return false
}
