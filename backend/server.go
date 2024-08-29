/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package main

import (
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

	e.Logger.Fatal(e.Start(":7011"))
}
