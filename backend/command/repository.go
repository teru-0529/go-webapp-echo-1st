/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"

	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
)

type iReceivingRepository interface {
	Save(context.Context, *sql.Tx, ordersdb.Receiving, []ordersdb.ReceivingDetail) (string, error)
	Query(context.Context, *sql.Tx, ReceivingQueryParam) (ordersdb.ReceivingSlice, bool, error)
	Get(context.Context, *sql.Tx, string) (*ordersdb.Receiving, error)
	UpdateOperator(context.Context, *sql.Tx, ordersdb.Receiving) error
	Exists(context.Context, *sql.Tx, string) (bool, error)
	DetailGet(context.Context, *sql.Tx, string, string) (*ordersdb.ReceivingDetail, error)
}

type iProductRepository interface {
	Exists(context.Context, *sql.Tx, string) (bool, error)
}

type iCancelInstructionRepository interface {
	Save(context.Context, *sql.Tx, ordersdb.CancelInstruction) (int, error)
}

type iShippingInstructionRepository interface {
	Save(context.Context, *sql.Tx, ordersdb.ShippingInstruction) (int, error)
}
