package mongo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type And struct {
	Operands []Criteria
}

func (o And) GetExpression() bson.M {
	var expressions []bson.M
	for _, operand := range o.Operands {
		expressions = append(expressions, operand.GetExpression())
	}

	return bson.M{"$and": expressions}
}
