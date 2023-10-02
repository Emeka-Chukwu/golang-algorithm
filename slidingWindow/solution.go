package slidingwindow

import "fmt"

type SlidingWindowAlg struct {
}

func (sl SlidingWindowAlg) Run() {

	fmt.Println("Start  MaxProfits============")

	fmt.Println("gotten result :", sl.maxProfit([]int{7, 1, 5, 3, 6, 4}), "     expected result : ", "5")
	fmt.Println("gotten result :", sl.maxProfit([]int{7, 6, 4, 3, 1}), "     expected result : ", "0")

	fmt.Println("gotten result :", sl.maxProfit([]int{8, 7, 2, 9, 1, 20}), "     expected result : ", "19")

	fmt.Println("gotten result :", sl.maxProfit([]int{1, 2, 4}), "     expected result : ", "3")
	fmt.Println("gotten result :", sl.maxProfit([]int{2, 1, 2, 1, 0, 1, 2}), "     expected result : ", "2")

	// fmt.Println("gotten result :", al.maxArea([]int{0, 0, 0}), "       expected result : ", "[[0,0,0]]")

	fmt.Println("End  MaxProfits============")
	fmt.Println()

}
