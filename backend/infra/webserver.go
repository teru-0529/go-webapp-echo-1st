/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package infra

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// FUNCTION: Web Server setting
func InitServer(isDebug bool) *echo.Echo {

	e := echo.New()
	e.Debug = isDebug

	// ロガーのミドルウェアを設定
	e.Use(middleware.Logger())
	// APIがエラーで落ちてもリカバーするミドルウェアを設定
	e.Use(middleware.Recover())

	// validaterを実装
	e.Validator = &CustomValidator{}

	// エラー応答のカスタマイズ
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	return e
}

// STRUCT:
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

// STRUCT:
type CustomValidator struct{}

// FUNCTION:
func (cv *CustomValidator) Validate(i interface{}) error {
	if c, ok := i.(validation.Validatable); ok {
		return c.Validate()
	}
	return nil
}
