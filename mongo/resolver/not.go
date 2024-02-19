package mongo

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo"
)

type Not struct {
}

func (n Not) Resolve(conv specification.Converter[mongo.Criteria], c specification.Criteria, sub any) (mongo.Criteria, bool) {
	mc, err := conv.Convert(c.(specification.Not).Operand, sub)
	if err != nil {
		return nil, false
	}

	return mongo.Not{Operand: mc}, true
}
