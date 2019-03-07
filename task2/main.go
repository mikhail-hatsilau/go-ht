package main

func MapTo(slice []int, modify func(elem, index int) string) (result []string) {
	for index, elem := range slice {
		elem := modify(elem, index)
		result = append(result, elem)
	}
	return
}

var intToStr = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func Convert(arr []int) []string {
	return MapTo(arr, func(elem, index int) string {
		if elem, ok := intToStr[elem]; ok {
			return elem
		}
		return "unknown"
	})
}

func main() {
}
