package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "books"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", getBook(db)).Methods("GET")
	router.HandleFunc("/books", createBook(db)).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM books")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		books := []Book{}
		for rows.Next() {
			book := Book{}
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			books = append(books, book)
		}
		json.NewEncoder(w).Encode(books)
	}
}

func getBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		book := Book{}
		err := db.QueryRow("SELECT * FROM books WHERE id=$1", id).Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(book)
	}
}

func createBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sqlStatement := `
        INSERT INTO books (title, author, description, price)
        VALUES ($1, $2, $3, $4)
        RETURNING id`
		id := 0
		err = db.QueryRow(sqlStatement, book.Title, book.Author, book.Description, book.Price).Scan(&id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		book.ID = id
		json.NewEncoder(w).Encode(book)
	}
}

func updateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sqlStatement := `UPDATE books SET title=$1, author=$2, description=$3, price=$4 WHERE id=$5`
		_, err = db.Exec(sqlStatement, book.Title, book.Author, book.Description, book.Price, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(book)
	}
}

func deleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		sqlStatement := `DELETE FROM books WHERE id=$1`
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
