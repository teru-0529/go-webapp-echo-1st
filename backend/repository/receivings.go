/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package repository

import (
	"context"
	"database/sql"

	"github.com/teru-0529/go-webapp-echo-1st/command"
	"github.com/teru-0529/go-webapp-echo-1st/spec/dbspec/ordersdb"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// TITLE: リポジトリ(受注)

type ReceivingRepository struct{}

// FUNCTION: new
func NewReceivingRepo() *ReceivingRepository {
	return &ReceivingRepository{}
}

// FUNCTION: save
func (r ReceivingRepository) Save(
	ctx context.Context, tx *sql.Tx, record ordersdb.Receiving, recordDetails []ordersdb.ReceivingDetail,
) (string, error) {

	// PROCESS: Receiving Insert
	cols := NewRegistrationCols()
	cols.add(ordersdb.ReceivingColumns.OperatorName)
	cols.add(ordersdb.ReceivingColumns.CustomerName)
	err := record.Insert(ctx, tx, boil.Whitelist(cols.InsertCols...))
	if err != nil {
		return "", err
	}

	// PROCESS: RECEIVING(OrderNo) Get ※OrderNoが登録時のフック処理で計算されるためTraceIdをキーに再検索して取得している。
	rec, err := ordersdb.Receivings(
		ordersdb.ReceivingWhere.CreatedBy.EQ(record.CreatedBy),
	).One(ctx, tx)
	if err != nil {
		return "", err
	}
	orderNo := rec.OrderNo

	// PROCESS: ReceivingDetail Insert
	dCols := NewRegistrationCols()
	dCols.add(ordersdb.ReceivingDetailColumns.OrderNo)
	dCols.add(ordersdb.ReceivingDetailColumns.ProductID)
	dCols.add(ordersdb.ReceivingDetailColumns.ReceivingQuantity)
	dCols.add(ordersdb.ReceivingDetailColumns.SelllingPrice)
	for _, detail := range recordDetails {
		detail.OrderNo = orderNo
		err := detail.Insert(ctx, tx, boil.Whitelist(dCols.InsertCols...))
		if err != nil {
			return "", err
		}
	}

	return orderNo, nil
}

// // FUNCTION: query
func (r ReceivingRepository) Query(
	ctx context.Context, tx *sql.Tx, qp command.ReceivingQueryParam) (ordersdb.ReceivingSlice, bool, error) {

	// PROCESS: Receiving Query
	results, err := ordersdb.Receivings(qp.Qm()...).All(ctx, tx)
	if err != nil {
		return nil, false, err
	}

	// PROCESS: 検索残があるかどうかのチェック(実際のクエリ:resultsではLimit+1で絞り込んでいる)
	isRemaining := len(results) > qp.Limit()
	if isRemaining {
		results = results[:qp.Limit()]
	}
	return results, isRemaining, nil
}

// FUNCTION: get
func (r ReceivingRepository) Get(ctx context.Context, tx *sql.Tx, orderNo string) (*ordersdb.Receiving, error) {

	// PROCESS: Receiving with detail Get
	return ordersdb.Receivings(
		ordersdb.ReceivingWhere.OrderNo.EQ(orderNo),
		qm.Load(ordersdb.ReceivingRels.OrderNoReceivingDetails),
	).One(ctx, tx)
}

// FUNCTION: update
func (r ReceivingRepository) UpdateOperator(ctx context.Context, tx *sql.Tx, record ordersdb.Receiving) error {

	// PROCESS: Receiving Update
	cols := NewRegistrationCols()
	cols.add(ordersdb.ReceivingColumns.OperatorName)
	_, err := record.Update(ctx, tx, boil.Whitelist(cols.UpdateCols...))
	return err
}

// FUNCTION: exists
func (r ReceivingRepository) Exists(ctx context.Context, tx *sql.Tx, orderNo string) (bool, error) {

	// PROCESS: Receiving Exists
	return ordersdb.Receivings(
		ordersdb.ReceivingWhere.OrderNo.EQ(orderNo),
	).Exists(ctx, tx)
}

// FUNCTION: get
func (r ReceivingRepository) DetailGet(ctx context.Context, tx *sql.Tx, orderNo string, productId string) (*ordersdb.ReceivingDetail, error) {

	// PROCESS: ReceivingDetail Get
	return ordersdb.ReceivingDetails(
		ordersdb.ReceivingDetailWhere.OrderNo.EQ(orderNo),
		ordersdb.ReceivingDetailWhere.ProductID.EQ(productId),
	).One(ctx, tx)
}
