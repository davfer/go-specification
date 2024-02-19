# go-specification

Simple Specification Pattern for Go.

## Usage

```go
package main

import (
	"github.com/davfer/go-specification"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	specGt := specification.Attr{
		Name:       "Age",
		Value:      20,
		Compatison: specification.ComparisonGt,
	}
	specLt := specification.Attr{
		Name:       "Age",
		Value:      50,
		Compatison: specification.ComparisonLt,
	}
	spec := specification.And{Operands: []specification.Criteria{specGt, specLt}}
	users := []User{
		{Name: "John", Age: 20},
		{Name: "Jane", Age: 30},
		{Name: "Jack", Age: 40},
		{Name: "Jill", Age: 50},
	}

	var result []User
	for _, user := range users {
		if spec.IsSatisfiedBy(user) {
			result = append(result, user)
		}
	}

	fmt.Println(result) // [{Name: "Jane", Age: 30}, {Name: "Jack", Age: 40}]
}
```