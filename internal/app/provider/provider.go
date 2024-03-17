package provider

import (
	"github.com/burhanwakhid/shopifyx_backend/config"
	"github.com/burhanwakhid/shopifyx_backend/internal"
	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/handler"
	"github.com/burhanwakhid/shopifyx_backend/internal/repository"
	"github.com/burhanwakhid/shopifyx_backend/internal/repository/auth"
	"github.com/burhanwakhid/shopifyx_backend/internal/repository/bank"
	"github.com/burhanwakhid/shopifyx_backend/internal/repository/product"
	"github.com/burhanwakhid/shopifyx_backend/internal/service"
	authService "github.com/burhanwakhid/shopifyx_backend/internal/service/auth"
	bankService "github.com/burhanwakhid/shopifyx_backend/internal/service/bank"
	imageService "github.com/burhanwakhid/shopifyx_backend/internal/service/image"
	productService "github.com/burhanwakhid/shopifyx_backend/internal/service/product"
	validates "github.com/burhanwakhid/shopifyx_backend/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var (
	PkgSet = wire.NewSet(
		ProvideErrorTranslator,
		config.GetDBConfig,
		ProvideLocalCache,
		ProviderValidator,
		wire.FieldsOf(new(*config.AppConfig), "TranslatorConfig"),
	)

	RepoSet = wire.NewSet(
		ProvideDBReplication,
		auth.NewUserRepository,
		wire.Bind(new(repository.AuthRepository), new(*auth.Repository)),
		bank.NewBankRepository,
		wire.Bind(new(repository.BankRepository), new(*bank.BankRepository)),
		product.NewProductRepository,
		wire.Bind(new(repository.ProductRepository), new(*product.ProductRepository)),
	)

	ServiceSet = wire.NewSet(
		authService.NewAuthService,
		wire.Bind(new(service.AuthService), new(*authService.Service)),
		bankService.NewBankService,
		wire.Bind(new(service.BankService), new(*bankService.BankService)),
		productService.NewProductService,
		wire.Bind(new(service.ProductService), new(*productService.ProductService)),
		imageService.NewImageService,
		wire.Bind(new(service.ImageService), new(*imageService.ImageService)),
	)

	// GRPCDeliverySet = wire.NewSet()

	RestDeliverySet = wire.NewSet(
		handler.NewHandlerContainer,
		handler.NewBankHandler,
		handler.NewProductHandler,
		handler.NewImageHandler,
		ProviderAuthHandler,
	)

	ConfigSet = wire.NewSet(
		config.GetAppConfig,
		config.GetVendorConfig,
	)

	// GRPCMiddlewareSet = wire.NewSet(

	// )

	// RESTMiddlewareSet = wire.NewSet(

	// )
)

func ProviderAuthHandler(
	service service.AuthService,
	errorTranslator internal.ErrorTranslator,
	validate validates.JSONValidator,
) *handler.AuthHandler {
	return handler.NewAuthHandler(service, errorTranslator, validate)
}

func ProviderValidator() validates.JSONValidator {
	s := validator.New()

	return *validates.NewValidator(s)
}
