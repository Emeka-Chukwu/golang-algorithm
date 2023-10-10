package main

import (
	arraysandhashing "ba/algorithms/arrays_and_hashing"
	"ba/algorithms/circular"
	slidingwindow "ba/algorithms/slidingWindow"
	"ba/algorithms/sorting"
	twopointers "ba/algorithms/two_pointers"
)

func main() {
	arr := arraysandhashing.ArrayAlg{}
	st := sorting.SortingAl{}
	two := twopointers.AlTwoPointers{}
	cir := circular.CircularAlgo{}
	sliding := slidingwindow.SlidingWindowAlg{}
	arr.Run()
	st.Run()
	two.Run()
	cir.Run()
	sliding.Run()

}
