package mongo

import (
	"github.com/davfer/go-specification"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Attr struct {
	Name       string
	Value      any
	Comparison specification.Comparator
}

func (a Attr) GetExpression() bson.M {
	return bson.M{a.Name: bson.M{ComparisonConversion[a.Comparison]: a.Value}}
}
