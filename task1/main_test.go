package main

import (
	"reflect"
	"testing"
)

// Tests filter function
func TestFilterFn(t *testing.T) {
	testTables := []struct {
		array     []int
		predicate func(int, int) bool
		result    []int
	}{
		{
			array:     []int{1, 2, 3, 4, 5},
			predicate: func(elem, _ int) bool { return elem == 3 },
			result:    []int{3},
		},
		{
			array:     []int{},
			predicate: func(elem, _ int) bool { return elem == 3 },
		},
		{
			array:     []int{1, 2, 4, 5, 3},
			predicate: func(elem, _ int) bool { return elem > 0 },
			result:    []int{1, 2, 4, 5, 3},
		},
	}

	for _, table := range testTables {
		result := Filter(table.array, table.predicate)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Expected result to be equal %d, but got %d", table.result, result)
		}
	}
}
