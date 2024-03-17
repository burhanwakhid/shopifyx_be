package bankService

import (
	"context"
	"fmt"

	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/burhanwakhid/shopifyx_backend/internal/repository"
)

type BankService struct {
	bankRepo repository.BankRepository
}

func NewBankService(repo repository.BankRepository) *BankService {
	return &BankService{
		bankRepo: repo,
	}
}

func (r *BankService) CreateBank(ctx context.Context, bank request.BankRequest, userId string) error {
	var bnk entity.Bank

	bnk.AccountName = bank.BankAccountName
	bnk.AccountNumber = bank.BankAccountNumber
	bnk.Name = bank.BankName

	err := r.bankRepo.CreateBank(ctx, bnk, userId)

	if err != nil {
		return err
	}

	return nil

}

func (r *BankService) GetBank(ctx context.Context, idUser string) ([]*entity.Bank, error) {
	banks, err := r.bankRepo.GetBank(ctx, idUser)

	if err != nil {
		return nil, err
	}

	return banks, nil
}

func (r *BankService) UpdateBank(ctx context.Context, bank request.BankRequest, bankId string) (*entity.Bank, error) {
	var bnk entity.Bank

	bnk.Id = bankId
	bnk.AccountName = bank.BankAccountName
	bnk.AccountNumber = bank.BankAccountNumber
	bnk.Name = bank.BankName

	banks, err := r.bankRepo.UpdateBank(ctx, bnk)

	if err != nil {
		fmt.Println("error haha")
		return nil, err
	}

	return banks, nil

}

func (r *BankService) DeleteBank(ctx context.Context, idBank string) error {
	return r.bankRepo.DeleteBank(ctx, idBank)
}
