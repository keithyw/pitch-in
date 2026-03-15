package database

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/keithyw/pitch-in/pkg/model"
	"github.com/keithyw/pitch-in/pkg/repository"
)

type DBStore interface {
	Attach(link model.LinkDef, leftID, rightID int64) error
	Count(m model.Model, filter repository.Filter) (int64, error)
	Create(m model.Model, data map[string]interface{}, result interface{}) error
	Delete(m model.Model) error
	Detach(link model.LinkDef, leftID, rightID int64) error
	FindBy(m model.Model, filter repository.Filter, result interface{}) error
	Get(m model.Model, result interface{}) error
	GetBy(m model.Model, key string, value any, result interface{}) error
	GetClient() DBClient
	GetContext() context.Context
	Select(m model.Model) sq.SelectBuilder
	Update(m model.Model, data map[string]interface{}, result interface{}) error
	MakeQueryFromFilter(filter repository.Filter, q sq.SelectBuilder) sq.SelectBuilder
}

type dbStoreImpl struct {
	Client DBClient
	ctx context.Context
}

func NewDBStore(ctx context.Context, client DBClient) DBStore {
	return &dbStoreImpl{
		Client: client,
		ctx: ctx,
	}
}

func (s *dbStoreImpl) Attach(link model.LinkDef, leftID, rightID  int64) error {
	data := map[string]interface{}{
		link.LeftKey: leftID,
		link.RightKey: rightID,
	}

	q := sq.Insert(link.TableName).SetMap(data).PlaceholderFormat(sq.Question)
	_, err := s.GetClient().Exec(s.ctx, q)
	return err
}

func (s *dbStoreImpl) Count(m model.Model, filter repository.Filter) (int64, error) {
	q := sq.Select("COUNT(*)").From(m.TableName())
	countFilter := filter
    countFilter.Limit = 0
    countFilter.Offset = 0
    countFilter.Sort = ""
    countFilter.Order = ""

	q = s.MakeQueryFromFilter(countFilter, q)
	sql, args, err := q.PlaceholderFormat(sq.Question).ToSql()
	if err != nil {
		return 0, err
	}

	var count int64
	err = s.Client.GetContext(s.ctx, &count, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("Failed getting count: %w", err)
	}
	return count, nil
}

func (s *dbStoreImpl) Create(m model.Model, data map[string]interface{}, result interface{}) error {
	q := sq.Insert(m.TableName()).
		SetMap(data).
		PlaceholderFormat(sq.Question)
	r, err := s.GetClient().Exec(s.ctx, q)
	if err != nil {
		return err
	}
	if m.IsAutoIncrementKey() {
		id, err := r.LastInsertId()
		if err != nil {
			return err
		}
		m.SetID(id)
	}
	return s.Get(m, result)
}

func (s *dbStoreImpl) Delete(m model.Model) error {
	key, val := m.PrimaryKey()
	b := sq.Delete(m.TableName()).Where(sq.Eq{key: val})
	_, err := s.GetClient().Exec(s.ctx, b)
	return err
}

func (s *dbStoreImpl) Detach(link model.LinkDef, leftID, rightID int64) error {
	q := sq.Delete(link.TableName).
		Where(sq.Eq{ link.LeftKey: leftID, link.RightKey: rightID }).
		PlaceholderFormat(sq.Question)
	_, err := s.GetClient().Exec(s.ctx, q)
	return err
}

func (s *dbStoreImpl) FindBy(m model.Model, filter repository.Filter, result interface{}) error {
	q := s.Select(m).From(m.TableName())
	q = s.MakeQueryFromFilter(filter, q)
	err := s.Client.Query(s.ctx, q, result)
	if err != nil {
		return err
	}
	return nil
}

func (s *dbStoreImpl) Get(m model.Model, result interface{}) error {
	key, val := m.PrimaryKey()
	q := s.Select(m).
		From(m.TableName()).
		Where(sq.Eq{key: val}).
		Limit(1)
	return s.GetClient().Get(s.ctx, q, result)
}

func (s *dbStoreImpl) GetBy(m model.Model, key string, value any, result interface{}) error {
	q := s.Select(m).
		From(m.TableName()).
		Where(sq.Eq{key: value}).
		Limit(1)
	return s.GetClient().Get(s.ctx, q, result)
}

func (s *dbStoreImpl) GetClient() DBClient {
	return s.Client
}

func (s *dbStoreImpl) GetContext() context.Context {
	return s.ctx
}

func(s *dbStoreImpl) Select(m model.Model) sq.SelectBuilder {
	cols := m.Columns()
	if len(cols) == 0 {
		cols = []string{"*"}
	}
	return sq.Select(cols...)
}

func (s *dbStoreImpl) Update(m model.Model, data map[string]interface{}, result interface{}) error {
	key, val := m.PrimaryKey()
	q := sq.Update(m.TableName()).
		SetMap(data).
		Where(sq.Eq{key: val})
	_, err := s.GetClient().Exec(s.ctx, q)
	if err != nil {
		return err
	}
	return s.Get(m, result)
}

func (s *dbStoreImpl) MakeQueryFromFilter(filter repository.Filter, q sq.SelectBuilder) sq.SelectBuilder {
	if filter.Fields != nil {
		for k, v := range filter.Fields {
			if v == nil || v == "" {
				continue
			}

			if op, ok := filter.Operators[k]; ok && op != "" {
				switch op {
				case "<=":
					q = q.Where(sq.LtOrEq{k: v})
				case ">=":
					q = q.Where(sq.GtOrEq{k: v})
				case "<":
					q = q.Where(sq.Lt{k: v})
				case ">":
					q = q.Where(sq.Gt{k: v})
				case "~=":
					q = q.Where(fmt.Sprintf("MATCH (%s) AGAINST (? IN BOOLEAN MODE)", k), v)
				case "between":
					if vals, ok := v.([]string); ok && len(vals) == 2 {
						q = q.Where(sq.Expr(fmt.Sprintf("%s BETWEEN ? AND ?", k), vals[0], vals[1]))
					}
				case "null":
					q = q.Where(sq.Eq{k: nil})
				case "in":
					fallthrough
				default:
					q = q.Where(sq.Eq{k: v})
				} 
			} else {
				q = q.Where(sq.Eq{k: v})
			}
		}
	}

	if filter.Sort != "" {
		q = q.OrderBy(filter.Sort + " " + filter.Order)
	}

	if filter.Limit > 0 {
		q = q.Limit(uint64(filter.Limit))
		if filter.Offset >= 0 {
			q = q.Offset(uint64(filter.Offset))
		}
	}
	return q
}