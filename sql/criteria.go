package sql

import (
	"database/sql"
	"github.com/davfer/go-specification"
)

var ComparisonConversion = map[specification.Comparator]string{
	specification.ComparisonEq:  "=",
	specification.ComparisonGt:  ">",
	specification.ComparisonGte: ">=",
	specification.ComparisonLt:  "<",
	specification.ComparisonLte: "<=",
	specification.ComparisonNe:  "!=",
}

type Criteria interface {
	GetQuery(table string) string
	GetParams() []sql.NamedArg
}
