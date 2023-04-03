package repository

import (
	"database/sql"
	"go-programming-microservices-into-the-specialization/session_03/domain"

	"github.com/pkg/errors"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *BookRepository) Create(book *domain.Book) error {
	query := `
        INSERT INTO books (title, author, publisher, publish_date)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at
    `
	err := r.db.QueryRow(query, book.Title, book.Author, book.Publisher, book.PublishDate).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return errors.Wrap(err, "BookRepository.Create")
	}
	return nil
}

func (r *BookRepository) Update(book *domain.Book) error {
	query := `UPDATE books SET title=$1, author=$2, publisher=$3, publish_date=$4, updated_at=NOW() WHERE id=$5`
	_, err := r.db.Exec(query, book.Title, book.Author, book.Publisher, book.PublishDate, book.ID)
	if err != nil {
		return errors.Wrap(err, "BookRepository.Update")
	}
	return nil
}

func (r *BookRepository) Delete(id int) error {
	query := `DELETE FROM books WHERE id=$1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "BookRepository.Delete")
	}
	return nil
}

func (r *BookRepository) GetByID(id int) (*domain.Book, error) {
	query := `SELECT id, title, author, publisher, publish_date, created_at, updated_at FROM books WHERE id=$1`
	book := &domain.Book{}
	err := r.db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.PublishDate, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "BookRepository.GetByID")
	}
	return book, nil
}

func (r *BookRepository) GetAll() ([]*domain.Book, error) {
	query := `SELECT id, title, author, publisher, publish_date, created_at, updated_at FROM books`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "BookRepository.GetAll")
	}
	defer rows.Close()
	books := []*domain.Book{}
	for rows.Next() {
		book := &domain.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publisher, &book.PublishDate, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "BookRepository.GetAll")
		}
		books = append(books, book)
	}

	return books, nil
}
