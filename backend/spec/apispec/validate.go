/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package apispec

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Limit 検索結果の取得数上限値
var LimitRule = []validation.Rule{
	validation.Min(0),
	validation.Max(1000),
}

// Offset 検索実行時のオフセット値
var OffsetRule = []validation.Rule{
	validation.Min(1),
}

// OrderStatus 受注状況(Query)
var OrderStatusQueryRule = []validation.Rule{
	validation.In(CANCELED, COMPLETED, PREPARING, WORKINPROGRESS).Error("must be in a valid enum type"),
}

// CustomerName 受注先企業名(Query)
var CustomerNameQueryRule = []validation.Rule{
	validation.Length(1, 50),
}

// OrderNo 受注番号
var OrderNoRule = []validation.Rule{
	validation.Required,
	validation.Match(regexp.MustCompile("^RO-[0-9]{7}$")).Error("must be in a valid format[RO-0000000]"),
}

// ProductId 商品ID
var ProductIdRule = []validation.Rule{
	validation.Required,
	validation.Match(regexp.MustCompile("^P[0-9]{4}$")).Error("must be in a valid format[P0000]"),
}

// OperatorName 担当者名称
var OperatorNameRule = []validation.Rule{
	validation.Required,
	validation.Length(3, 30),
}

// CustomerName 受注先企業名
var CustomerNameRule = []validation.Rule{
	validation.Required,
	validation.Length(1, 50),
}

// Quantity 数量
var QuantityRule = []validation.Rule{
	validation.Required.Error("must be no less than 1"),
	validation.Min(1),
	validation.Max(1000),
}

// SellingPrice 金額
var PriceRule = []validation.Rule{
	validation.Required.Error("must be no less than 1"),
	validation.Min(1),
	validation.Max(9999999),
}

// Details (受注登録子要素)
var ReceivingPostDetailsRule = []validation.Rule{
	validation.Required.Error("The number of items must be at least 1"),
}

// FUNCTION:
func (s OrdersReceivingsGetParams) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Limit, LimitRule...),
		validation.Field(&s.Offset, OffsetRule...),
		validation.Field(&s.CustomerName, CustomerNameQueryRule...),
		validation.Field(&s.OrderStatus, OrderStatusQueryRule...),
	)
}

// FUNCTION:
func (s ReceivingPostBody) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.OperatorName, OperatorNameRule...),
		validation.Field(&s.CustomerName, CustomerNameRule...),
		validation.Field(&s.Details, ReceivingPostDetailsRule...),
	)
}

// FUNCTION:
func (s ReceivingPostDetail) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ProductId, ProductIdRule...),
		validation.Field(&s.OrderQuantity, QuantityRule...),
		validation.Field(&s.SellingPrice, PriceRule...),
	)
}

// FUNCTION:
func (s ReceivingOperatorBody) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.OperatorName, OperatorNameRule...),
	)
}

// FUNCTION:
func (s CancelInstructionBody) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.OperatorName, OperatorNameRule...),
		validation.Field(&s.OrderNo, OrderNoRule...),
		validation.Field(&s.ProductId, ProductIdRule...),
		validation.Field(&s.Quantity, QuantityRule...),
	)
}

// FUNCTION:
func (s ShippingInstructionBody) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.OperatorName, OperatorNameRule...),
		validation.Field(&s.OrderNo, OrderNoRule...),
		validation.Field(&s.ProductId, ProductIdRule...),
		validation.Field(&s.Quantity, QuantityRule...),
	)
}
