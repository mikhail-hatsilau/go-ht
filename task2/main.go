package main

func MapTo(slice []int, f func(item, index int) string) []string {
	var result []string
	for i, elem := range slice {
		result = append(result, f(elem, i))
	}
	return result
}

func Convert(slice []int) []string {
	convertedNums := map[int]string{
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
	return MapTo(slice, func(item, index int) string {
		var result string
		if val, ok := convertedNums[item]; ok {
			result = val
		} else {
			result = "unknown"
		}
		return result
	})
}

func main() {
}
