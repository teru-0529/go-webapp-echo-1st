/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
	"github.com/volatiletech/null/v8"
)

// STRUCT:
type CancelInstructionPostCommand struct {
	body      spec.CancelInstructionBody
	OrderNo   spec.OrderNo
	orderRepo iReceivingRepository
	instRepo  iCancelInstructionRepository
}

// FUNCTION:
func NewCancelInstructionPostCommand(body spec.CancelInstructionBody, orderRepo iReceivingRepository, instRepo iCancelInstructionRepository) *CancelInstructionPostCommand {
	return &CancelInstructionPostCommand{body: body, orderRepo: orderRepo, instRepo: instRepo}
}

// FUNCTION:
func (cmd *CancelInstructionPostCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 存在チェック(受注明細)
	detail, err := cmd.orderRepo.DetailGet(ctx, tx, cmd.body.OrderNo, cmd.body.ProductId)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("ReceivingDetail: %w", ErrNotFound)
	} else if err != nil {
		return err
	}
	if detail.RemainingQuantity < cmd.body.Quantity {
		return fmt.Errorf("quantity: %w", ErrOverflow)
	}

	// PROCESS:
	// 構造体
	record := ordersdb.CancelInstruction{
		OrderNo:        cmd.body.OrderNo,
		ProductID:      cmd.body.ProductId,
		OperatorName:   cmd.body.OperatorName,
		CancelQuantity: cmd.body.Quantity,
		CancelReason:   null.StringFromPtr(cmd.body.Reason),
		CreatedBy:      traceId(ctx),
		UpdatedBy:      traceId(ctx),
	}

	// PROCESS:
	// 登録
	_, err = cmd.instRepo.Save(ctx, tx, record)
	if err != nil {
		return err
	}

	cmd.OrderNo = cmd.body.OrderNo
	return nil
}
