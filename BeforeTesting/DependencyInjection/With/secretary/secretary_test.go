package secretary

import (
	"testing"
)

type MockDb struct{}

func (db *MockDb) SaveSomeData(data string) (response string, err error) {
	return "Good response", nil
}

func TestNewSecretary(t *testing.T) {

	secretary := NewSecretary(&MockDb{})

	if secretary == nil {
		t.Error("Error: expecting interface implementation but got nil instead")
	}

}

func TestSaveSomeData(t *testing.T) {

	s := Secretary{db: &MockDb{}}

	_, err := s.SaveSomeData("test data")

	if err != nil {
		t.Error(err)
	}

}
