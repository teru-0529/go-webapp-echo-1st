/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"fmt"

	"github.com/teru-0529/go-webapp-echo-1st/infra"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

// STRUCT:
type ReceivingPutOperatorCommand struct {
	ctx     context.Context
	orderNo spec.OrderNo
	body    spec.ReceivingOperatorBody
}

// FUNCTION:
func NewReceivingPutOperatorCommand(ctx context.Context, orderNo spec.OrderNo, body spec.ReceivingOperatorBody) ReceivingPutOperatorCommand {
	return ReceivingPutOperatorCommand{ctx: ctx, orderNo: orderNo, body: body}
}

// FUNCTION:
func (cmd *ReceivingPutOperatorCommand) Ececute(ctx context.Context) error {

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
