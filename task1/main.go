package main

func Filter(input []int, predicate func(int, int) bool) []int {
	var output []int
	for index, value := range input {
		if (predicate(value, index)) {
			output = append(output, value)
		}
	}
	return output
}

func main() {

}
