package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"go-programming-microservices-into-the-specialization/session_03/domain"
	"go-programming-microservices-into-the-specialization/session_03/usecase"
)

type BookHandler struct {
	usecase *usecase.BookUseCase
}

func NewBookHandler(usecase *usecase.BookUseCase) *BookHandler {
	return &BookHandler{usecase: usecase}
}

func (h *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	book := &domain.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.Create").Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = h.usecase.Create(book)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.Create").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	book := &domain.Book{}
	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.Update").Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = h.usecase.Update(book)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.Update").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.Delete").Error(), http.StatusBadRequest)
		return
	}

	err = h.usecase.Delete(id)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.Delete").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.GetByID").Error(), http.StatusBadRequest)
		return
	}

	book, err := h.usecase.GetByID(id)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.GetByID").Error(), http.StatusInternalServerError)
		return
	}

	if book == nil {
		http.NotFound(w, r)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.GetByID").Error(), http.StatusInternalServerError)
		return
	}
}

func (h *BookHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	books, err := h.usecase.GetAll()
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.GetAll").Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, errors.Wrap(err, "BookHandler.GetAll").Error(), http.StatusInternalServerError)
		return
	}
}
