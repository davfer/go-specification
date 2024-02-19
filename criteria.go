package specification

import "context"

// Criteria is the interface that wraps the basic method to check if a value satisfies a condition.
type Criteria interface {
	// IsSatisfiedBy returns true if the value satisfies the condition.
	IsSatisfiedBy(any) bool
}

// Repository is the interface that wraps the basic methods to match entities.
type Repository[K any] interface {
	Match(context.Context, Criteria) ([]K, error)
	MatchOne(context.Context, Criteria) (K, error)
}
