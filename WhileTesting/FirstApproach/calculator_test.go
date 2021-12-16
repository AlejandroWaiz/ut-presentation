package calculator

import "testing"

func TestCalculator(t *testing.T) {

	result, _ := Calculate("add", 1, 2)

	if result != 3 {
		t.Errorf("Expecting 3 but got %v instead", result)
	}

	result, _ = Calculate("substract", 3, 2)

	if result != 1 {
		t.Errorf("Expecting 1 but got %v instead", result)
	}

	result, _ = Calculate("multiply", 1, 1)

	if result != 1 {
		t.Errorf("Expecting 1 but got %v instead", result)
	}

	result, _ = Calculate("divide", 2, 2)

	if result != 1 {
		t.Errorf("Expecting 1 but got %v instead", result)
	}

	_, err := Calculate("divide", 2, 0)

	if err == nil {
		t.Error("Expecting err but got nil instead")
	}

}
