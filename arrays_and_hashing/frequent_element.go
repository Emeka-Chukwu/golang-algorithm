package arraysandhashing

import (
	"sort"
)

// 347. Top K Frequent Elements
// Medium
// 16K
// 571
// Companies
// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.

// Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

func (ArrayAlg) topKFrequent(nums []int, k int) []int {
	var frequency = make(map[int]int)
	for _, value := range nums {
		frequency[value] += 1
	}
	type FrequencyT struct {
		Key   int
		Value int
	}
	data := make([]FrequencyT, 0)
	result := make([]int, 0)
	for key, value := range frequency {
		data = append(data, FrequencyT{Key: key, Value: value})
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})
	for i := 0; i < k; i++ {
		result = append(result, data[i].Key)
	}
	return result

}

// // needs modification
func (ArrayAlg) topKFrequentWithoutSoriung(nums []int, k int) []int {
	var frequency = make(map[int]int)
	var frequencyInverse = make(map[int]int)
	result := []int{}
	for _, value := range nums {
		frequency[value] += 1
	}
	for key, value := range frequency {
		frequencyInverse[value] = key
	}

	// fmt.Println(frequencyInverse)

	for i := len(nums) - 1; i > 0; i-- {
		if value := frequencyInverse[i]; value != 0 {
			result = append(result, value)
		}
		if len(result) >= 2 {
			return result
		}
	}
	return result
}
