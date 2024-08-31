/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"fmt"

	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// STRUCT:
type ReceivingPostCommand struct {
	traceId spec.TraceId
	body    spec.ReceivingPostBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewReceivingPostCommand(traceId spec.TraceId, body spec.ReceivingPostBody) ReceivingPostCommand {
	return ReceivingPostCommand{traceId: traceId, body: body}
}

// FUNCTION:
func (cmd *ReceivingPostCommand) Ececute() error {

	// PROCESS:
	// 存在チェック(商品)

	// PROCESS:
	// 登録

	// FIXME:
	fmt.Println(cmd.traceId)
	fmt.Println(cmd.body)
	cmd.OrderNo = "RO-0000058"
	// FIXME:

	return nil
}
