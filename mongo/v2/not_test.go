package mongo

import (
	"reflect"
	"testing"

	"github.com/davfer/go-specification"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMongoNot_GetExpression(t *testing.T) {
	tests := []struct {
		name    string
		operand Criteria
		want    bson.M
	}{
		{
			name: "Test operand",
			operand: Attr{
				Name:       "some_column",
				Value:      12,
				Comparison: specification.ComparisonEq,
			},
			want: bson.M{"$not": bson.M{"some_column": bson.M{"$eq": 12}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Not{
				Operand: tt.operand,
			}
			if got := n.GetExpression(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
