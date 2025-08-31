package mongo

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo/v2"
)

type And struct {
}

func (o And) Resolve(conv specification.Converter[mongo.Criteria], c specification.Criteria, sub any) (mongo.Criteria, bool) {
	an, ok := c.(specification.And)
	if !ok {
		return nil, false
	}

	var ops []mongo.Criteria
	for _, operand := range an.Operands {
		mc, err := conv.Convert(operand, sub)
		if err != nil {
			return nil, false
		}
		ops = append(ops, mc)
	}
	return mongo.And{Operands: ops}, true
}
