package specification

import "testing"

type testNotEntity struct {
	Attr1    string
	IntNice  int
	Floating float64
}

func TestNot_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name    string
		operand Criteria
		value   any
		want    bool
	}{
		{
			name: "Test true",
			operand: Attr{
				Name:       "Attr1",
				Value:      "test",
				Comparison: ComparisonEq,
			},
			value: testNotEntity{
				Attr1: "testa",
			},
			want: true,
		},
		{
			name: "Test false",
			operand: Attr{
				Name:       "Attr1",
				Value:      "test",
				Comparison: ComparisonEq,
			},
			value: testNotEntity{
				Attr1: "test",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := Not{
				Operand: tt.operand,
			}
			if got := n.IsSatisfiedBy(tt.value); got != tt.want {
				t.Errorf("IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
