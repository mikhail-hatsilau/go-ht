package main

import (
	"fmt"
)

// Specify Filter function here

func Filter(array []int, predicate func(elem, _ int) bool) []int {
	var result []int

	for index, val := range array {
		if predicate(val, index) == true {
			result = append(result, val)
		}
	}
	return result
}

func main() {

	fmt.Println(Filter([]int{1, 2, 3, 4, 5, 6, 7}, func(elem, _ int) bool { return elem > 0 }))
}
