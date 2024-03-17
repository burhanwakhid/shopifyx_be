package bank

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/burhanwakhid/shopifyx_backend/pkg/orm"
)

type BankRepository struct {
	db *orm.Replication
}

func NewBankRepository(db *orm.Replication) *BankRepository {
	return &BankRepository{db: db}
}
func (r *BankRepository) CreateBank(ctx context.Context, bank entity.Bank, userId string) error {
	tx, err := r.db.BeginTransaction(ctx)
	if err != nil {
		fmt.Printf("ini error 0: %s ", err)
		return err
	}

	// Build the SQL query with placeholders
	query := "INSERT INTO bank ( name, account_name, account_number, id_user) VALUES ($1, $2, $3, $4)"

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		_ = tx.Rollback() // rollback on prepare error

		fmt.Printf("ini error 1: %s ", err)

		return err
	}
	defer stmt.Close() // ensure statement is closed

	// Execute the query with the transaction and scanned arguments
	_, err = stmt.ExecContext(ctx, bank.Name, bank.AccountName, bank.AccountNumber, userId)
	if err != nil {
		_ = tx.Rollback() // rollback on insert error
		fmt.Printf("ini error 2: %s ", err)
		return err
	}

	// Commit the transaction if successful
	if err := tx.Commit(); err != nil {
		fmt.Printf("ini error 3: %s ", err)
		return err
	}

	return nil
}

func (r *BankRepository) GetBank(ctx context.Context, idUser string) ([]*entity.Bank, error) {
	query := "SELECT * FROM bank WHERE id_user = $1"
	stmt, err := r.db.GetInstance(orm.WorkloadRead).PrepareContext(ctx, query)
	if err != nil {
		fmt.Printf("ini error 0: %s ", err)
		return nil, err
	}
	defer stmt.Close() // ensure statement is closed

	rows, err := stmt.QueryContext(ctx, idUser)
	if err != nil {
		fmt.Printf("ini error 1: %s ", err)
		return nil, err
	}
	defer rows.Close() // ensure rows are closed

	var banks []*entity.Bank
	for rows.Next() {
		var b entity.Bank
		err := rows.Scan(&b.Id, &b.Name, &b.AccountName, &b.AccountNumber, &b.IdUser)
		if err != nil {
			fmt.Printf("ini error 2: %s ", err)
			if errors.Is(err, sql.ErrNoRows) {
				return nil, entity.ErrDataNotFound
			}
			return nil, err
		}
		banks = append(banks, &b)
	}

	fmt.Printf("Result Length: %d", len(banks))

	return banks, nil
}

func (r *BankRepository) UpdateBank(ctx context.Context, bank entity.Bank) (*entity.Bank, error) {
	query := "UPDATE bank SET name = $1, account_name = $2, account_number = $3 WHERE id = $4"

	tx, err := r.db.BeginTransaction(ctx)
	if err != nil {
		fmt.Printf("ini error 0: %s ", err)
		return nil, err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		_ = tx.Rollback() // rollback on prepare error
		fmt.Printf("ini error 1: %s ", err)
		return nil, err
	}
	defer stmt.Close() // ensure statement is closed

	result, err := stmt.ExecContext(ctx, bank.Name, bank.AccountName, bank.AccountNumber, bank.Id)
	if err != nil {
		_ = tx.Rollback() // rollback on update error
		fmt.Printf("ini error 2: %s ", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		_ = tx.Rollback() // rollback on error getting rows affected
		fmt.Printf("ini error 3: %s ", err)
		return nil, err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows affected during update")
		_ = tx.Rollback() // rollback if no rows affected
		return nil, entity.ErrDataNotFound
	}

	// Commit the transaction if successful
	if err := tx.Commit(); err != nil {
		fmt.Printf("ini error 4: %s ", err)
		return nil, err
	}

	return &bank, nil
}

func (r *BankRepository) DeleteBank(ctx context.Context, idBank string) error {
	query := "DELETE FROM bank WHERE id = $1"

	tx, err := r.db.BeginTransaction(ctx)
	if err != nil {
		fmt.Printf("ini error 0: %s ", err)
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		_ = tx.Rollback() // rollback on prepare error
		fmt.Printf("ini error 1: %s ", err)
		return err
	}
	defer stmt.Close() // ensure statement is closed

	result, err := stmt.ExecContext(ctx, idBank)
	if err != nil {
		_ = tx.Rollback() // rollback on delete error
		fmt.Printf("ini error 2: %s ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		_ = tx.Rollback() // rollback on error getting rows affected
		fmt.Printf("ini error 3: %s ", err)
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("No rows affected during delete")
		_ = tx.Rollback() // rollback if no rows affected
		return entity.ErrDataNotFound
	}

	// Commit the transaction if successful
	if err := tx.Commit(); err != nil {
		fmt.Printf("ini error 4: %s ", err)
		return err
	}

	return nil
}
