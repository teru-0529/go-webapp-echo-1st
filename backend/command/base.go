/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package command

import (
	"github.com/teru-0529/go-webapp-echo-1st/spec"
)

// STRUCT:
type Command interface {
	Execute()
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
