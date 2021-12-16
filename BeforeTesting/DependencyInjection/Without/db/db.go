package db

import "errors"

type SuperDb struct {
}

func NewDB() (*SuperDb, error) {

	return nil, errors.New("Not connected to database")

}

func (db *SuperDb) SaveSomeData(data string) (response string, err error) {

	if data == "" {
		return "", errors.New("Error saving the data")
	}

	return "Data saved with success", nil

}
