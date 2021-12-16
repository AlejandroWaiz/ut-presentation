package calculator

import "testing"

func TestCalculate(t *testing.T) {

	testTable := []struct {
		operation                              string
		FirstOperand, SecondOperand, expResult int
	}{
		{"add", 1, 2, 3},
		{"substract", 3, 2, 1},
		{"multiply", 1, 2, 2},
		{"divide", 2, 2, 1},
	}

	for _, testCase := range testTable {

		result, _ := Calculate(testCase.operation, testCase.FirstOperand, testCase.SecondOperand)

		if result != testCase.expResult {
			t.Errorf("Expecting %v as a result but got %v instead", testCase.expResult, result)
		}

	}

}
