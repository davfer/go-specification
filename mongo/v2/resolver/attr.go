package mongo

import (
	"reflect"
	"strings"

	"github.com/davfer/archit/helpers/str"

	"github.com/davfer/go-specification"
	"github.com/davfer/go-specification/mongo/v2"
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
		return mongo.Attr{Name: ca.Name, Value: ca.Value, Comparison: ca.Comparison}, true
	}

	tag := field.Tag.Get("bson")
	switch {
	case strings.Contains(tag, ","):
		tag = strings.Split(tag, ",")[0]
	case tag == "-":
		return nil, false
	case tag == "":
		_, cs := str.GetWords(ca.Name)
		if len(cs) > 0 {
			tag = str.Convert(ca.Name, cs[0], str.Camel)
		}
	}

	return mongo.Attr{Name: tag, Value: ca.Value, Comparison: ca.Comparison}, true
}
