package main

import "fmt"

type predicate func(int, int) bool

func Filter(slice []int, function predicate) []int {
	var result []int
	for i, item := range slice {
		if function(item, i) == true {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	fmt.Println("Even", Filter([]int{1, 2, 3, 4, 5},
		func(item, index int) bool { return item%2 == 0 }))
}
