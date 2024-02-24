package main

import (
	"fmt"
	"strconv"
)

func main() {
	First("a4bc2d5e")
}

func First(s string) string {
	result := ""
	lastIndex := 0
	for index, r := range s {

		number, err := strconv.Atoi(string(r))
		if err != nil {
			result = result + string(r)
			lastIndex = index
			continue
		}

		prevNum, err1 := strconv.Atoi(string(s[lastIndex]))

		fmt.Printf("%d", prevNum)

		if err1 == nil {
			fmt.Printf("%d", prevNum)
			continue
		}

		for i := 0; i < (number - 1); i++ {
			result = result + string(s[lastIndex])
		}

		lastIndex = index
	}
	fmt.Printf("\nresult = %s \n", result)

	return result
}
