package mongo

import (
	"reflect"
	"testing"

	"github.com/davfer/go-specification"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMongoAnd_GetExpression(t *testing.T) {
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
			want: bson.M{"$and": []bson.M{{"some_column": bson.M{"$eq": 12}}}},
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
					Name:       "some_column",
					Value:      12,
					Comparison: specification.ComparisonEq,
				},
			},
			want: bson.M{"$and": []bson.M{{"some_column": bson.M{"$eq": 12}}, {"some_column": bson.M{"$eq": 12}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := And{
				Operands: tt.operands,
			}
			if got := a.GetExpression(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
