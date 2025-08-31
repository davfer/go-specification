package specification

import (
	"fmt"
)

// ErrNoResolverFound is the error returned when no resolver is found for a given criteria
var ErrNoResolverFound = fmt.Errorf("no resolver found")

// CriteriaPrimitive is a helper interface to allow the conversion of a Criteria to a primitive Criteria.
// Useful to work with custom criteria types.
type CriteriaPrimitive interface {
	Criteria
	GetPrimitive() Criteria
}

// Resolver is a helper interface to allow the conversion of a Criteria to a different implementation
type Resolver[K any] interface {
	Resolve(Converter[K], Criteria, any) (K, bool)
}

// Converter is a helper to convert a Criteria to a different implementation with a set of resolvers
type Converter[K any] struct {
	Resolvers []Resolver[K]
}

func (c Converter[K]) Convert(criteria Criteria, subject any) (k K, err error) {
	for _, resolver := range c.Resolvers {
		var ok bool
		if k, ok = resolver.Resolve(c, criteria, subject); ok {
			return
		}
	}

	err = ErrNoResolverFound
	return
}
