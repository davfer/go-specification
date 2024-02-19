package mongo

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo"
)

type Primitive struct {
}

func (o Primitive) Resolve(conv specification.Converter[mongo.Criteria], c specification.Criteria, sub any) (mongo.Criteria, bool) {
	pr, ok := c.(specification.CriteriaPrimitive)
	if !ok {
		return nil, false
	}

	res, err := conv.Convert(pr.GetPrimitive(), sub)
	if err != nil {
		return nil, false
	}

	return res, true
}
