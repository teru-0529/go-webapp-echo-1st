/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"fmt"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

// STRUCT:
type ReceivingQueryParam struct {
	customerName *spec.CustomerName
	orderStatus  *spec.OrderStatus
}

// FUNCTION:
func NewReceivingQueryParam(params spec.OrdersReceivingsGetParams) ReceivingQueryParam {
	return ReceivingQueryParam{
		customerName: params.CustomerName,
		orderStatus:  params.OrderStatus,
	}
}

// STRUCT:
type ReceivingQueryCommand struct {
	accountId   spec.AccountId
	queryBase   QueryBase
	queryParam  ReceivingQueryParam
	Response    spec.ReceivingArray
	IsRemaining bool
}

// FUNCTION:
func NewReceivingQueryCommand(accountId spec.AccountId, queryBase QueryBase, queryParam ReceivingQueryParam) ReceivingQueryCommand {
	return ReceivingQueryCommand{accountId: accountId, queryBase: queryBase, queryParam: queryParam}
}

// FUNCTION:
func (cmd *ReceivingQueryCommand) Ececute() error {

	// PROCESS:
	// 取得(受注)

	// FIXME:
	fmt.Println(cmd.accountId)
	fmt.Println(cmd.queryBase)
	fmt.Println(*cmd.queryParam.customerName)
	fmt.Println(*cmd.queryParam.orderStatus)

	const layout = "2006-01-02"
	cmd.Response = []spec.Receiving{}
	t, _ := time.Parse(layout, "2024-01-01")
	cmd.Response = append(cmd.Response, spec.Receiving{
		OrderNo:             "RO-0000001",
		OrderDate:           openapi_types.Date{Time: t},
		OperatorName:        "織田信長",
		CustomerName:        "徳川物産株式会社",
		TotalOrderPrice:     280000,
		RemainingOrderPrice: 280000,
		OrderStatus:         "COMPLETED",
	})
	t, _ = time.Parse(layout, "2024-03-14")
	cmd.Response = append(cmd.Response, spec.Receiving{
		OrderNo:             "RO-0000002",
		OrderDate:           openapi_types.Date{Time: t},
		OperatorName:        "織田信長",
		CustomerName:        "株式会社島津製作所",
		TotalOrderPrice:     0,
		RemainingOrderPrice: 0,
		OrderStatus:         "CANCELED",
	})
	t, _ = time.Parse(layout, "2024-04-26")
	cmd.Response = append(cmd.Response, spec.Receiving{
		OrderNo:             "RO-0000003",
		OrderDate:           openapi_types.Date{Time: t},
		OperatorName:        "上杉謙信",
		CustomerName:        "徳川物産株式会社",
		TotalOrderPrice:     145000,
		RemainingOrderPrice: 34000,
		OrderStatus:         "WORK_IN_PROGRESS",
	})
	cmd.IsRemaining = true
	// FIXME:

	return nil
}
