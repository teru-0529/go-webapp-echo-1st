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

// TITLE: リポジトリ(キャンセル指示)

type CancelInstructionRepository struct{}

// FUNCTION: new
func NewCancelRepo() *CancelInstructionRepository {
	return &CancelInstructionRepository{}
}

// FUNCTION: save
func (r CancelInstructionRepository) Save(
	ctx context.Context, tx *sql.Tx, record ordersdb.CancelInstruction,
) (int, error) {

	// PROCESS: CancelInstruction Insert
	cols := NewRegistrationCols()
	cols.add(ordersdb.CancelInstructionColumns.OrderNo)
	cols.add(ordersdb.CancelInstructionColumns.ProductID)
	cols.add(ordersdb.CancelInstructionColumns.OperatorName)
	cols.add(ordersdb.CancelInstructionColumns.CancelQuantity)
	if record.CancelReason.Valid {
		cols.add(ordersdb.CancelInstructionColumns.CancelReason)
	}
	err := record.Insert(ctx, tx, boil.Whitelist(cols.InsertCols...))
	if err != nil {
		return -1, err
	}

	return record.CancelNo, nil
}
