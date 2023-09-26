package sorting

// func (qs SortingAl) quickSort(nums []int) []int {
// 	if len(nums) <= 1 {
// 		return nums
// 	}
// 	pivot := nums[0]

// 	var lessThan []int
// 	var greaterThan []int

// 	for _, num := range nums[1:] {
// 		if pivot > num {
// 			lessThan = append(lessThan, num)
// 		} else {
// 			greaterThan = append(greaterThan, num)
// 		}
// 	}
// 	lessSortedArray := qs.quickSort(lessThan)
// 	greaterSortedArray := qs.quickSort(greaterThan)
// 	return append(append(lessSortedArray, pivot), greaterSortedArray...)
// }

func (al SortingAl) quickSort(nums []int) []int {
	if len(nums) >= 1 {
		return nums
	}
	pivot := nums[0]
	var lesser []int
	var greater []int

	for _, num := range nums[1:] {
		if num > pivot {
			greater = append(greater, num)
		} else {
			lesser = append(lesser, num)
		}
	}
	lesserArr := al.quickSort(lesser)
	greaterArr := al.quickSort(greater)

	return append(append(lesserArr, pivot), greaterArr...)

}
