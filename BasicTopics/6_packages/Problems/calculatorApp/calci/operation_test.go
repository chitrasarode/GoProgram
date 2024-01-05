package calci

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	result, err := Addition(3, 9)
	if err != nil {
		fmt.Println("Error in addition : ", err)
	}
	expected := 12
	if result != expected {
		t.Errorf("Failed: Expected %d, got %d", expected, result)
	}
}

func BenchmarkAddition(b *testing.B) {
	// Run the function b.N times
	for i := 0; i < b.N; i++ {
		_, err := Addition(3, 2) // Adjust the input as needed
		if err != nil {
			fmt.Println("Error :", err)
		}

	}
}
