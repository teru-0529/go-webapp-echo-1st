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
type ReceivingPostCommand struct {
	ctx     context.Context
	body    spec.ReceivingPostBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewReceivingPostCommand(ctx context.Context, body spec.ReceivingPostBody) ReceivingPostCommand {
	return ReceivingPostCommand{ctx: ctx, body: body}
}

// FUNCTION:
func (cmd *ReceivingPostCommand) Ececute() error {

	// PROCESS:
	// 存在チェック(商品)

	// PROCESS:
	// 登録

	// FIXME:
	fmt.Println(infra.TraceId(cmd.ctx))
	fmt.Println(cmd.body)
	cmd.OrderNo = "RO-0000058"
	// FIXME:

	return nil
}
