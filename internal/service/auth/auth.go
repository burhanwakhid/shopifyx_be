package authService

import (
	"context"
	"fmt"

	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/burhanwakhid/shopifyx_backend/internal/repository"
	"github.com/burhanwakhid/shopifyx_backend/pkg/bcrypt"
	"github.com/burhanwakhid/shopifyx_backend/pkg/uuid"
)

type Service struct {
	authRepo repository.AuthRepository
}

func NewAuthService(
	authRepo repository.AuthRepository,
) *Service {
	return &Service{
		authRepo: authRepo,
	}
}

func (s *Service) RegisterUser(ctx context.Context, user request.RegisterRequest) (*entity.LoginData, error) {
	var usr entity.User

	hashedPassword, err := bcrypt.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	usr.Username = user.Username
	usr.Name = user.Name
	usr.Password = hashedPassword
	usr.ID = uuid.GenerateV4()

	dataUser, errs := s.authRepo.RegisterUser(ctx, usr)

	if errs != nil {
		fmt.Printf("ini erroer: %s", err)
		return nil, errs
	}

	return dataUser.ToLoginData(), nil
}

func (s *Service) LoginUser(ctx context.Context, username, password string) (*entity.LoginData, error) {

	usr, err := s.authRepo.LoginUser(ctx, username, password)

	if err != nil {
		return nil, err
	}

	return usr.ToLoginData(), nil

}
