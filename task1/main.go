package main

import "fmt"

func check(item int, index int) bool {
	return item%2 == 0
}

func Filter(list []int, check func(int, int) bool) (result []int) {
	for index, item := range list {
		if check(item, index) {
			result = append(result, item)
		}
	}
	return
}

func main() {
	input := []int{1, 3, 2, 4, 7, 6, 0}
	fmt.Println(Filter(input, check))
}
