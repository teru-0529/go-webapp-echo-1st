/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package infra

// TITLE:context設定

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type contextKey string

const (
	REQUEST_ID contextKey = "REQUEST_ID"
	ACCOUNT_ID contextKey = "ACCOUNT_ID"
)

// FUNCTION: context setting
func ConvertCtx(eCtx echo.Context, accountId string) context.Context {
	ctx := eCtx.Request().Context()
	ctx = context.WithValue(ctx, REQUEST_ID, ulid.Make().String())
	ctx = context.WithValue(ctx, ACCOUNT_ID, accountId)
	return ctx
}

// FUNCTION:: getVal(AccountId)
func AccountId(ctx context.Context) string {
	val, ok := ctx.Value(ACCOUNT_ID).(string)
	if !ok {
		return "N/A"
	}
	return val
}

// FUNCTION:: getVal(TraceId)
func TraceId(ctx context.Context) string {
	requestId, ok := ctx.Value(REQUEST_ID).(string)
	if !ok {
		return "N/A"
	}
	accountId, ok := ctx.Value(ACCOUNT_ID).(string)
	if !ok {
		return "N/A"
	}
	return fmt.Sprintf("%s::%s", requestId, accountId)
}
