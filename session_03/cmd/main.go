package main

import (
	"log"

	"go-programming-microservices-into-the-specialization/session_03/delivery"
	"go-programming-microservices-into-the-specialization/session_03/repository"
	"go-programming-microservices-into-the-specialization/session_03/usecase"
)

func main() {
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	bookUseCase := usecase.NewBookUseCase(bookRepo)
	err = delivery.StartServer(bookUseCase)
	if err != nil {
		log.Fatal(err)
	}
}
