/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package controller

import (
	"fmt"
	"net/http"
	"path"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/teru-0529/go-webapp-echo-1st/command"
	"github.com/teru-0529/go-webapp-echo-1st/infra"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

// STRUCT:
type ApiController struct{}

// 受注登録
// (POST /receivings)
// FUNCTION:
func (ac ApiController) OrdersReceivingsPost(ctx echo.Context, params spec.OrdersReceivingsPostParams) error {
	// PROCESS: アプリコンテキスト
	apCtx := infra.ConvertCtx(ctx, params.XAccountId)

	// PROCESS: Bodyパース/バリデーション
	receiving := spec.ReceivingPostBody{}
	if err := ctx.Bind(&receiving); err != nil {
		return err
	}
	if err := ctx.Validate(receiving); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			ctx.Logger().Error(k + ": " + err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// PROCESS: コマンド実行
	// FIXME:repository
	cmd := command.NewReceivingPostCommand(apCtx, receiving)
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
	// PROCESS: アプリコンテキスト
	apCtx := infra.ConvertCtx(ctx, params.XAccountId)

	// PROCESS: Queryバリデーション/整形
	if err := ctx.Validate(params); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			ctx.Logger().Error(k + ": " + err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	qb := command.NewQueryBase(params.Limit, params.Offset)
	qp := command.NewReceivingQueryParam(params)

	// PROCESS: コマンド実行
	// FIXME:repository
	cmd := command.NewReceivingQueryCommand(apCtx, qb, qp)
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
	// PROCESS: アプリコンテキスト
	apCtx := infra.ConvertCtx(ctx, params.XAccountId)

	// PROCESS: Pathバリデーション
	if err := validation.Validate(&orderNo, spec.OrderNoRule...); err != nil {
		message := "orderNo: " + err.Error()
		ctx.Logger().Error(message)
		return echo.NewHTTPError(http.StatusBadRequest, message)
	}

	// PROCESS: コマンド実行
	// FIXME:repository
	cmd := command.NewReceivingGetCommand(apCtx, orderNo)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, cmd.Response)
}

// 受注修正
// (PUT /receivings/{order_no}/operator)
// FUNCTION:
func (ac ApiController) OrdersReceivingsNoOperatorPut(ctx echo.Context, orderNo spec.OrderNo, params spec.OrdersReceivingsNoOperatorPutParams) error {
	// PROCESS: アプリコンテキスト
	apCtx := infra.ConvertCtx(ctx, params.XAccountId)

	// PROCESS: Pathバリデーション
	if err := validation.Validate(&orderNo, spec.OrderNoRule...); err != nil {
		message := "orderNo: " + err.Error()
		ctx.Logger().Error(message)
		return echo.NewHTTPError(http.StatusBadRequest, message)
	}

	// PROCESS: Bodyパース/バリデーション
	receivingOperator := spec.ReceivingOperatorBody{}
	if err := ctx.Bind(&receivingOperator); err != nil {
		return err
	}
	if err := ctx.Validate(receivingOperator); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			ctx.Logger().Error(k + ": " + err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// PROCESS: コマンド実行
	// FIXME:repository
	cmd := command.NewReceivingPutOperatorCommand(apCtx, orderNo, receivingOperator)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// キャンセル指示登録
// (POST /cancel-instructions)
// FUNCTION:
func (ac ApiController) OrdersCancelInstructionsPost(ctx echo.Context, params spec.OrdersCancelInstructionsPostParams) error {
	// PROCESS: アプリコンテキスト
	apCtx := infra.ConvertCtx(ctx, params.XAccountId)

	// PROCESS: Bodyパース/バリデーション
	cancelInstruction := spec.CancelInstructionBody{}
	if err := ctx.Bind(&cancelInstruction); err != nil {
		return err
	}
	if err := ctx.Validate(cancelInstruction); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			ctx.Logger().Error(k + ": " + err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// PROCESS: コマンド実行
	// FIXME:repository
	cmd := command.NewCancelInstructionPostCommand(apCtx, cancelInstruction)
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
	// PROCESS: アプリコンテキスト
	apCtx := infra.ConvertCtx(ctx, params.XAccountId)

	// PROCESS: Bodyパース/バリデーション
	shippingInstruction := spec.ShippingInstructionBody{}
	if err := ctx.Bind(&shippingInstruction); err != nil {
		return err
	}
	if err := ctx.Validate(shippingInstruction); err != nil {
		errs := err.(validation.Errors)
		for k, err := range errs {
			ctx.Logger().Error(k + ": " + err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// PROCESS: コマンド実行
	// FIXME:repository
	cmd := command.NewShippingIsntructionPostCommand(apCtx, shippingInstruction)
	if err := cmd.Ececute(); err != nil {
		return err
	}

	ctx.Response().Header().Set("location", path.Join("/receivings", cmd.OrderNo))
	return ctx.NoContent(http.StatusCreated)
}
