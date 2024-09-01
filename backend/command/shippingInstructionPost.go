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
type ShippingIsntructionPostCommand struct {
	body    spec.ShippingInstructionBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewShippingIsntructionPostCommand(body spec.ShippingInstructionBody) *ShippingIsntructionPostCommand {
	return &ShippingIsntructionPostCommand{body: body}
}

// FUNCTION:
func (cmd *ShippingIsntructionPostCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 存在チェック(受注明細)

	// PROCESS:
	// 登録

	// FIXME:
	fmt.Println(infra.TraceId(ctx))
	fmt.Println(cmd.body)
	cmd.OrderNo = cmd.body.OrderNo
	// return errors.New("XXX")
	// FIXME:

	return nil
}
