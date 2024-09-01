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
)

// STRUCT:
type ReceivingPutOperatorCommand struct {
	orderNo spec.OrderNo
	body    spec.ReceivingOperatorBody
}

// FUNCTION:
func NewReceivingPutOperatorCommand(orderNo spec.OrderNo, body spec.ReceivingOperatorBody) *ReceivingPutOperatorCommand {
	return &ReceivingPutOperatorCommand{orderNo: orderNo, body: body}
}

// FUNCTION:
func (cmd *ReceivingPutOperatorCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 取得(受注)

	// PROCESS:
	// 更新

	// FIXME:
	fmt.Println(infra.TraceId(ctx))
	fmt.Println(cmd.orderNo)
	fmt.Println(cmd.body)
	// FIXME:

	return nil
}
