package circular

import "fmt"

type CircularAlgo struct {
}

func (cl CircularAlgo) Run() {

	fmt.Println("Start  moveElementsToRight ============")

	fmt.Println("gotten result :", cl.moveElementsToRight([]int{7, 1, 5, 3, 6, 4}, 2), "     expected result : ", "[6 4 7 1 5 3]")
	fmt.Println("gotten result :", cl.moveElementsToRight([]int{7, 6, 4, 3, 1}, 1), "     expected result : ", "[1 7 6 4 3]")

	fmt.Println("End  moveElementsToRight============")
	fmt.Println()

	fmt.Println("Start  moveElementsToLeft ============")

	fmt.Println("gotten result :", cl.moveElementsToLeft([]int{7, 1, 5, 3, 6, 4}, 2), "     expected result : ", "[5 3 6 4 7 1 ]")
	fmt.Println("gotten result :", cl.moveElementsToLeft([]int{7, 6, 4, 3, 1}, 1), "     expected result : ", "[6 4 3 1 7]")

	// fmt.Println("gotten result :", al.maxArea([]int{0, 0, 0}), "       expected result : ", "[[0,0,0]]")

	fmt.Println("End  moveElementsToLeft============")
	fmt.Println()

}
