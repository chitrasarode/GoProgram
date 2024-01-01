package calci

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(3, 5)
	expected := 8
	if result == expected {
		fmt.Println("Passed")
	} else {
		t.Errorf("failed")
	}
}
