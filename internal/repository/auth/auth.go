package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/burhanwakhid/shopifyx_backend/pkg/bcrypt"
	"github.com/burhanwakhid/shopifyx_backend/pkg/orm"
	"github.com/lib/pq"
)

type Repository struct {
	db *orm.Replication
}

func NewUserRepository(db *orm.Replication) *Repository {
	return &Repository{db: db}
}

func (r *Repository) RegisterUser(ctx context.Context, user entity.User) (*entity.User, error) {
	tx, err := r.db.BeginTransaction(ctx)
	if err != nil {
		fmt.Printf("ini error 0: %s ", err)
		return nil, err
	}

	// Build the SQL query with placeholders
	query := "INSERT INTO users ( username, name, password) VALUES ($1, $2, $3)"

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		_ = tx.Rollback() // rollback on prepare error

		fmt.Printf("ini error 1: %s ", err)

		return nil, err
	}
	defer stmt.Close() // ensure statement is closed

	// Execute the query with the transaction and scanned arguments
	_, err = stmt.ExecContext(ctx, user.Username, user.Name, user.Password)
	if err != nil {
		_ = tx.Rollback() // rollback on insert error
		fmt.Printf("ini error 2: %s ", err)

		// Check if the error is a unique constraint violation
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return nil, entity.ErrDataAlreadyExists
			}
		}
		return nil, err
	}

	// Commit the transaction if successful
	if err := tx.Commit(); err != nil {
		fmt.Printf("ini error 3: %s ", err)
		return nil, err
	}

	return &user, nil

}

func (r *Repository) LoginUser(ctx context.Context, username, password string) (*entity.User, error) {

	var user *entity.User

	// Build the SQL query with placeholders
	query := "SELECT * FROM users WHERE username = $1"

	// Prepare the statement with the context
	stmt, err := r.db.GetInstance(orm.WorkloadRead).PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close() // ensure statement is closed

	// Execute the query with scanned arguments
	row := stmt.QueryRowContext(ctx, username)

	user = &entity.User{}

	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Name)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, entity.ErrDataNotFound
		}
		return nil, err
	}

	isPasswordTrue := bcrypt.CheckPassword(user.Password, password)

	if isPasswordTrue {
		fmt.Printf("ini name 1: %s ", user.Name)
		return user, nil
	} else {
		return nil, entity.ErrInvalidPassword
	}

}
