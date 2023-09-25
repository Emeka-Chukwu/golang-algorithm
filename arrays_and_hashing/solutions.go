package arraysandhashing

import "fmt"

type ArrayAlg struct {
}

func (arr ArrayAlg) Run() {
	input := []int{1, 2, 3, 1}
	input2 := []int{1, 2, 3, 4}
	input3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println()
	fmt.Println("Running ContainsDuplivate ============")
	fmt.Println("gotten result :", arr.ContainsDuplicate(input), "expected result : true")
	fmt.Println("gotten result :", arr.ContainsDuplicate(input2), "expected result : false")
	fmt.Println("gotten result :", arr.ContainsDuplicate(input3), "expected result : true")
	fmt.Println("End ContainsDuplivate ============")
	s := "anagram"
	t := "nagaram"
	s1 := "rat"
	t1 := "car"
	fmt.Println()
	fmt.Println("Running  IsAnagram  ============")
	fmt.Println("gotten result :", arr.isAnagram(s, t), "expected result : true")
	fmt.Println("gotten result :", arr.isAnagram(s1, t1), "expected result : false")
	// fmt.Println("gotten result :", arr.isAnagram(s, t), "expected result : true")
	fmt.Println("End IsAnagram ============")
	fmt.Println()
}
