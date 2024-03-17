package service

import (
	"context"
	"mime/multipart"

	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
)

type AuthService interface {
	RegisterUser(ctx context.Context, user request.RegisterRequest) (*entity.LoginData, error)
	LoginUser(ctx context.Context, username, password string) (*entity.LoginData, error)
}

type BankService interface {
	CreateBank(ctx context.Context, bank request.BankRequest, userId string) error
	GetBank(ctx context.Context, idUser string) ([]*entity.Bank, error)
	UpdateBank(ctx context.Context, bank request.BankRequest, bankId string) (*entity.Bank, error)
	DeleteBank(ctx context.Context, idBank string) error
}

type ProductService interface {
	// Product Management
	CreateProduct(ctx context.Context, product request.Product, userId string) error
	UpdateProduct(ctx context.Context, product request.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, idProduct string) error

	// Product Page
	GetProduct(ctx context.Context, idOwner string) ([]*entity.Product, error)
	GetProductById(ctx context.Context, idProduct string) (*entity.ProductDetail, error)
}

type ImageService interface {
	UploadImage(form *multipart.FileHeader) (string, error)
}
