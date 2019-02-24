package main

func Filter(numbers []int, predicate func(item int, index int) bool) []int {
	var result []int
	
	for i, number := range numbers {
		if predicate(number, i) {
			result = append(result, number)
		}
	}
	
	return result;
}

func main() {

}
