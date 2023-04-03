package delivery

import (
	"go-programming-microservices-into-the-specialization/session_03/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(usecase *usecase.BookUseCase) error {
	r := mux.NewRouter()

	bookHandler := NewBookHandler(usecase)

	r.HandleFunc("/books", bookHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", bookHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/books/{id}", bookHandler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/books/{id}", bookHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/books", bookHandler.GetAll).Methods(http.MethodGet)

	return http.ListenAndServe(":8080", r)
}
