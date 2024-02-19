package mongo

import (
	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo"
	"reflect"
	"strings"
)

type Attr struct {
}

func (o Attr) Resolve(conv specification.Converter[mongo.Criteria], c specification.Criteria, sub any) (mongo.Criteria, bool) {
	ca, ok := c.(specification.Attr)
	if !ok {
		return nil, false
	}

	field, ok := reflect.TypeOf(sub).Elem().FieldByName(ca.Name)
	if !ok {
		return nil, false
	}

	tag := field.Tag.Get("bson")
	if tag == "" {
		return nil, false
	}
	if strings.Contains(tag, ",") {
		tag = strings.Split(tag, ",")[0]
	}

	return mongo.Attr{Name: tag, Value: ca.Value, Comparison: ca.Comparison}, true
}
