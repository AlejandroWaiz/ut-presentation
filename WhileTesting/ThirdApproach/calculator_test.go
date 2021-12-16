package calculator

import "testing"

func TestCalculate(t *testing.T) {

	testTable := []struct {
		name, operation                        string
		FirstOperand, SecondOperand, expResult int
		expErr                                 bool
	}{
		{"Normal add", "add", 1, 2, 3, false},
		{"Normal substract", "substract", 3, 2, 1, false},
		{"Normal multiply", "multiply", 1, 2, 2, false},
		{"Normal divide", "divide", 2, 2, 1, false},
		{"Err divide", "divide", 2, 0, 0, true},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {

			result, err := Calculate(testCase.operation, testCase.FirstOperand, testCase.SecondOperand)

			if result != testCase.expResult {
				t.Errorf("Expecting %v as a result but got %v instead", testCase.expResult, result)
			}

			if err != nil && testCase.expErr == false {
				t.Errorf("Not expecting err but got %v instead", err)
			}

		})

	}

}
