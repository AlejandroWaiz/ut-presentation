package secretary

import (
	"testing"
)

func TestNewSecretary(t *testing.T) {

	_, err := NewSecretary()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestSaveSomeData(t *testing.T) {

	secretary, err := NewSecretary()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	_, err = secretary.SaveSomeData("Test data")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
