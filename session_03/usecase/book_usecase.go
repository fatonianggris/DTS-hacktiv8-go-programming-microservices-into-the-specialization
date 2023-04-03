package usecase

import (
	"go-programming-microservices-into-the-specialization/session_03/domain"

	"github.com/pkg/errors"
)

type BookUseCase struct {
	repo domain.BookRepository
}

func NewBookUseCase(repo domain.BookRepository) *BookUseCase {
	return &BookUseCase{repo: repo}
}

func (u *BookUseCase) Create(book *domain.Book) error {
	err := u.repo.Create(book)
	if err != nil {
		return errors.Wrap(err, "BookUseCase.Create")
	}
	return nil
}

func (u *BookUseCase) Update(book *domain.Book) error {
	err := u.repo.Update(book)
	if err != nil {
		return errors.Wrap(err, "BookUseCase.Update")
	}
	return nil
}

func (u *BookUseCase) Delete(id int) error {
	err := u.repo.Delete(id)
	if err != nil {
		return errors.Wrap(err, "BookUseCase.Delete")
	}
	return nil
}

func (u *BookUseCase) GetByID(id int) (*domain.Book, error) {
	book, err := u.repo.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "BookUseCase.GetByID")
	}
	return book, nil
}

func (u *BookUseCase) GetAll() ([]*domain.Book, error) {
	books, err := u.repo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "BookUseCase.GetAll")
	}
	return books, nil
}
