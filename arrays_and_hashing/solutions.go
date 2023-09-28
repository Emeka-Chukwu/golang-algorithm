package arraysandhashing

import (
	"fmt"
)

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
	fmt.Println("Running  TwoSums  ============")
	fmt.Println("gotten result :", arr.twoSums([]int{2, 7, 11, 15}, 9), "expected result : [0 1]")
	fmt.Println("gotten result :", arr.twoSums([]int{3, 2, 4}, 6), "expected result : [1 2]")
	fmt.Println("gotten result :", arr.twoSums([]int{3, 3}, 6), "expected result : [0 1]")
	fmt.Println("gotten result :", arr.twoSums([]int{2, 1, 5, 3}, 4), "expected result : [1 3]")
	// fmt.Println("gotten result :", arr.isAnagram(s, t), "expected result : true")
	fmt.Println("End TwoSums ============")
	fmt.Println()

	fmt.Println()
	fmt.Println("Running  GroupAnagram  ============")
	fmt.Println("gotten result :", arr.groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), "expected result : ", "[[bat],[nat,tan],[ate,eat,tea]]")
	fmt.Println("gotten result :", arr.groupAnagrams([]string{""}), "expected result : [[]]")
	fmt.Println("gotten result :", arr.groupAnagrams([]string{"a"}), "expected result : [[a]]")

	fmt.Println("gotten result :", arr.groupAnagramsWithoutsort([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), "expected result : ", "[[bat],[nat,tan],[ate,eat,tea]]")
	fmt.Println("gotten result :", arr.groupAnagramsWithoutsort([]string{""}), "expected result : [[]]")
	fmt.Println("gotten result :", arr.groupAnagramsWithoutsort([]string{"a"}), "expected result : [[a]]")
	// fmt.Println("gotten result :", arr.isAnagram(s, t), "expected result : true")
	fmt.Println("End GroupAnagram ============")
	fmt.Println()

	fmt.Println()
	fmt.Println("Running  Top Frequency  ============")

	fmt.Println("gotten result :", arr.topKFrequent([]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6}, 2), "expected result : ", "[1 2]")
	fmt.Println("gotten result :", arr.topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), "expected result : ", "[1 2]")
	fmt.Println("gotten result :", arr.topKFrequent([]int{1}, 1), "expected result : ", "[1]")

	fmt.Println("gotten result :", arr.topKFrequentWithoutSoriung([]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6}, 2), "expected result : ", "[1 2]")
	fmt.Println("gotten result :", arr.topKFrequentWithoutSoriung([]int{1, 1, 1, 2, 2, 3}, 2), "expected result : ", "[1 2]")
	fmt.Println("gotten result :", arr.topKFrequentWithoutSoriung([]int{1}, 1), "expected result : ", "[1]")

	// fmt.Println("gotten result :", arr.isAnagram(s, t), "expected result : true")
	fmt.Println("End Top Frequency============")
	fmt.Println()

	fmt.Println()
	fmt.Println("Running   Product of Array Except Self	============")

	fmt.Println("gotten result :", arr.productExceptSelf([]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6}), "expected result : ", "[1 2]")
	fmt.Println("gotten result :", arr.productExceptSelf([]int{1, 2, 3, 4}), "expected result : ", "[24,12,8,6]")
	fmt.Println("gotten result :", arr.productExceptSelf([]int{-1, 1, 0, -3, 3}), "expected result : ", "[0,0,9,0,0]")

	fmt.Println("gotten result :", arr.productExceptSelfWithoutDivider([]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 6}), "expected result : ", "[1 2]")
	fmt.Println("gotten result :", arr.productExceptSelfWithoutDivider([]int{1, 2, 3, 4}), "expected result : ", "[24,12,8,6]")
	fmt.Println("gotten result :", arr.productExceptSelfWithoutDivider([]int{-1, 1, 0, -3, 3}), "expected result : ", "[0,0,9,0,0]")

	// fmt.Println("gotten result :", arr.isAnagram(s, t), "expected result : true")
	fmt.Println("End  Product of Array Except Self ============")
	fmt.Println()

	fmt.Println("Running  Sokudo	============")

	fmt.Println("gotten result :", arr.isValidSudokuL([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}), "expected result : ", "true")

	fmt.Println("gotten result :", arr.isValidSudokuL([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}), "expected result : ", "false")
	// fmt.Println("gotten result :", arr.isAnagram(s, t), "expected result : true")
	fmt.Println("End  Sokudo ============")
	fmt.Println()

	fmt.Println()
	fmt.Println("Running   longestConsecutive	============")

	fmt.Println("gotten result :", arr.longestConsecutive([]int{100, 4, 200, 1, 3, 2}), "expected result : ", "4")
	fmt.Println("gotten result :", arr.longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), "expected result : ", "9")
	fmt.Println("gotten result :", arr.longestConsecutive([]int{0, 0}), "expected result : ", "1")
	fmt.Println("gotten result :", arr.longestConsecutive([]int{1, 2, 0, 1}), "expected result : ", "3")
	fmt.Println("End  longestConsecutive ============")
	fmt.Println()
}
