package main

import (
	"reflect"
	"strconv"
	"testing"
)

// Tests map function
func TestMapFn(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	cb := func(elem, _ int) string { return strconv.Itoa(elem) }
	expected := []string{"1", "2", "3", "4", "5"}

	result := MapTo(array, cb)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be equal %d, but got %d", expected, result)
	}
}

func TestConvertion(t *testing.T) {
	testsTable := []struct {
		array    []int
		expected []string
	}{
		{[]int{1, 2, 3, 4, 5}, []string{"one", "two", "three", "four", "five"}},
		{array: []int{}},
		{[]int{1, 5, 9, 10, 11}, []string{"one", "five", "nine", "unknown", "unknown"}},
	}
	for _, table := range testsTable {
		result := Convert(table.array)
		if !reflect.DeepEqual(result, table.expected) {
			t.Errorf("Expected result to be equal %d, but got %d", table.expected, result)
		}
	}
}
