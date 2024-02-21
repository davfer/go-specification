package resolver

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/sql"
	"strings"
)

type Attr struct {
}

func (o Attr) Resolve(_ specification.Converter[sql.Criteria], c specification.Criteria, sub any) (sql.Criteria, bool) {
	ca, ok := c.(specification.Attr)
	if !ok {
		return nil, false
	}

	return sql.Attr{Name: o.nameToColumn(ca.Name), Value: ca.Value, Comparison: ca.Comparison}, true
}

func (o Attr) nameToColumn(name string) string {
	var column strings.Builder
	for i, r := range name {
		if i > 0 && r >= 'A' && r <= 'Z' {
			column.WriteRune('_')
		}
		column.WriteRune(r)
	}

	return strings.ToLower(column.String())
}
