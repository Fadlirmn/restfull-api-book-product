package services

import(
	"go-roadmap/models"
	"go-roadmap/repository"
)


type ProductService struct{
	repo repository.ProductRepository
}
func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) GetProducts()[]models.Product  {
	return  s.repo.FindAllProduct()
}

func (s *ProductService) CreateProduct(Product models.Product)  {
	s.repo.SaveProduct(Product)
}

func (s *ProductService) UpdateProduct(id int,Product models.Product)  error{
	return s.repo.UpdateProduct(id, Product)
}
func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}