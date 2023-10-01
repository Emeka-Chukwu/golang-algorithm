package twopointers

import "fmt"

type AlTwoPointers struct{}

func (al AlTwoPointers) Run() {
	fmt.Println()
	fmt.Println("Start isPalindrome ============")

	fmt.Println("gotten result :", al.isPalindrome2("A man, a plan, a canal: Panama"), "     expected result : ", "true")
	fmt.Println("gotten result :", al.isPalindrome2("race a car"), "     expected result : ", "false")
	fmt.Println("gotten result :", al.isPalindrome2(" "), "       expected result : ", "true")
	fmt.Println("gotten result :", al.isPalindrome2("0P"), "      expected result : ", "false")
	fmt.Println("End  isPalindrome ============")
	fmt.Println()

	fmt.Println("Start sum of two ============")

	fmt.Println("gotten result :", al.twoSumSort([]int{2, 7, 11, 15}, 9), "     expected result : ", "[1 2]")
	fmt.Println("gotten result :", al.twoSumSort([]int{2, 3, 4}, 6), "     expected result : ", "[1 3")
	fmt.Println("gotten result :", al.twoSumSort([]int{-1, 0}, -1), "       expected result : ", "[1 2]")

	fmt.Println("End  sum of two ============")

	fmt.Println()

	fmt.Println("Start sum of Three============")

	fmt.Println("gotten result :", al.threeSum([]int{-1, 0, 1, 2, -1, -4}), "     expected result : ", "[[-1,-1,2],[-1,0,1]]")
	fmt.Println("gotten result :", al.threeSum([]int{0, 1, 1}), "     expected result : ", "[[]]")
	fmt.Println("gotten result :", al.threeSum([]int{0, 0, 0}), "       expected result : ", "[[0,0,0]]")

	fmt.Println("End  sum of three ============")

	fmt.Println("Start  Maxrea============")

	fmt.Println("gotten result :", al.maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), "     expected result : ", "49")
	fmt.Println("gotten result :", al.maxArea([]int{1, 1}), "     expected result : ", "1")

	fmt.Println("gotten result :", al.maxAreaPointers([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), "     expected result : ", "49")
	fmt.Println("gotten result :", al.maxAreaPointers([]int{1, 1}), "     expected result : ", "1")
	// fmt.Println("gotten result :", al.maxArea([]int{0, 0, 0}), "       expected result : ", "[[0,0,0]]")

	fmt.Println("End  MaxArea============")
	fmt.Println()

	fmt.Println("Start  Trap============")

	fmt.Println("gotten result :", al.trap([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), "     expected result : ", "19")
	fmt.Println("gotten result :", al.trap([]int{1, 1}), "     expected result : ", "0")

	fmt.Println("gotten result :", al.trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), "     expected result : ", "6")
	fmt.Println("gotten result :", al.trap([]int{4, 2, 0, 3, 2, 5}), "     expected result : ", "9")
	// fmt.Println("gotten result :", al.maxArea([]int{0, 0, 0}), "       expected result : ", "[[0,0,0]]")

	fmt.Println("End  Trap============")
	fmt.Println()
}
