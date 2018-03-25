package list

import (
	"strings"
	"testing"
)

type Object struct {
	name  string
	age   int
	marks []float32
}

func TestAdjust(t *testing.T) {

	var testArr1 = []int{1, 2, 3, 4, 5, 6, 7}
	var testArr2 = []string{"Matt", "Joanna", "Helga", "Levy", "Alice"}

	square := func(x int) int {
		return x * x
	}

	// Test 1
	case1 := AdjustInt(square, 2, testArr1)
	if case1[2] != 9 {
		t.Errorf("Expected %d Got %d \n", 9, case1)
	}
	// Test 2
	case2 := AdjustStr(strings.ToUpper, 2, testArr2)
	if case2[2] != "HELGA" {
		t.Errorf("Expected %s Got %s\n", "HELGA", case2)
	}

}
