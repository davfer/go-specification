package specification

import (
	"reflect"
	"testing"
)

type testConverter struct {
}

type testResolver struct {
}

func (r testResolver) Resolve(conv Converter[int], c Criteria, sub any) (int, bool) {
	return 12, true
}

func TestConverter_Convert(t *testing.T) {
	type testCase[K any] struct {
		name     string
		c        Converter[int]
		criteria Criteria
		subject  any
		wantK    int
		wantErr  bool
	}
	tests := []testCase[testConverter]{
		{
			name: "Test no resolver",
			c:    Converter[int]{},
			criteria: Attr{
				Name:       "some_column",
				Value:      12,
				Comparison: ComparisonEq,
			},
			subject: 12,
			wantK:   0,
			wantErr: true,
		},
		{
			name: "Test one resolver",
			c: Converter[int]{Resolvers: []Resolver[int]{
				testResolver{},
			}},
			criteria: Attr{
				Name:       "some_column",
				Value:      12,
				Comparison: ComparisonEq,
			},
			subject: 12,
			wantK:   12,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK, err := tt.c.Convert(tt.criteria, tt.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotK, tt.wantK) {
				t.Errorf("Convert() gotK = %v, want %v", gotK, tt.wantK)
			}
		})
	}
}
