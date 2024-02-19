package specification

type Not struct {
	Operand Criteria
}

func (n Not) IsSatisfiedBy(value any) bool {
	return !n.Operand.IsSatisfiedBy(value)
}
