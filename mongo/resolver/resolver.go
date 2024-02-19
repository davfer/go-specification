package mongo

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo"
)

func NewMongoResolver() specification.Converter[mongo.Criteria] {
	return specification.Converter[mongo.Criteria]{
		Resolvers: []specification.Resolver[mongo.Criteria]{
			Or{},
			And{},
			Attr{},
			Not{},
		},
	}
}
