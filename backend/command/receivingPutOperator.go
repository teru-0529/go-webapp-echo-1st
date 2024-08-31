/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"fmt"

	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// STRUCT:
type ReceivingPutOperatorCommand struct {
	traceId spec.TraceId
	orderNo spec.OrderNo
	body    spec.ReceivingOperatorBody
}

// FUNCTION:
func NewReceivingPutOperatorCommand(traceId spec.TraceId, orderNo spec.OrderNo, body spec.ReceivingOperatorBody) ReceivingPutOperatorCommand {
	return ReceivingPutOperatorCommand{traceId: traceId, orderNo: orderNo, body: body}
}

// FUNCTION:
func (cmd *ReceivingPutOperatorCommand) Ececute() error {

	// FIXME:
	fmt.Println(cmd.traceId)
	fmt.Println(cmd.orderNo)
	fmt.Println(cmd.body)
	// FIXME:

	return nil
}
