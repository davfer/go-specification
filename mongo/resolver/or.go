package resolver

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo"
)

type Or struct {
}

func (o Or) Resolve(conv specification.Converter[mongo.Criteria], c specification.Criteria, sub any) (mongo.Criteria, bool) {
	s, ok := c.(specification.Or)
	if !ok {
		return nil, false
	}

	var ops []mongo.Criteria
	for _, operand := range s.Operands {
		mc, err := conv.Convert(operand, sub)
		if err != nil {
			return nil, false
		}
		ops = append(ops, mc)
	}
	return mongo.Or{Operands: ops}, true
}
