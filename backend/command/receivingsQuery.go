/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"

	openapi_types "github.com/oapi-codegen/runtime/types"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// STRUCT:
type ReceivingQueryParam struct {
	qb           QueryBase
	customerName *spec.CustomerName
	orderStatus  *spec.OrderStatus
}

// FUNCTION:
func NewReceivingQueryParam(params spec.OrdersReceivingsGetParams) ReceivingQueryParam {
	return ReceivingQueryParam{
		qb:           NewQueryBase(params.Limit, params.Offset),
		customerName: params.CustomerName,
		orderStatus:  params.OrderStatus,
	}
}

// FUNCTION:
func (qp *ReceivingQueryParam) Qm() []qm.QueryMod {

	// PROCESS:
	mods := []qm.QueryMod{}
	if qp.customerName != nil {
		mods = append(mods, ordersdb.ReceivingWhere.CustomerName.EQ(*qp.customerName))
	}
	if qp.orderStatus != nil {
		mods = append(mods, ordersdb.ReceivingWhere.OrderStatus.EQ(ordersdb.OrderStatus(*qp.orderStatus)))
	}
	mods = append(mods, qp.qb.Qm()...)
	return mods
}

// FUNCTION:
func (qp *ReceivingQueryParam) Limit() int {
	return qp.qb.limit
}

// STRUCT:
type ReceivingQueryCommand struct {
	qp          ReceivingQueryParam
	Response    spec.ReceivingArray
	IsRemaining bool
	orderRepo   iReceivingRepository
}

// FUNCTION:
func NewReceivingQueryCommand(params spec.OrdersReceivingsGetParams, orderRepo iReceivingRepository) *ReceivingQueryCommand {
	return &ReceivingQueryCommand{qp: NewReceivingQueryParam(params), orderRepo: orderRepo}
}

// FUNCTION:
func (cmd *ReceivingQueryCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 検索(受注)
	receivings, isRemaining, err := cmd.orderRepo.Query(ctx, tx, cmd.qp)
	if err != nil {
		return err
	}

	result := []spec.Receiving{}
	for _, receiving := range receivings {
		result = append(result, spec.Receiving{
			OrderNo:             receiving.OrderNo,
			OrderDate:           openapi_types.Date{Time: receiving.OrderDate},
			OperatorName:        receiving.OperatorName,
			CustomerName:        receiving.CustomerName,
			TotalOrderPrice:     receiving.TotalOrderPrice,
			RemainingOrderPrice: receiving.RemainingOrderPrice,
			OrderStatus:         spec.OrderStatus(receiving.OrderStatus),
		})
	}
	cmd.Response = result
	cmd.IsRemaining = isRemaining
	return nil
}
