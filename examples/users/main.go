package main

import (
	"github.com/davfer/go-specification"
)

type User struct {
	Username string
	Age      int
}

type AdultUsersCriteria struct {
}

func (a *AdultUsersCriteria) IsSatisfiedBy(u any) bool {
	spec := specification.Attr{
		Name:       "Age",
		Value:      18,
		Comparison: specification.ComparisonGte,
	}

	return spec.IsSatisfiedBy(u)
}

func main() {
	adultUsersCriteria := &AdultUsersCriteria{}
	user1 := User{
		Username: "user1",
		Age:      20,
	}
	user2 := User{
		Username: "user2",
		Age:      15,
	}

	if adultUsersCriteria.IsSatisfiedBy(user1) {
		println("User 1 is an adult")
	} else {
		println("User 1 is not an adult")
	}
	if adultUsersCriteria.IsSatisfiedBy(user2) {
		println("User 2 is an adult")
	} else {
		println("User 2 is not an adult")
	}
}
