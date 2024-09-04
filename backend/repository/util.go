/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package repository

import (
	"context"

	"github.com/teru-0529/go-webapp-echo-1st/infra"
	"github.com/volatiletech/null/v8"
)

// TITLE: 変更対象フィールド

type RegistrationCols struct {
	InsertCols []string
	UpdateCols []string
}

// FUNCTION: new
func NewRegistrationCols() *RegistrationCols {
	return &RegistrationCols{
		InsertCols: []string{"created_by", "updated_by"},
		UpdateCols: []string{"updated_by"},
	}
}

// FUNCTION:
func (cs *RegistrationCols) add(col string) {
	cs.InsertCols = append(cs.InsertCols, col)
	cs.UpdateCols = append(cs.UpdateCols, col)
}

// FUNCTION: generateTraceId
func traceId(ctx context.Context) null.String {
	return null.StringFrom(infra.TraceId(ctx))
}
