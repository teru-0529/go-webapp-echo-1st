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
type ReceivingPostCommand struct {
	body    spec.ReceivingPostBody
	OrderNo spec.OrderNo
}

// FUNCTION:
func NewReceivingPostCommand(body spec.ReceivingPostBody) *ReceivingPostCommand {
	return &ReceivingPostCommand{body: body}
}

// FUNCTION:
func (cmd *ReceivingPostCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 存在チェック(商品)

	// PROCESS:
	// 登録

	// FIXME:
	fmt.Println(infra.TraceId(ctx))
	fmt.Println(cmd.body)
	cmd.OrderNo = "RO-0000058"
	// FIXME:

	return nil
}
