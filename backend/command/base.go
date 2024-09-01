/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/teru-0529/go-webapp-echo-1st/infra"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
)

// STRUCT:
type Command interface {
	Execute(ctx context.Context)
}

// STRUCT:
type Invoker struct {
	Cmd Command
}

// FUNCTION:
func NewInvoker(eCtx echo.Context, accountId string, cmd Command) (*Invoker, func()) {
	ivc := &Invoker{Cmd: cmd}

	// PROCESS: アプリコンテキストの変換/セット
	ctx := infra.ConvertCtx(eCtx, accountId)
	eCtx.Logger().Debug("traceId: " + infra.TraceId(ctx))

	return ivc, nil
}

// STRUCT:
type QueryBase struct {
	limit  spec.Limit
	offset spec.Offset
}

// FUNCTION:
func NewQueryBase(limit *spec.Limit, offset *spec.Offset) QueryBase {
	qb := QueryBase{limit: 20, offset: 0}
	if limit != nil {
		qb.limit = *limit
	}
	if offset != nil {
		qb.offset = *offset
	}
	return qb
}
