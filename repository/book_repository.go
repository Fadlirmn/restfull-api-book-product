package repository

import (
	"go-roadmap/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type BookRepository interface {
	FindAllBook() []models.Book
	SaveBook(Book models.Book)
	UpdateBook(id int, Book models.Book) error
	DeleteBook(id int) error
}

type BookRepo struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) BookRepository {
	return &BookRepo{db: db}
}

func (r *BookRepo) FindAllBook() []models.Book {
	var books []models.Book
	 err := r.db.Select(&books,"SELECT id, name_book, item, type FROM books")

	if err != nil {
		log.Println("error query", err)
		return nil
	}
	return books
}

func (r *BookRepo) SaveBook(book models.Book) {
	_, err := r.db.NamedExec(
		`INSERT INTO books(name_book,genre) VALUES (:name_book,genre)`,
		book,
	)
	if err != nil {
		log.Println("gagal Menambahkan Book", err)
	}
}

func (r *BookRepo) UpdateBook(id int, book models.Book) error {
	_, err := r.db.Exec("UPDATE books SET name_book=$1, genre=$2 WHERE id=$3", book.NameBook, book.Genre, id)
	return err
}

func (r *BookRepo) DeleteBook(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}
