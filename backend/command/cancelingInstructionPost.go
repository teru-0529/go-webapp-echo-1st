/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/teru-0529/go-webapp-echo-1st/infra"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// STRUCT:
type CancelInstructionPostCommand struct {
	body    spec.CancelInstructionBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewCancelInstructionPostCommand(body spec.CancelInstructionBody) *CancelInstructionPostCommand {
	return &CancelInstructionPostCommand{body: body}
}

// FUNCTION:
func (cmd *CancelInstructionPostCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 存在チェック(受注明細)

	// PROCESS:
	// 登録

	// FIXME:
	// DB接続確認
	record := &ordersdb.Product{
		ProductName: "日本刀",
		CostPrice:   20000,
		CreatedBy:   traceId(ctx),
		UpdatedBy:   traceId(ctx),
	}
	cols := NewRegistrationCols()
	cols.add(ordersdb.ProductColumns.ProductName)
	cols.add(ordersdb.ProductColumns.CostPrice)
	err := record.Insert(
		ctx,
		tx,
		boil.Whitelist(cols.InsertCols...),
	)
	fmt.Println(err)

	// FIXME:
	fmt.Println(infra.TraceId(ctx))
	fmt.Println(cmd.body)
	cmd.OrderNo = cmd.body.OrderNo
	// FIXME:

	return nil
}

// TITLE: 変更対象フィールド

type RegistrationCols struct {
	InsertCols []string
	UpdateCols []string
}

// FUNCTION: new
func NewRegistrationCols() *RegistrationCols {
	return &RegistrationCols{
		InsertCols: []string{"created_by", "updated_by"},
		UpdateCols: []string{"updated_by"},
	}
}

// FUNCTION:
func (rc *RegistrationCols) add(col string) {
	rc.InsertCols = append(rc.InsertCols, col)
	rc.UpdateCols = append(rc.UpdateCols, col)
}

// FUNCTION: generateTraceId
func traceId(ctx context.Context) null.String {
	return null.StringFrom(infra.TraceId(ctx))
}
