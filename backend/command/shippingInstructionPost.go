/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"fmt"

	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

// STRUCT:
type ShippingIsntructionPostCommand struct {
	accountId spec.AccountId
	body      spec.ShippingInstructionBody
	OrderNo   spec.OrderNo
}

// FUNCTION:
func NewShippingIsntructionPostCommand(accountId spec.AccountId, body spec.ShippingInstructionBody) ShippingIsntructionPostCommand {
	return ShippingIsntructionPostCommand{accountId: accountId, body: body}
}

// FUNCTION:
func (cmd *ShippingIsntructionPostCommand) Ececute() error {

	// PROCESS:
	// 存在チェック(受注明細)

	// PROCESS:
	// 登録

	// FIXME:
	fmt.Println(cmd.accountId)
	fmt.Println(cmd.body)
	cmd.OrderNo = cmd.body.OrderNo
	// FIXME:

	return nil
}
