package mongo

import (
	"github.com/davfer/go-specification"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestMongoOr_GetExpression(t *testing.T) {
	tests := []struct {
		name     string
		operands []Criteria
		want     bson.M
	}{
		{
			name: "Test one operand",
			operands: []Criteria{
				Attr{
					Name:       "some_column",
					Value:      12,
					Comparison: specification.ComparisonEq,
				},
			},
			want: bson.M{"$or": []bson.M{{"some_column": bson.M{"$eq": 12}}}},
		},
		{
			name: "Test two operands",
			operands: []Criteria{
				Attr{
					Name:       "some_column",
					Value:      12,
					Comparison: specification.ComparisonEq,
				},
				Attr{
					Name:       "another_column",
					Value:      "howard",
					Comparison: specification.ComparisonNe,
				},
			},
			want: bson.M{"$or": []bson.M{{"some_column": bson.M{"$eq": 12}}, {"another_column": bson.M{"$ne": "howard"}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Or{
				Operands: tt.operands,
			}
			if got := o.GetExpression(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
