// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package spec

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// キャンセル指示登録
	// (POST /cancel-instructions)
	OrdersCancelInstructionsPost(ctx echo.Context) error
	// 受注一覧検索
	// (GET /receivings)
	OrdersReceivingsGet(ctx echo.Context, params OrdersReceivingsGetParams) error
	// 受注登録
	// (POST /receivings)
	OrdersReceivingsPost(ctx echo.Context) error
	// 受注取得
	// (GET /receivings/{order_no})
	OrdersReceivingsNoGet(ctx echo.Context, orderNo OrderNo) error
	// 受注修正
	// (PATCH /receivings/{order_no})
	OrdersReceivingsNoPatch(ctx echo.Context, orderNo OrderNo) error
	// 出荷指示登録
	// (POST /shipping-instructions)
	OrdersShippingInstructionsPost(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// OrdersCancelInstructionsPost converts echo context to params.
func (w *ServerInterfaceWrapper) OrdersCancelInstructionsPost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OrdersCancelInstructionsPost(ctx)
	return err
}

// OrdersReceivingsGet converts echo context to params.
func (w *ServerInterfaceWrapper) OrdersReceivingsGet(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OrdersReceivingsGetParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "customer_name" -------------

	err = runtime.BindQueryParameter("form", true, false, "customer_name", ctx.QueryParams(), &params.CustomerName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter customer_name: %s", err))
	}

	// ------------- Optional query parameter "order_status" -------------

	err = runtime.BindQueryParameter("form", true, false, "order_status", ctx.QueryParams(), &params.OrderStatus)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order_status: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OrdersReceivingsGet(ctx, params)
	return err
}

// OrdersReceivingsPost converts echo context to params.
func (w *ServerInterfaceWrapper) OrdersReceivingsPost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OrdersReceivingsPost(ctx)
	return err
}

// OrdersReceivingsNoGet converts echo context to params.
func (w *ServerInterfaceWrapper) OrdersReceivingsNoGet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "order_no" -------------
	var orderNo OrderNo

	err = runtime.BindStyledParameterWithLocation("simple", false, "order_no", runtime.ParamLocationPath, ctx.Param("order_no"), &orderNo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order_no: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OrdersReceivingsNoGet(ctx, orderNo)
	return err
}

// OrdersReceivingsNoPatch converts echo context to params.
func (w *ServerInterfaceWrapper) OrdersReceivingsNoPatch(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "order_no" -------------
	var orderNo OrderNo

	err = runtime.BindStyledParameterWithLocation("simple", false, "order_no", runtime.ParamLocationPath, ctx.Param("order_no"), &orderNo)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order_no: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OrdersReceivingsNoPatch(ctx, orderNo)
	return err
}

// OrdersShippingInstructionsPost converts echo context to params.
func (w *ServerInterfaceWrapper) OrdersShippingInstructionsPost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.OrdersShippingInstructionsPost(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/cancel-instructions", wrapper.OrdersCancelInstructionsPost)
	router.GET(baseURL+"/receivings", wrapper.OrdersReceivingsGet)
	router.POST(baseURL+"/receivings", wrapper.OrdersReceivingsPost)
	router.GET(baseURL+"/receivings/:order_no", wrapper.OrdersReceivingsNoGet)
	router.PATCH(baseURL+"/receivings/:order_no", wrapper.OrdersReceivingsNoPatch)
	router.POST(baseURL+"/shipping-instructions", wrapper.OrdersShippingInstructionsPost)

}
