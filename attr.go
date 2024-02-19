package specification

import "reflect"

const (
	ComparisonEq  Comparator = "eq"
	ComparisonNe  Comparator = "ne"
	ComparisonGt  Comparator = "gt"
	ComparisonGte Comparator = "gte"
	ComparisonLt  Comparator = "lt"
	ComparisonLte Comparator = "lte"
)

type Attr struct {
	Name       string
	Value      any
	Comparison Comparator
}

func (a Attr) IsSatisfiedBy(obj any) bool {
	if a.Name == "" {
		panic("Name cannot be empty")
	}

	if a.Name[0] < 'A' || a.Name[0] > 'Z' {
		panic("Name must be an exported attribute (start with a capital letter)")
	}

	refObj := reflect.Indirect(reflect.ValueOf(obj))
	if refObj.Kind() != reflect.Struct {
		return false
	}

	var found bool
	var value any
	for i := 0; i < refObj.NumField(); i++ {
		if refObj.Type().Field(i).Name == a.Name {
			found = true
			value = refObj.Field(i).Interface()
			break
		}
	}

	if !found {
		return false
	}

	if a.Comparison == ComparisonEq {
		return value == a.Value
	} else if a.Comparison == ComparisonNe {
		return value != a.Value
	}

	switch a.Comparison {
	case ComparisonGt:
		switch value.(type) {
		case int:
			return value.(int) > a.Value.(int)
		case float64:
			return value.(float64) > a.Value.(float64)
		case string:
			return value.(string) > a.Value.(string)
		}
	case ComparisonGte:
		switch value.(type) {
		case int:
			return value.(int) >= a.Value.(int)
		case float64:
			return value.(float64) >= a.Value.(float64)
		case string:
			return value.(string) >= a.Value.(string)
		}
	case ComparisonLt:
		switch value.(type) {
		case int:
			return value.(int) < a.Value.(int)
		case float64:
			return value.(float64) < a.Value.(float64)
		case string:
			return value.(string) < a.Value.(string)
		}
	case ComparisonLte:
		switch value.(type) {
		case int:
			return value.(int) <= a.Value.(int)
		case float64:
			return value.(float64) <= a.Value.(float64)
		case string:
			return value.(string) <= a.Value.(string)
		}
	}

	return false
}
