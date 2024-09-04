/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package main

import (
	_ "github.com/lib/pq"
	"github.com/teru-0529/go-webapp-echo-1st/controller"
	"github.com/teru-0529/go-webapp-echo-1st/infra"
	"github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

const DEBUG_MODE = true

// FUNCTION:
func main() {

	// データベース(Sqlboiler)の設定
	cleanUp := infra.InitDB(DEBUG_MODE)
	defer cleanUp()

	// Webサーバー(echo)の設定
	e := infra.InitServer(DEBUG_MODE)

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	api := controller.ApiController{}
	apispec.RegisterHandlers(e, api)

	// サーバースタート
	e.Logger.Fatal(e.Start(":7011"))
}
