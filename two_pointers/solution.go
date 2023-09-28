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
}
