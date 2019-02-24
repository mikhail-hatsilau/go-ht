package main

func Filter(input []int, filter func(int, int) bool) []int {
	var output []int
	for index, value := range input {
		if (filter(value, index)) {
			output = append(output, value)
		}
	}
	return output
}

func main() {
}
