package mongo

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo"
)

func NewMongoConverter(extraResolvers ...specification.Resolver[mongo.Criteria]) specification.Converter[mongo.Criteria] {
	resolvers := []specification.Resolver[mongo.Criteria]{
		Or{},
		And{},
		Attr{},
		Not{},
		Primitive{},
	}
	resolvers = append(resolvers, extraResolvers...)

	return specification.Converter[mongo.Criteria]{
		Resolvers: resolvers,
	}
}
