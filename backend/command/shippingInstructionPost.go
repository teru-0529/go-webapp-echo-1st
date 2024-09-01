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
type ShippingIsntructionPostCommand struct {
	ctx     context.Context
	body    spec.ShippingInstructionBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewShippingIsntructionPostCommand(ctx context.Context, body spec.ShippingInstructionBody) ShippingIsntructionPostCommand {
	return ShippingIsntructionPostCommand{ctx: ctx, body: body}
}

// FUNCTION:
func (cmd *ShippingIsntructionPostCommand) Ececute() error {

	// PROCESS:
	// 存在チェック(受注明細)

	// PROCESS:
	// 登録

	// FIXME:
	fmt.Println(infra.TraceId(cmd.ctx))
	fmt.Println(cmd.body)
	cmd.OrderNo = cmd.body.OrderNo
	// FIXME:

	return nil
}
