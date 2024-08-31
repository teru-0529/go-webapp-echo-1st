/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"fmt"

	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// STRUCT:
type ShippingIsntructionPostCommand struct {
	traceId spec.TraceId
	body    spec.ShippingInstructionBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewShippingIsntructionPostCommand(traceId spec.TraceId, body spec.ShippingInstructionBody) ShippingIsntructionPostCommand {
	return ShippingIsntructionPostCommand{traceId: traceId, body: body}
}

// FUNCTION:
func (cmd *ShippingIsntructionPostCommand) Ececute() error {

	// FIXME:
	fmt.Println(cmd.traceId)
	fmt.Println(cmd.body)
	cmd.OrderNo = cmd.body.OrderNo
	// FIXME:

	return nil
}
