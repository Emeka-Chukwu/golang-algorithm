package circular

// example 1
// move the elements to right by k position that
// k = 2
// nums = [2,3,4,5,6,7,8]
// result = [7,8,2,3,4,5,6]

// formula
// ((length + index) - k) % length

func (CircularAlgo) moveElementsToLeft(nums []int, k int) []int {
	length := len(nums)
	result := make([]int, length)
	for index, value := range nums {
		newK := ((index + length) - k) % length
		result[newK] = value
	}
	return result
}
