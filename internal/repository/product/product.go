package product

import (
	"context"
	"fmt"
	"strings"

	"github.com/burhanwakhid/shopifyx_backend/internal/delivery/rest/request"
	"github.com/burhanwakhid/shopifyx_backend/internal/entity"
	"github.com/burhanwakhid/shopifyx_backend/pkg/orm"
)

type ProductRepository struct {
	db *orm.Replication
}

func NewProductRepository(db *orm.Replication) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Product Management
func (r *ProductRepository) CreateProduct(ctx context.Context, product entity.Product, userId string) error {
	tx, err := r.db.BeginTransaction(ctx)
	if err != nil {
		fmt.Printf("ini error 0: %s ", err)
		return err
	}

	// Build the SQL query with placeholders
	query := "INSERT INTO product ( name, price, imageurl, stock, condition, tags, owner_id, ispurchaseable) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	// tagsString := strings.Join(product.Tags, ",")
	tagsString := "{'" + strings.Join(product.Tags, "', '") + "'}"

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		_ = tx.Rollback() // rollback on prepare error

		fmt.Printf("ini error 1: %s ", err)

		return err
	}
	defer stmt.Close() // ensure statement is closed

	// Execute the query with the transaction and scanned arguments
	_, err = stmt.ExecContext(ctx, product.Name, product.Price, product.ImageUrl, product.Stock, product.Condition, tagsString, userId, product.IsPurchasable)
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

func (r *ProductRepository) UpdateProduct(ctx context.Context, product entity.Product) (*entity.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, idProduct string) error {
	panic("not implemented") // TODO: Implement
}

// Product Page
func (r *ProductRepository) GetProduct(ctx context.Context, request request.Product, idOwner string) ([]*entity.Product, error) {
	// query := `
	// 	SELECT
	// 		p.id,
	// 		p.name,
	// 		p.price,
	// 		p.imageurl,
	// 		p.stock,
	// 		p.condition,
	// 		p.tags,
	// 		p.ispurchaseable,
	// 		COUNT(DISTINCT py.id) AS purchase_count
	// 	FROM
	// 		product p
	// 	LEFT JOIN
	// 		payment py ON py.product_id = p.id
	// 	WHERE
	// 		($1  IS NULL OR p.owner_id  = $1)  -- Filter by user (optional)
	// 		-- AND ($tags IS NULL OR string_to_array($tags, ',')::text[] <@ p.tags OR cardinality(string_to_array($tags, ',')) = 0)
	// 		AND ($2 IS NULL OR p.condition = $2)  -- Filter by condition (optional)
	// 		AND ($showEmptyStock IS NULL OR p.stock > 0 OR $showEmptyStock = TRUE)  -- Filter by stock (optional)
	// 		AND ($maxPrice IS NULL OR p.price <= $maxPrice)  -- Filter by max price (optional)
	// 		AND ($minPrice IS NULL OR p.price >= $minPrice)  -- Filter by min price (optional)
	// 		AND ($search IS NULL OR p.name ILIKE '%' || $search || '%')  -- Filter by search term (optional)
	// 	GROUP BY
	// 		p.id, p.name, p.price, p.imageurl, p.stock, p.condition, p.tags, p.ispurchaseable
	// 	ORDER BY
	// 		CASE WHEN $sortBy = 'price' THEN p.price END $orderBy,
	// 		CASE WHEN $sortBy = 'date' THEN p.id END $orderBy
	// 	LIMIT
	// 		$limit OFFSET $offset
	// `
	// stmt, err := r.db.GetInstance(orm.WorkloadRead).PrepareContext(ctx, query)
	// if err != nil {
	// 	fmt.Printf("ini error 0: %s ", err)
	// 	return nil, err
	// }
	// defer stmt.Close() // ensure statement is closed

	// rows, err := stmt.QueryContext(ctx, idUser)
	// if err != nil {
	// 	fmt.Printf("ini error 1: %s ", err)
	// 	return nil, err
	// }
	// defer rows.Close() // ensure rows are closed

	// var banks []*entity.Bank
	// for rows.Next() {
	// 	var b entity.Bank
	// 	err := rows.Scan(&b.Id, &b.Name, &b.AccountName, &b.AccountNumber, &b.IdUser)
	// 	if err != nil {
	// 		fmt.Printf("ini error 2: %s ", err)
	// 		if errors.Is(err, sql.ErrNoRows) {
	// 			return nil, entity.ErrDataNotFound
	// 		}
	// 		return nil, err
	// 	}
	// 	banks = append(banks, &b)
	// }

	// fmt.Printf("Result Length: %d", len(banks))

	// return banks, nil

	return nil, nil
}

func (r *ProductRepository) GetProductById(ctx context.Context, idProduct string) (*entity.ProductDetail, error) {
	panic("not implemented") // TODO: Implement
}
