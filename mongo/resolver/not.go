package mongo

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo"
)

type Not struct {
}

func (n Not) Resolve(conv specification.Converter[mongo.Criteria], c specification.Criteria, sub any) (mongo.Criteria, bool) {
	no, ok := c.(specification.Not)
	if !ok {
		return nil, false
	}

	mc, err := conv.Convert(no.Operand, sub)
	if err != nil {
		return nil, false
	}

	return mongo.Not{Operand: mc}, true
}
