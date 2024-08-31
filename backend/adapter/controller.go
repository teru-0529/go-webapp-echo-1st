/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package adapter

import (
	"fmt"
	"net/http"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/teru-0529/go-webapp-echo-1st/command"
	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// STRUCT:
type ApiController struct{}

// 受注登録
// (POST /receivings)
// FUNCTION:
func (ac ApiController) OrdersReceivingsPost(ctx echo.Context, params spec.OrdersReceivingsPostParams) error {
	traceId := params.XTaraceId
	receiving := spec.ReceivingPostBody{}
	if err := ctx.Bind(&receiving); err != nil {
		return err
	}
	// FIXME:repository

	// PROCESS: コマンド実行
	cmd := command.NewReceivingPostCommand(traceId, receiving)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	ctx.Response().Header().Set("location", path.Join("/receivings", cmd.OrderNo))
	return ctx.NoContent(http.StatusCreated)
}

// 受注一覧検索
// (GET /receivings)
// FUNCTION:
func (ac ApiController) OrdersReceivingsGet(ctx echo.Context, params spec.OrdersReceivingsGetParams) error {
	traceId := params.XTaraceId
	qb := command.NewQueryBase(params.Limit, params.Offset)
	qp := command.NewReceivingQueryParam(params)
	// FIXME:repository

	// PROCESS: コマンド実行
	cmd := command.NewReceivingQueryCommand(traceId, qb, qp)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	ctx.Response().Header().Set("is_remaining", fmt.Sprintf("%t", cmd.IsRemaining))
	return ctx.JSON(http.StatusOK, cmd.Response)
}

// 受注取得
// (GET /receivings/{order_no})
// FUNCTION:
func (ac ApiController) OrdersReceivingsNoGet(ctx echo.Context, orderNo spec.OrderNo, params spec.OrdersReceivingsNoGetParams) error {
	traceId := params.XTaraceId
	// FIXME:repository

	// PROCESS: コマンド実行
	cmd := command.NewReceivingGetCommand(traceId, orderNo)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, cmd.Response)
}

// 受注修正
// (PUT /receivings/{order_no}/operator)
// FUNCTION:
func (ac ApiController) OrdersReceivingsNoOperatorPut(ctx echo.Context, orderNo spec.OrderNo, params spec.OrdersReceivingsNoOperatorPutParams) error {
	traceId := params.XTaraceId
	receivingOperator := spec.ReceivingOperatorBody{}
	if err := ctx.Bind(&receivingOperator); err != nil {
		return err
	}
	// FIXME:repository

	// PROCESS: コマンド実行
	cmd := command.NewReceivingPutOperatorCommand(traceId, orderNo, receivingOperator)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	// FIXME:エラー処理サンプル
	// return echo.NewHTTPError(http.StatusNotFound, "orderDetail is not found.")

	return ctx.NoContent(http.StatusNoContent)
}

// キャンセル指示登録
// (POST /cancel-instructions)
// FUNCTION:
func (ac ApiController) OrdersCancelInstructionsPost(ctx echo.Context, params spec.OrdersCancelInstructionsPostParams) error {
	traceId := params.XTaraceId
	cancelInstruction := spec.CancelInstructionBody{}
	if err := ctx.Bind(&cancelInstruction); err != nil {
		return err
	}
	// FIXME:repository

	// PROCESS: コマンド実行
	cmd := command.NewCancelInstructionPostCommand(traceId, cancelInstruction)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	ctx.Response().Header().Set("location", path.Join("/receivings", cmd.OrderNo))
	return ctx.NoContent(http.StatusCreated)
}

// 出荷指示登録
// (POST /shipping-instructions)
// FUNCTION:
func (ac ApiController) OrdersShippingInstructionsPost(ctx echo.Context, params spec.OrdersShippingInstructionsPostParams) error {
	traceId := params.XTaraceId
	shippingInstruction := spec.ShippingInstructionBody{}
	if err := ctx.Bind(&shippingInstruction); err != nil {
		return err
	}
	// FIXME:repository

	// PROCESS: コマンド実行
	cmd := command.NewShippingIsntructionPostCommand(traceId, shippingInstruction)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	ctx.Response().Header().Set("location", path.Join("/receivings", cmd.OrderNo))
	return ctx.NoContent(http.StatusCreated)
}
