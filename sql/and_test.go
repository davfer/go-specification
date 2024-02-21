package sql

import (
	"database/sql"
	"github.com/davfer/go-specification"
	"reflect"
	"testing"
)

func TestMongoAnd_GetExpression(t *testing.T) {
	tests := []struct {
		name     string
		operands []Criteria
		params   []sql.NamedArg
		want     string
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
			params: []sql.NamedArg{
				{Name: "some_column_2758240694", Value: 12},
			},
			want: "(some_table.some_column = :some_column_2758240694)",
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
					Name:       "some_column2",
					Value:      12,
					Comparison: specification.ComparisonEq,
				},
			},
			params: []sql.NamedArg{
				{Name: "some_column_2758240694", Value: 12},
				{Name: "some_column2_2051320817", Value: 12},
			},
			want: "(some_table.some_column = :some_column_2758240694 AND some_table.some_column2 = :some_column2_2051320817)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := And{
				Operands: tt.operands,
			}
			if got := o.GetQuery("some_table"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpression() = %v, want %v", got, tt.want)
			}
			if got := o.GetParams(); !reflect.DeepEqual(got, tt.params) {
				t.Errorf("GetParams() = %v, want %v", got, tt.params)
			}
		})
	}
}
