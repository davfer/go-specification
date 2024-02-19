package specification

type Or struct {
	Operands []Criteria
}

func (o Or) IsSatisfiedBy(value any) bool {
	if len(o.Operands) == 0 {
		return true
	}

	for _, operand := range o.Operands {
		if operand.IsSatisfiedBy(value) {
			return true
		}
	}

	return false
}
