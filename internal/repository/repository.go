package repository

import (
	"context"

	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
)

type AuthRepository interface {
	RegisterUser(ctx context.Context, user entity.User) (*entity.User, error)
	LoginUser(ctx context.Context, username, password string) (*entity.User, error)
}

type BankRepository interface {
	CreateBank(ctx context.Context, bank entity.Bank, userId string) error
	GetBank(ctx context.Context, idUser string) ([]*entity.Bank, error)
	UpdateBank(ctx context.Context, bank entity.Bank) (*entity.Bank, error)
	DeleteBank(ctx context.Context, idBank string) error
}

type ProductRepository interface {
	// Product Management
	CreateProduct(ctx context.Context, product entity.Product, userId string) error
	UpdateProduct(ctx context.Context, product entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, idProduct string) error

	// Product Page
	GetProduct(ctx context.Context, request request.Product, idOwner string) ([]*entity.Product, error)
	GetProductById(ctx context.Context, idProduct string) (*entity.ProductDetail, error)
}

type PaymentRepository interface {
	BuyProduct(ctx context.Context) error
}

type StockRepository interface {
	UpdateStock(ctx context.Context) error
}

type ImageUploadRepository interface {
	UploadImage(ctx context.Context) error
}
