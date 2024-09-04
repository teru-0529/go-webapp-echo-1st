/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"
	"fmt"

	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
)

// STRUCT:
type ReceivingPostCommand struct {
	body        spec.ReceivingPostBody
	OrderNo     spec.OrderNo
	orderRepo   iReceivingRepository
	productRepo iProductRepository
}

// FUNCTION:
func NewReceivingPostCommand(body spec.ReceivingPostBody, orderRepo iReceivingRepository, productRepo iProductRepository) *ReceivingPostCommand {
	return &ReceivingPostCommand{body: body, orderRepo: orderRepo, productRepo: productRepo}
}

// FUNCTION:
func (cmd *ReceivingPostCommand) Execute(ctx context.Context, tx *sql.Tx) error {

	// PROCESS:
	// 存在チェック(商品)
	for i, detail := range cmd.body.Details {
		exist, err := cmd.productRepo.Exists(ctx, tx, detail.ProductId)
		if err != nil {
			return err
		} else if !exist {
			return fmt.Errorf("details: (%d: (Product: %w).).", i, ErrNotFound)
		}
	}

	// PROCESS:
	// 構造体
	record := ordersdb.Receiving{
		OperatorName: cmd.body.OperatorName,
		CustomerName: cmd.body.CustomerName,
		CreatedBy:    traceId(ctx),
		UpdatedBy:    traceId(ctx),
	}
	recordDetails := []ordersdb.ReceivingDetail{}
	for _, detail := range cmd.body.Details {
		recordDetails = append(recordDetails, ordersdb.ReceivingDetail{
			ProductID:         detail.ProductId,
			ReceivingQuantity: detail.OrderQuantity,
			SelllingPrice:     detail.SellingPrice,
			CreatedBy:         traceId(ctx),
			UpdatedBy:         traceId(ctx),
		})
	}

	// PROCESS:
	// 登録
	orderNo, err := cmd.orderRepo.Save(ctx, tx, record, recordDetails)
	if err != nil {
		return err
	}

	cmd.OrderNo = orderNo
	return nil
}
