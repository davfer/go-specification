package mongo

import (
	"github.com/davfer/go-specification"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var ComparisonConversion = map[specification.Comparator]string{
	specification.ComparisonEq:  "$eq",
	specification.ComparisonGt:  "$gt",
	specification.ComparisonGte: "$gte",
	specification.ComparisonLt:  "$lt",
	specification.ComparisonLte: "$lte",
	specification.ComparisonNe:  "$ne",
}

type Criteria interface {
	GetExpression() bson.M
}
