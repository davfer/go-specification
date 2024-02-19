package specification

type Comparator string

type Criteria interface {
	IsSatisfiedBy(any) bool
}
