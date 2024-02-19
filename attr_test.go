package specification

import (
	"math"
	"testing"
)

type testAttrEntity struct {
	id            string
	Attr1         string
	someOtherAttr bool
	IntNice       int
	Floating      float64
}

func TestAttr_IsSatisfiedBy(t *testing.T) {
	type fields struct {
		Name       string
		Value      any
		Comparison Comparator
	}
	tests := []struct {
		name   string
		fields fields
		value  any
		want   bool
		panic  bool
	}{
		{
			name: "Test unexported field",
			fields: fields{
				Name:       "id",
				Value:      "123",
				Comparison: ComparisonEq,
			},
			panic: true,
		},
		{
			name: "Test empty field",
			fields: fields{
				Name:       "",
				Value:      "123",
				Comparison: ComparisonEq,
			},
			panic: true,
		},
		{
			name: "Test eq true",
			fields: fields{
				Name:       "Floating",
				Value:      math.Pi,
				Comparison: ComparisonEq,
			},
			value: testAttrEntity{
				Floating: math.Pi,
			},
			want: true,
		},
		{
			name: "Test eq false",
			fields: fields{
				Name:       "IntNice",
				Value:      1,
				Comparison: ComparisonEq,
			},
			value: testAttrEntity{
				IntNice: 2,
			},
			want: false,
		},
		{
			name: "Test ne true",
			fields: fields{
				Name:       "IntNice",
				Value:      1,
				Comparison: ComparisonNe,
			},
			value: testAttrEntity{
				IntNice: 2,
			},
			want: true,
		},
		{
			name: "Test ne false",
			fields: fields{
				Name:       "IntNice",
				Value:      1,
				Comparison: ComparisonNe,
			},
			value: testAttrEntity{
				IntNice: 1,
			},
			want: false,
		},
		{
			name: "Test gt true",
			fields: fields{
				Name:       "Floating",
				Value:      1.0,
				Comparison: ComparisonGt,
			},
			value: testAttrEntity{
				Floating: 1.1,
			},
			want: true,
		},
		{
			name: "Test gt string true",
			fields: fields{
				Name:       "Attr1",
				Value:      "a",
				Comparison: ComparisonGt,
			},
			value: testAttrEntity{
				Attr1: "a",
			},
			want: false,
		},
		{
			name: "Test gt false",
			fields: fields{
				Name:       "Floating",
				Value:      1.0,
				Comparison: ComparisonGt,
			},
			value: testAttrEntity{
				Floating: 1.0,
			},
		},
		{
			name: "Test gte true",
			fields: fields{
				Name:       "Floating",
				Value:      2.34,
				Comparison: ComparisonGte,
			},
			value: testAttrEntity{
				Floating: 2.34,
			},
			want: true,
		},
		{
			name: "Test gte false",
			fields: fields{
				Name:       "IntNice",
				Value:      12,
				Comparison: ComparisonGte,
			},
			value: testAttrEntity{
				IntNice: 1,
			},
			want: false,
		},

		{
			name: "Test gte false",
			fields: fields{
				Name:       "Attr1",
				Value:      "a",
				Comparison: ComparisonGte,
			},
			value: testAttrEntity{
				Attr1: "a",
			},
			want: true,
		},
		{
			name: "Test lt true",
			fields: fields{
				Name:       "IntNice",
				Value:      12,
				Comparison: ComparisonLt,
			},
			value: testAttrEntity{
				IntNice: 1,
			},
			want: true,
		},
		{
			name: "Test lt false",
			fields: fields{
				Name:       "IntNice",
				Value:      12,
				Comparison: ComparisonLt,
			},
			value: testAttrEntity{
				IntNice: 12,
			},
		},
		{
			name: "Test lte true",
			fields: fields{
				Name:       "IntNice",
				Value:      12,
				Comparison: ComparisonLte,
			},
			value: testAttrEntity{
				IntNice: 12,
			},
			want: true,
		},
		{
			name: "Test lte false",
			fields: fields{
				Name:       "IntNice",
				Value:      12,
				Comparison: ComparisonLte,
			},
			value: testAttrEntity{
				IntNice: 13,
			},
			want: false,
		},
		{
			name: "Test non struct",
			fields: fields{
				Name:       "IntNice",
				Value:      12,
				Comparison: ComparisonLte,
			},
			value: 12,
			want:  false,
		},
		{
			name: "Test non existing field",
			fields: fields{
				Name:       "NonExisting",
				Value:      12,
				Comparison: ComparisonLte,
			},
			value: testAttrEntity{},
			want:  false,
		},
		{
			name: "Test non existing comparison",
			fields: fields{
				Name:       "IntNice",
				Value:      12,
				Comparison: "NonExisting",
			},
			value: testAttrEntity{},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}
			a := Attr{
				Name:       tt.fields.Name,
				Value:      tt.fields.Value,
				Comparison: tt.fields.Comparison,
			}
			if got := a.IsSatisfiedBy(tt.value); got != tt.want {
				t.Errorf("IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
