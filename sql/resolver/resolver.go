package resolver

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/sql"
)

func NewSqlConverter(extraResolvers ...specification.Resolver[sql.Criteria]) specification.Converter[sql.Criteria] {
	resolvers := []specification.Resolver[sql.Criteria]{
		Attr{},
		And{},
		Or{},
	}
	resolvers = append(resolvers, extraResolvers...)

	return specification.Converter[sql.Criteria]{
		Resolvers: resolvers,
	}
}
