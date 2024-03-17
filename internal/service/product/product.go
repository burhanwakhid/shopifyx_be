package productService

import (
	"context"

	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/burhanwakhid/shopifyx_backend/internal/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}

// Product Management
func (s *ProductService) CreateProduct(ctx context.Context, product request.Product, userId string) error {
	prod := entity.Product{
		Name:          product.Name,
		Price:         product.Price,
		ImageUrl:      product.ImageUrl,
		Stock:         product.Stock,
		Condition:     product.Condition,
		Tags:          product.Tags,
		IsPurchasable: product.IsPurchasable,
	}

	err := s.repository.CreateProduct(ctx, prod, userId)

	if err != nil {
		return err
	}

	return nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, product request.Product) (*entity.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ProductService) DeleteProduct(ctx context.Context, idProduct string) error {
	panic("not implemented") // TODO: Implement
}

// Product Page
func (s *ProductService) GetProduct(ctx context.Context, idOwner string) ([]*entity.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (s *ProductService) GetProductById(ctx context.Context, idProduct string) (*entity.ProductDetail, error) {
	panic("not implemented") // TODO: Implement
}
