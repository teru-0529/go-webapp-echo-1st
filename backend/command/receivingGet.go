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
type ReceivingGetCommand struct {
	traceId  spec.TraceId
	orderNo  spec.OrderNo
	Response *spec.ReceivingWithDetail
}

// FUNCTION:
func NewReceivingGetCommand(traceId spec.TraceId, orderNo spec.OrderNo) ReceivingGetCommand {
	return ReceivingGetCommand{traceId: traceId, orderNo: orderNo}
}

// FUNCTION:
func (cmd *ReceivingGetCommand) Ececute() error {

	// PROCESS:
	// 取得(受注)

	// FIXME:
	fmt.Println(cmd.traceId)
	fmt.Println(cmd.orderNo)

	const layout = "2006-01-02"
	t, _ := time.Parse(layout, "2024-05-26")
	cmd.Response = &spec.ReceivingWithDetail{
		OrderNo:             "RO-0000056",
		OrderDate:           openapi_types.Date{Time: t},
		OperatorName:        "織田信長",
		CustomerName:        "徳川物産株式会社",
		TotalOrderPrice:     230200,
		RemainingOrderPrice: 111200,
		OrderStatus:         "WORK_IN_PROGRESS",
		Details:             []spec.ReceivingDetail{},
	}
	cmd.Response.Details = append(cmd.Response.Details, spec.ReceivingDetail{
		ProductId:         "P0001",
		OrderQuantity:     5,
		ShippingQuantity:  1,
		CancelQuantity:    0,
		RemainingQuantity: 4,
		SellingPrice:      27800,
		CostPrice:         19800,
		ProfitRate:        0.29,
		OrderStatus:       "WORK_IN_PROGRESS",
	})
	cmd.Response.Details = append(cmd.Response.Details, spec.ReceivingDetail{
		ProductId:         "P0005",
		OrderQuantity:     3,
		ShippingQuantity:  2,
		CancelQuantity:    1,
		RemainingQuantity: 0,
		SellingPrice:      45600,
		CostPrice:         28700,
		ProfitRate:        0.37,
		OrderStatus:       "COMPLETED",
	})
	cmd.Response.Details = append(cmd.Response.Details, spec.ReceivingDetail{
		ProductId:         "P0006",
		OrderQuantity:     1,
		ShippingQuantity:  0,
		CancelQuantity:    1,
		RemainingQuantity: 0,
		SellingPrice:      100200,
		CostPrice:         73800,
		ProfitRate:        0.26,
		OrderStatus:       "CANCELED",
	})
	// FIXME:

	return nil
}
