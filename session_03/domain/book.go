package domain

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publish_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookRepository interface {
	Create(book *Book) error
	Update(book *Book) error
	Delete(id int) error
	GetByID(id int) (*Book, error)
	GetAll() ([]*Book, error)
}
