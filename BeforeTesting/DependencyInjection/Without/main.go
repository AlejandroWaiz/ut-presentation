package main

import (
	"log"
	secretary "testing/BeforeTesting/DependencyInjection/Without/secretary"
)

func main() {

	secretary, err := secretary.NewSecretary()

	if err != nil {
		panic(err)
	}

	_, err = secretary.SaveSomeData("data")

	if err != nil {
		log.Println(err)
	}
}
