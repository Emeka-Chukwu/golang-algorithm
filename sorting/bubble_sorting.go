package sorting

func (SortingAl) bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		swapped := false
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return nums
}
