/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

// STRUCT:
type ReceivingGetCommand struct {
	orderNo   spec.OrderNo
	Response  *spec.ReceivingWithDetail
	orderRepo iReceivingRepository
}

// FUNCTION:
func NewReceivingGetCommand(orderNo spec.OrderNo, orderRepo iReceivingRepository) *ReceivingGetCommand {
	return &ReceivingGetCommand{orderNo: orderNo, orderRepo: orderRepo}
}

// FUNCTION:
func (cmd *ReceivingGetCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 取得(受注)
	receiving, err := cmd.orderRepo.Get(ctx, tx, cmd.orderNo)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("Receiving: %w", ErrNotFound)
	} else if err != nil {
		return err
	}

	result := spec.ReceivingWithDetail{
		OrderNo:             receiving.OrderNo,
		OrderDate:           openapi_types.Date{Time: receiving.OrderDate},
		OperatorName:        receiving.OperatorName,
		CustomerName:        receiving.CustomerName,
		TotalOrderPrice:     receiving.TotalOrderPrice,
		RemainingOrderPrice: receiving.RemainingOrderPrice,
		OrderStatus:         spec.OrderStatus(receiving.OrderStatus),
		Details:             []spec.ReceivingDetail{},
	}
	for _, detail := range receiving.R.OrderNoReceivingDetails {
		profitRate, _ := detail.ProfitRate.Float64()
		fmt.Println(profitRate)
		result.Details = append(result.Details, spec.ReceivingDetail{
			ProductId:         detail.ProductID,
			OrderQuantity:     detail.ReceivingQuantity,
			ShippingQuantity:  detail.ShippingQuantity,
			CancelQuantity:    detail.CancelQuantity,
			RemainingQuantity: detail.RemainingQuantity,
			SellingPrice:      detail.SelllingPrice,
			CostPrice:         detail.CostPrice,
			ProfitRate:        float32(profitRate),
			OrderStatus:       spec.OrderStatus(detail.OrderStatus),
		})
	}
	cmd.Response = &result
	return nil
}
