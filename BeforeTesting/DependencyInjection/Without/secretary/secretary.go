package secretary

import (
	"log"
	db "testing/BeforeTesting/DependencyInjection/Without/db"
)

type Secretary struct {
	db *db.SuperDb
}

func NewSecretary() (*Secretary, error) {

	db, err := db.NewDB()

	if err != nil {
		return nil, err
	}

	S := Secretary{db: db}

	return &S, nil

}

func (s *Secretary) SaveSomeData(data string) (response string, err error) {

	response, err = s.db.SaveSomeData(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return "Data saved", nil
}
