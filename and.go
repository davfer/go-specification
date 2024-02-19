package specification

type And struct {
	Operands []Criteria
}

func (a And) IsSatisfiedBy(value any) bool {
	if len(a.Operands) == 0 {
		return true
	}

	for _, operand := range a.Operands {
		if !operand.IsSatisfiedBy(value) {
			return false
		}
	}

	return true
}
