/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teru-0529/go-webapp-echo-1st/infra"
	spec "github.com/teru-0529/go-webapp-echo-1st/spec/apispec"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var ErrNotFound = errors.New("resource not found.")
var ErrOverflow = errors.New("more than permissible value.")

// STRUCT:
type Command interface {
	Execute(ctx context.Context, tx *sql.Tx) error
}

// STRUCT:
type Invoker struct {
	ctx context.Context
	Cmd Command
}

// FUNCTION:
func NewInvoker(eCtx echo.Context, accountId string, cmd Command) *Invoker {
	// PROCESS: アプリコンテキストの変換/セット
	ctx := infra.ConvertCtx(eCtx, accountId)
	eCtx.Logger().Debug("traceId: " + infra.TraceId(ctx))

	return &Invoker{ctx: ctx, Cmd: cmd}
}

// FUNCTION:
func (inv *Invoker) Execute() error {

	// PROCESS: トランザクションの取得/セット
	tx, err := boil.BeginTx(inv.ctx, nil)
	if err != nil {
		return err
	}

	// PROCESS: 処理実行
	if err := inv.Cmd.Execute(inv.ctx, tx); err != nil {
		tx.Rollback()
		if errors.Is(err, ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		} else if errors.Is(err, ErrOverflow) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	tx.Commit()
	return nil
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

// FUNCTION:
func (qb *QueryBase) Qm() []qm.QueryMod {

	// PROCESS:
	mods := []qm.QueryMod{}
	mods = append(mods, qm.Limit(qb.limit+1))
	mods = append(mods, qm.Offset(qb.offset))
	return mods
}

// FUNCTION: generateTraceId
func traceId(ctx context.Context) null.String {
	return null.StringFrom(infra.TraceId(ctx))
}
