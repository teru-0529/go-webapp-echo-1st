/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package repository

import (
	"context"
	"database/sql"

	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// TITLE: リポジトリ(出荷指示)

type ShippingInstructionRepository struct{}

// FUNCTION: new
func NewShippingRepo() *ShippingInstructionRepository {
	return &ShippingInstructionRepository{}
}

// FUNCTION: save
func (r ShippingInstructionRepository) Save(
	ctx context.Context, tx *sql.Tx, record ordersdb.ShippingInstruction,
) (int, error) {

	// PROCESS: ShippingInstruction Insert
	cols := NewRegistrationCols()
	cols.add(ordersdb.ShippingInstructionColumns.OrderNo)
	cols.add(ordersdb.ShippingInstructionColumns.ProductID)
	cols.add(ordersdb.ShippingInstructionColumns.OperatorName)
	cols.add(ordersdb.ShippingInstructionColumns.ShippingQuantity)
	err := record.Insert(ctx, tx, boil.Whitelist(cols.InsertCols...))
	if err != nil {
		return -1, err
	}

	return record.SippingNo, nil
}
