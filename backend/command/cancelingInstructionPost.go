/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"fmt"

	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// STRUCT:
type CancelInstructionPostCommand struct {
	traceId spec.TraceId
	body    spec.CancelInstructionBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewCancelInstructionPostCommand(traceId spec.TraceId, body spec.CancelInstructionBody) CancelInstructionPostCommand {
	return CancelInstructionPostCommand{traceId: traceId, body: body}
}

// FUNCTION:
func (cmd *CancelInstructionPostCommand) Ececute() error {

	// FIXME:
	fmt.Println(cmd.traceId)
	fmt.Println(cmd.body)
	cmd.OrderNo = cmd.body.OrderNo
	// FIXME:

	return nil
}
