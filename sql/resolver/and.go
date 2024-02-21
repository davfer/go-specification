package resolver

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/sql"
)

type And struct {
}

func (o And) Resolve(conv specification.Converter[sql.Criteria], c specification.Criteria, sub any) (sql.Criteria, bool) {
	an, ok := c.(specification.And)
	if !ok {
		return nil, false
	}

	var ops []sql.Criteria
	for _, operand := range an.Operands {
		mc, err := conv.Convert(operand, sub)
		if err != nil {
			return nil, false
		}
		ops = append(ops, mc)
	}
	return sql.And{Operands: ops}, true
}
