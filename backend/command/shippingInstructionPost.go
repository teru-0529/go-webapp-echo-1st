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
)

// STRUCT:
type ShippingIsntructionPostCommand struct {
	body      spec.ShippingInstructionBody
	OrderNo   spec.OrderNo
	orderRepo iReceivingRepository
	instRepo  iShippingInstructionRepository
}

// FUNCTION:
func NewShippingIsntructionPostCommand(body spec.ShippingInstructionBody, orderRepo iReceivingRepository, instRepo iShippingInstructionRepository) *ShippingIsntructionPostCommand {
	return &ShippingIsntructionPostCommand{body: body, orderRepo: orderRepo, instRepo: instRepo}
}

// FUNCTION:
func (cmd *ShippingIsntructionPostCommand) Execute(ctx context.Context, tx *sql.Tx) error {

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
	record := ordersdb.ShippingInstruction{
		OrderNo:          cmd.body.OrderNo,
		ProductID:        cmd.body.ProductId,
		OperatorName:     cmd.body.OperatorName,
		ShippingQuantity: cmd.body.Quantity,
		CreatedBy:        traceId(ctx),
		UpdatedBy:        traceId(ctx),
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
