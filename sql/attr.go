package sql

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"github.com/davfer/go-specification"
	"hash/crc32"
	"strconv"
	"strings"
)

type Attr struct {
	Name       string
	Value      any
	Comparison specification.Comparator
}

func (a Attr) GetQuery(table string) string {
	column := fmt.Sprintf("%s.%s", table, a.Name)
	id := fmt.Sprintf(":%s", a.getParamHash())

	return fmt.Sprintf("%s %s %s", column, ComparisonConversion[a.Comparison], id)
}

func (a Attr) GetParams() []sql.NamedArg {
	return []sql.NamedArg{
		sql.Named(a.getParamHash(), a.Value),
	}
}

func (a Attr) getParamHash() string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(a)
	return fmt.Sprintf("%s_%s", strings.ToLower(a.Name), strconv.Itoa(int(crc32.ChecksumIEEE(b.Bytes()))))
}
