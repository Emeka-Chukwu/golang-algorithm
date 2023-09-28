package twopointers

import "fmt"

type AlTwoPointers struct{}

func (al AlTwoPointers) Run() {
	fmt.Println()
	fmt.Println("Start isPalindrome ============")

	fmt.Println("gotten result :", al.isPalindrome("A man, a plan, a canal: Panama"), "     expected result : ", "true")
	fmt.Println("gotten result :", al.isPalindrome("race a car"), "     expected result : ", "false")
	fmt.Println("gotten result :", al.isPalindrome(" "), "       expected result : ", "true")
	fmt.Println("gotten result :", al.isPalindrome("0P"), "      expected result : ", "false")
	fmt.Println("End  isPalindrome ============")
	fmt.Println()

	fmt.Println("Start sum of two ============")

	fmt.Println("gotten result :", al.twoSum([]int{2, 7, 11, 15}, 9), "     expected result : ", "[1 2]")
	fmt.Println("gotten result :", al.twoSum([]int{2, 3, 4}, 6), "     expected result : ", "[1 3")
	fmt.Println("gotten result :", al.twoSum([]int{-1, 0}, -1), "       expected result : ", "[1 2]")

	fmt.Println("End  sum of two ============")
	fmt.Println()
}
