/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package repository

import (
	"context"
	"database/sql"

	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
)

// TITLE: リポジトリ(商品)

type ProductRepository struct{}

// FUNCTION: new
func NewProductRepo() *ProductRepository {
	return &ProductRepository{}
}

// FUNCTION: exists
func (r ProductRepository) Exists(ctx context.Context, tx *sql.Tx, productId string) (bool, error) {

	// PROCESS: Product Exists
	return ordersdb.Products(
		ordersdb.ProductWhere.ProductID.EQ(productId),
	).Exists(ctx, tx)
}
