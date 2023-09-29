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

	fmt.Println()
}
