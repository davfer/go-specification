package specification

import "testing"

type testAndEntity struct {
	Attr1    string
	IntNice  int
	Floating float64
}

func TestAnd_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name     string
		operands []Criteria
		value    any
		want     bool
	}{
		{
			name:     "Test empty operands",
			operands: []Criteria{},
			value:    testAndEntity{},
			want:     true,
		},
		{
			name: "Test single operand true",
			operands: []Criteria{
				Attr{
					Name:       "Attr1",
					Value:      "test",
					Comparison: ComparisonEq,
				},
			},
			value: testAndEntity{
				Attr1: "test",
			},
			want: true,
		},
		{
			name: "Test single operand false",
			operands: []Criteria{
				Attr{
					Name:       "Attr1",
					Value:      "testa",
					Comparison: ComparisonEq,
				},
			},
			value: testAndEntity{
				Attr1: "test",
			},
			want: false,
		},
		{
			name: "Test single operand true, true",
			operands: []Criteria{
				Attr{
					Name:       "Attr1",
					Value:      "test",
					Comparison: ComparisonEq,
				},
				Attr{
					Name:       "IntNice",
					Value:      1,
					Comparison: ComparisonGt,
				},
			},
			value: testAndEntity{
				Attr1:   "test",
				IntNice: 2,
			},
			want: true,
		},
		{
			name: "Test single operand true, false",
			operands: []Criteria{
				Attr{
					Name:       "Attr1",
					Value:      "test",
					Comparison: ComparisonEq,
				},
				Attr{
					Name:       "IntNice",
					Value:      1,
					Comparison: ComparisonGt,
				},
			},
			value: testAndEntity{
				Attr1:   "test",
				IntNice: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := And{
				Operands: tt.operands,
			}
			if got := a.IsSatisfiedBy(tt.value); got != tt.want {
				t.Errorf("IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
