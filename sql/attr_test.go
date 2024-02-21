package sql

import (
	"database/sql"
	"github.com/davfer/go-specification"
	"reflect"
	"testing"
)

func TestMongoAttr_GetExpression(t *testing.T) {
	tests := []struct {
		name       string
		attr       string
		value      any
		params     []sql.NamedArg
		table      string
		comparison specification.Comparator
		want       string
	}{
		{
			name:  "Test eq",
			attr:  "some_column",
			value: 12,
			params: []sql.NamedArg{
				{Name: "some_column_2758240694", Value: 12},
			},
			table:      "some_table",
			comparison: specification.ComparisonEq,
			want:       "some_table.some_column = :some_column_2758240694",
		},
		{
			name:  "Test gt",
			attr:  "some_column",
			value: 12,
			params: []sql.NamedArg{
				{Name: "some_column_3667151261", Value: 12},
			},
			table:      "some_table",
			comparison: specification.ComparisonGt,
			want:       "some_table.some_column > :some_column_3667151261",
		},
		{
			name:  "Test gte",
			attr:  "some_column",
			value: 12,
			params: []sql.NamedArg{
				{Name: "some_column_1789139740", Value: 12},
			},
			table:      "some_table",
			comparison: specification.ComparisonGte,
			want:       "some_table.some_column >= :some_column_1789139740",
		},
		{
			name:  "Test lt",
			attr:  "some_column",
			value: 12,
			params: []sql.NamedArg{
				{Name: "some_column_3603015292", Value: 12},
			},
			table:      "some_table",
			comparison: specification.ComparisonLt,
			want:       "some_table.some_column < :some_column_3603015292",
		},
		{
			name:  "Test lte",
			attr:  "some_column",
			value: 12,
			params: []sql.NamedArg{
				{Name: "some_column_3181747229", Value: 12},
			},
			table:      "some_table",
			comparison: specification.ComparisonLte,
			want:       "some_table.some_column <= :some_column_3181747229",
		},
		{
			name:  "Test ne",
			attr:  "some_column",
			value: 12,
			params: []sql.NamedArg{
				{Name: "some_column_2258391298", Value: 12},
			},
			table:      "some_table",
			comparison: specification.ComparisonNe,
			want:       "some_table.some_column != :some_column_2258391298",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Attr{
				Name:       tt.attr,
				Value:      tt.value,
				Comparison: tt.comparison,
			}
			if got := a.GetQuery(tt.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpression() = %v, want %v", got, tt.want)
			}
			if got := a.GetParams(); !reflect.DeepEqual(got, tt.params) {
				t.Errorf("GetParams() = %v, want %v", got, tt.params)
			}
		})
	}
}
