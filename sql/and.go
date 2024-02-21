package sql

import (
	"database/sql"
	"strings"
)

type And struct {
	Operands []Criteria
}

func (o And) GetQuery(table string) string {
	var expressions []string
	for _, operand := range o.Operands {
		expressions = append(expressions, operand.GetQuery(table))
	}

	return "(" + strings.Join(expressions, " AND ") + ")"
}

func (o And) GetParams() []sql.NamedArg {
	var params []sql.NamedArg
	for _, operand := range o.Operands {
		params = append(params, operand.GetParams()...)
	}

	return params
}
