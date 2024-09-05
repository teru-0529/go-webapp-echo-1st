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

// FUNCTION:
func main() {
	// PROCESS: envファイルのロード
	config := infra.LeadEnv()

	// PROCESS: データベース(Sqlboiler)の設定
	cleanUp := infra.InitDB(config)
	defer cleanUp()

	// PROCESS: Webサーバー(echo)の設定
	e := infra.InitServer(config.DebugMode)

	// PROCESS: OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	api := controller.ApiController{}
	apispec.RegisterHandlers(e, api)

	// サーバースタート
	e.Logger.Fatal(e.Start(":8080"))
}
