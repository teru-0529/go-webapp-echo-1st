/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teru-0529/go-webapp-echo-1st/adapter"
	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// FUNCTION:
func main() {
	e := echo.New()

	// ロガーのミドルウェアを設定
	e.Use(middleware.Logger())
	// APIがエラーで落ちてもリカバーするミドルウェアを設定
	e.Use(middleware.Recover())

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	api := adapter.ApiController{}
	spec.RegisterHandlers(e, api)

	e.HTTPErrorHandler = CustomHTTPErrorHandler

	e.Logger.Fatal(e.Start(":7011"))
}

// TITLE:
type ApiError struct {
	Types  string `json:"types"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// FUNCTION:
func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusBadRequest
	message := ""
	// https://godoc.org/github.com/labstack/echo#HTTPError
	if ee, ok := err.(*echo.HTTPError); ok {
		code = ee.Code
		message = ee.Message.(string)
	}
	body := ApiError{
		Types:  "/errors/description",
		Title:  fmt.Sprintf("%d %s", code, http.StatusText(code)),
		Detail: message,
	}
	c.JSON(code, body)
}
