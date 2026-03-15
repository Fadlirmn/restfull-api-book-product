package services

import(
	"go-roadmap/models"
	"go-roadmap/repository"
)


type BookService struct{
	repo repository.BookRepository
}
func NewBookService(r repository.BookRepository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) GetBooks()[]models.Book  {
	return  s.repo.FindAllBook()
}

func (s *BookService) CreateBook(Book models.Book)  {
	s.repo.SaveBook(Book)
}

func (s *BookService) UpdateBook(id int,Book models.Book)  error{
	return s.repo.UpdateBook(id, Book)
}
func (s *BookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}