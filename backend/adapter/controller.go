/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package adapter

import (
	"fmt"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// TITLE:
type ApiController struct{}

// 受注登録
// (POST /receivings)
// FUNCTION:
func (ac ApiController) OrdersReceivingsPost(ctx echo.Context) error {
	traceId := ctx.Request().Header.Get("x-tarace-id")
	receiving := spec.ReceivingPostBody{}
	if err := ctx.Bind(&receiving); err != nil {
		return err
	}

	// FIXME:
	fmt.Println(traceId)
	fmt.Println(receiving)
	orderNo := "RO-0000024"
	// FIXME:

	ctx.Response().Header().Set("location", path.Join("/receivings", orderNo))
	return ctx.NoContent(http.StatusCreated)
}

// 受注一覧検索
// (GET /receivings)
// FUNCTION:
func (ac ApiController) OrdersReceivingsGet(ctx echo.Context, params spec.OrdersReceivingsGetParams) error {
	traceId := ctx.Request().Header.Get("x-tarace-id")

	// FIXME:
	fmt.Println(traceId)
	fmt.Println(*params.CustomerName)
	fmt.Println(*params.OrderStatus)
	// FIXME:

	return ctx.NoContent(http.StatusOK)
}

// 受注取得
// (GET /receivings/{order_no})
// FUNCTION:
func (ac ApiController) OrdersReceivingsNoGet(ctx echo.Context, orderNo spec.OrderNo) error {
	traceId := ctx.Request().Header.Get("x-tarace-id")

	// FIXME:
	fmt.Println(traceId)
	fmt.Println(orderNo)
	// FIXME:

	return ctx.NoContent(http.StatusOK)
}

// 受注修正
// (PATCH /receivings/{order_no})
// FUNCTION:
func (ac ApiController) OrdersReceivingsNoPatch(ctx echo.Context, orderNo spec.OrderNo) error {
	traceId := ctx.Request().Header.Get("x-tarace-id")
	receivingPatch := spec.ReceivingPatchBody{}
	if err := ctx.Bind(&receivingPatch); err != nil {
		return err
	}

	// FIXME:
	fmt.Println(traceId)
	fmt.Println(orderNo)
	fmt.Println(receivingPatch)
	// FIXME:

	return ctx.NoContent(http.StatusOK)
}

// キャンセル指示登録
// (POST /cancel-instructions)
// FUNCTION:
func (ac ApiController) OrdersCancelInstructionsPost(ctx echo.Context) error {
	traceId := ctx.Request().Header.Get("x-tarace-id")
	cancelInstruction := spec.CancelInstructionBody{}
	if err := ctx.Bind(&cancelInstruction); err != nil {
		return err
	}

	// FIXME:
	fmt.Println(traceId)
	fmt.Println(cancelInstruction)
	// FIXME:

	ctx.Response().Header().Set("location", path.Join("/receivings", cancelInstruction.OrderNo))
	return ctx.NoContent(http.StatusCreated)
}

// 出荷指示登録
// (POST /shipping-instructions)
// FUNCTION:
func (ac ApiController) OrdersShippingInstructionsPost(ctx echo.Context) error {
	traceId := ctx.Request().Header.Get("x-tarace-id")
	shippingInstruction := spec.ShippingInstructionBody{}
	if err := ctx.Bind(&shippingInstruction); err != nil {
		return err
	}

	// FIXME:
	fmt.Println(traceId)
	fmt.Println(shippingInstruction)
	// FIXME:

	ctx.Response().Header().Set("location", path.Join("/receivings", shippingInstruction.OrderNo))
	return ctx.NoContent(http.StatusCreated)
}
