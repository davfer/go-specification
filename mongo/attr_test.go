package mongo

import (
	"github.com/davfer/go-specification"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestMongoAttr_GetExpression(t *testing.T) {
	tests := []struct {
		name       string
		attr       string
		value      any
		comparison specification.Comparator
		want       bson.M
	}{
		{
			name:       "Test eq",
			attr:       "some_column",
			value:      12,
			comparison: specification.ComparisonEq,
			want:       bson.M{"some_column": bson.M{"$eq": 12}},
		},
		{
			name:       "Test gt",
			attr:       "some_column",
			value:      12,
			comparison: specification.ComparisonGt,
			want:       bson.M{"some_column": bson.M{"$gt": 12}},
		},
		{
			name:       "Test gte",
			attr:       "some_column",
			value:      12,
			comparison: specification.ComparisonGte,
			want:       bson.M{"some_column": bson.M{"$gte": 12}},
		},
		{
			name:       "Test lt",
			attr:       "some_column",
			value:      12,
			comparison: specification.ComparisonLt,
			want:       bson.M{"some_column": bson.M{"$lt": 12}},
		},
		{
			name:       "Test lte",
			attr:       "some_column",
			value:      12,
			comparison: specification.ComparisonLte,
			want:       bson.M{"some_column": bson.M{"$lte": 12}},
		},
		{
			name:       "Test ne",
			attr:       "some_column",
			value:      12,
			comparison: specification.ComparisonNe,
			want:       bson.M{"some_column": bson.M{"$ne": 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Attr{
				Name:       tt.attr,
				Value:      tt.value,
				Comparison: tt.comparison,
			}
			if got := a.GetExpression(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
