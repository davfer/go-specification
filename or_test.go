package specification

import "testing"

type testOrEntity struct {
	Attr1    string
	IntNice  int
	Floating float64
}

func TestOr_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name     string
		operands []Criteria
		value    any
		want     bool
	}{
		{
			name:     "Test empty operands",
			operands: []Criteria{},
			value:    testOrEntity{},
			want:     true,
		},
		{
			name: "Test single operand true",
			operands: []Criteria{
				Attr{
					Name:       "IntNice",
					Value:      42,
					Comparison: ComparisonGt,
				},
			},
			value: testOrEntity{
				IntNice: 50,
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
			value: testOrEntity{
				Attr1: "test",
			},
			want: false,
		},
		{
			name: "Test double operand true, true",
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
			value: testOrEntity{
				Attr1:   "test",
				IntNice: 2,
			},
			want: true,
		},
		{
			name: "Test double operand true, false",
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
			value: testOrEntity{
				Attr1:   "test",
				IntNice: 1,
			},
			want: true,
		},
		{
			name: "Test double operand false, false",
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
			value: testOrEntity{
				Attr1:   "testa",
				IntNice: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := Or{
				Operands: tt.operands,
			}
			if got := o.IsSatisfiedBy(tt.value); got != tt.want {
				t.Errorf("IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
