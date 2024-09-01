/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"fmt"

	"github.com/teru-0529/go-webapp-echo-1st/infra"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// STRUCT:
type CancelInstructionPostCommand struct {
	ctx     context.Context
	body    spec.CancelInstructionBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewCancelInstructionPostCommand(ctx context.Context, body spec.CancelInstructionBody) CancelInstructionPostCommand {
	return CancelInstructionPostCommand{ctx: ctx, body: body}
}

// FUNCTION:
func (cmd *CancelInstructionPostCommand) Ececute() error {

	// PROCESS:
	// 存在チェック(受注明細)

	// PROCESS:
	// 登録

	// FIXME:
	// DB接続確認
	record := &ordersdb.Product{
		ProductName: "日本刀",
		CostPrice:   20000,
		CreatedBy:   traceId(cmd.ctx),
		UpdatedBy:   traceId(cmd.ctx),
	}
	cols := NewRegistrationCols()
	cols.add(ordersdb.ProductColumns.ProductName)
	cols.add(ordersdb.ProductColumns.CostPrice)
	err := record.InsertG(
		cmd.ctx,
		boil.Whitelist(cols.InsertCols...),
	)
	fmt.Println(err)

	// FIXME:
	fmt.Println(infra.TraceId(cmd.ctx))
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
