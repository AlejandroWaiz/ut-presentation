package main

import (
	"log"
	db "testing/BeforeTesting/DependencyInjection/With/db"
	"testing/BeforeTesting/DependencyInjection/With/secretary"
)

func main() {

	db, err := db.NewDB()

	if err != nil {
		panic(err)
	}

	secretary := secretary.NewSecretary(db)

	_, err = secretary.SaveSomeData("data")

	if err != nil {
		log.Println(err)
	}
}
