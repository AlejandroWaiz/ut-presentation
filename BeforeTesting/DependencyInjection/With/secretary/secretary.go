package secretary

import (
	"log"
	db "testing/BeforeTesting/DependencyInjection/With/db"
)

type Secretary struct {
	db db.ISuperDb
}

type ISecretary interface {
	SaveSomeData(data string) (response string, err error)
}

func NewSecretary(db db.ISuperDb) *Secretary {

	S := Secretary{db: db}

	return &S

}

func (s *Secretary) SaveSomeData(data string) (response string, err error) {

	response, err = s.db.SaveSomeData(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return "Data saved", nil
}
