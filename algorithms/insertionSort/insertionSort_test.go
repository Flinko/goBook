package insertionSort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		A []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{2, 2, 2, 2, 2, 2, 2, 2}},
		{[]int{2, 2, 2, 2, 2, 2, 2, 2}},
		{[]int{4, 3, 2, 10, 12, 1, 5, 6}},
	}

	fmt.Printf("Alina")

	for _, test := range tests {
		result := insertionSort(test.A)
		for i := 1; i < len(result); i++ {
			if result[i-1] > result[i] {
				t.Errorf("sorting worked bad for input: %v got output: %v", test.A, result)
			}
		}
	}
	fmt.Printf("end")
}
