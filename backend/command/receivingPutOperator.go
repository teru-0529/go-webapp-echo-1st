/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"
	"fmt"

	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
)

// STRUCT:
type ReceivingPutOperatorCommand struct {
	orderNo   spec.OrderNo
	body      spec.ReceivingOperatorBody
	orderRepo iReceivingRepository
}

// FUNCTION:
func NewReceivingPutOperatorCommand(orderNo spec.OrderNo, body spec.ReceivingOperatorBody, orderRepo iReceivingRepository) *ReceivingPutOperatorCommand {
	return &ReceivingPutOperatorCommand{orderNo: orderNo, body: body, orderRepo: orderRepo}
}

// FUNCTION:
func (cmd *ReceivingPutOperatorCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 存在チェック(受注)
	exist, err := cmd.orderRepo.Exists(ctx, tx, cmd.orderNo)
	if err != nil {
		return err
	} else if !exist {
		return fmt.Errorf("Receiving: %w", ErrNotFound)
	}

	// PROCESS:
	// 更新
	record := ordersdb.Receiving{
		OrderNo:      cmd.orderNo,
		OperatorName: cmd.body.OperatorName,
		UpdatedBy:    traceId(ctx),
	}
	return cmd.orderRepo.UpdateOperator(ctx, tx, record)
}
