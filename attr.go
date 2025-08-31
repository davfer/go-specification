package specification

import (
	"fmt"
	"reflect"
)

type Comparator string

const (
	ComparisonEq  Comparator = "eq"
	ComparisonNe  Comparator = "ne"
	ComparisonGt  Comparator = "gt"
	ComparisonGte Comparator = "gte"
	ComparisonLt  Comparator = "lt"
	ComparisonLte Comparator = "lte"
	//ComparisonCo  Comparator = "co"
	//ComparisonSw  Comparator = "sw"
	//ComparisonEw  Comparator = "ew"
	//ComparisonIn  Comparator = "in"
	//ComparisonNin Comparator = "nin"
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

	switch a.Comparison {
	case ComparisonEq:
		return value == a.Value
	case ComparisonNe:
		return value != a.Value
	case ComparisonGt:
		switch v := value.(type) {
		case int:
			return v > a.Value.(int)
		case float64:
			return v > a.Value.(float64)
		case string:
			if s, ok := a.Value.(fmt.Stringer); ok {
				return v > s.String()
			}
			return v > a.Value.(string)
		}
	case ComparisonGte:
		switch v := value.(type) {
		case int:
			return v >= a.Value.(int)
		case float64:
			return v >= a.Value.(float64)
		case string:
			if s, ok := a.Value.(fmt.Stringer); ok {
				return v >= s.String()
			}
			return v >= a.Value.(string)
		}
	case ComparisonLt:
		switch v := value.(type) {
		case int:
			return v < a.Value.(int)
		case float64:
			return v < a.Value.(float64)
		case string:
			if s, ok := a.Value.(fmt.Stringer); ok {
				return v < s.String()
			}
			return v < a.Value.(string)
		}
	case ComparisonLte:
		switch v := value.(type) {
		case int:
			return v <= a.Value.(int)
		case float64:
			return v <= a.Value.(float64)
		case string:
			if s, ok := a.Value.(fmt.Stringer); ok {
				return v <= s.String()
			}
			return v <= a.Value.(string)
		}
	}

	return false
}
