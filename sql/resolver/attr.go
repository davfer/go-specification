package resolver

import (
	"github.com/davfer/archit/helpers/str"
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

	return sql.Attr{Name: strings.ToLower(str.Convert(ca.Name, str.Pascal, str.Snake)), Value: ca.Value, Comparison: ca.Comparison}, true
}
