package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Not struct {
	Operand Criteria
}

func (n Not) GetExpression() bson.M {
	return bson.M{"$not": n.Operand.GetExpression()}
}
