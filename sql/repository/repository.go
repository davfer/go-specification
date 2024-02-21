package repository

import (
	"context"
	"database/sql"
	"github.com/davfer/go-specification"
	sqlSpecification "github.com/davfer/go-specification/mongo"
	"github.com/pkg/errors"
)

var ErrNoMatch = errors.New("no match found")

type CriteriaRepository[K any] struct {
	Db        *sql.DB
	Converter specification.Converter[sqlSpecification.Criteria]
}

func (r *CriteriaRepository[K]) Match(ctx context.Context, c specification.Criteria) ([]K, error) {
	//var subject K
	//mc, err := r.Converter.Convert(c, subject)
	//if err != nil {
	//	return nil, errors.Wrap(err, "error converting criteria")
	//}
	//
	//
	//
	//var entities []K
	//cursor, err := r.Collection.Find(ctx, mc.GetExpression())
	//if err != nil {
	//	return nil, errors.Wrap(err, "error finding match")
	//}
	//if err = cursor.All(ctx, &entities); err != nil {
	//	return nil, errors.Wrap(err, "error reading match")
	//}
	//
	//if len(entities) == 0 {
	return []K{}, nil
	//}
	//
	//return entities, nil
}

func (r *CriteriaRepository[K]) MatchOne(ctx context.Context, c specification.Criteria) (k K, err error) {
	ks, err := r.Match(ctx, c)
	if err != nil {
		return k, err
	}
	if len(ks) == 0 {
		return k, ErrNoMatch
	}

	k = ks[0]
	return
}
