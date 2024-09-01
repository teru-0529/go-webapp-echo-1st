/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"fmt"

	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

// STRUCT:
type ReceivingPutOperatorCommand struct {
	accountId spec.AccountId
	orderNo   spec.OrderNo
	body      spec.ReceivingOperatorBody
}

// FUNCTION:
func NewReceivingPutOperatorCommand(accountId spec.AccountId, orderNo spec.OrderNo, body spec.ReceivingOperatorBody) ReceivingPutOperatorCommand {
	return ReceivingPutOperatorCommand{accountId: accountId, orderNo: orderNo, body: body}
}

// FUNCTION:
func (cmd *ReceivingPutOperatorCommand) Ececute() error {

	// PROCESS:
	// 取得(受注)

	// PROCESS:
	// 更新

	// FIXME:
	fmt.Println(cmd.accountId)
	fmt.Println(cmd.orderNo)
	fmt.Println(cmd.body)
	// FIXME:

	return nil
}
