package bubbleSort

import (
	"reflect"
	"testing"
)

func TestMyBubbleSort(t *testing.T) {
	// arr := []int{10, 23, 65, 34, 51, 46}
	// result, err := MyBubbleSort(Ascending, arr)
	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// }
	// expected := []int{10, 23, 34, 46, 51, 65}
	// if result != expected {
	// 	t.Errorf("Failed: Expected %d, got %d", expected, result)
	// }

	tests := []struct {
		name     string
		sortFunc func([]int, int) []int
		arr      []int
		expected []int
	}{
		{
			name:     "Ascending Sort",
			sortFunc: Ascending,
			arr:      []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name:     "Descending Sort",
			sortFunc: Descending,
			arr:      []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{9, 6, 5, 5, 5, 4, 3, 3, 2, 1, 1},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MyBubbleSort(tt.sortFunc, tt.arr)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unexpected result. Got: %v, Expected: %v", result, tt.expected)
			}
		})
	}
}
