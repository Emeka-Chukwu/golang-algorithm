package sorting

import "fmt"

type SortingAl struct {
}

func (al SortingAl) Run() {
	fmt.Println(al.bubbleSort([]int{13, 3, 2, 7, 8, 1, 9, 1, 4, 2}))
	fmt.Println(al.bubbleSort([]int{13, 3, 2, 9, 50}))
	fmt.Println(al.bubbleSort([]int{1, 2, 3, 4, 5}))

	// Quicksorting
	fmt.Println(al.quickSort([]int{13, 3, 2, 7, 8, 1, 9, 1, 4, 2}))
	fmt.Println(al.quickSort([]int{13, 3, 2, 9, 50}))
	fmt.Println(al.quickSort([]int{1, 2, 3, 4, 5}))
}
