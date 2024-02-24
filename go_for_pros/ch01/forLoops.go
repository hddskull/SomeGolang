package main

import "fmt"

func forLoops() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i*i, " ")
	//}
	//fmt.Println()

	//i := 0
	//for ok := true; ok; ok = (i != 10) {
	//	fmt.Println(i*i, " ")
	//	i++
	//}
	//fmt.Println()

	//i := 0
	//for {
	//	if i == 10 {
	//		break
	//	}
	//	i++
	//	fmt.Println()
	//}

	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index: ", i, "| value: ", v)
	}
}
