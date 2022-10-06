// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package index

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/bangumi/server/internal/dal/dao"
	"github.com/bangumi/server/internal/dal/query"
	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/errgo"
	"github.com/bangumi/server/internal/subject"
)

func NewMysqlRepo(q *query.Query, log *zap.Logger) (domain.IndexRepo, error) {
	return mysqlRepo{q: q, log: log.Named("index.mysqlRepo")}, nil
}

type mysqlRepo struct {
	q   *query.Query
	log *zap.Logger
}

func (r mysqlRepo) isNsfw(ctx context.Context, id uint32) (bool, error) {
	i, err := r.q.IndexSubject.WithContext(ctx).
		Join(r.q.Subject, r.q.IndexSubject.Sid.EqCol(r.q.Subject.ID)).
		Where(r.q.IndexSubject.Rid.Eq(id), r.q.Subject.Nsfw.Is(true)).Count()
	if err != nil {
		r.log.Error("unexpected error when checking index nsfw", zap.Uint32("index_id", id))
		return false, errgo.Wrap(err, "dal")
	}

	return i != 0, nil
}

func (r mysqlRepo) Get(ctx context.Context, id uint32) (model.Index, error) {
	i, err := r.q.Index.WithContext(ctx).
		Where(r.q.Index.ID.Eq(id), r.q.Index.Ban.Is(false)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Index{}, domain.ErrNotFound
		}

		return model.Index{}, errgo.Wrap(err, "dal")
	}

	nsfw, err := r.isNsfw(ctx, id)
	if err != nil {
		return model.Index{}, err
	}

	ret := daoToModel(i)
	ret.NSFW = nsfw
	return *ret, nil
}

func (r mysqlRepo) CountSubjects(
	ctx context.Context, id uint32, subjectType model.SubjectType,
) (int64, error) {
	q := r.q.IndexSubject.WithContext(ctx).Where(r.q.IndexSubject.Rid.Eq(id))
	if subjectType != 0 {
		q = q.Where(r.q.IndexSubject.Type.Eq(subjectType))
	}

	i, err := q.Count()
	if err != nil {
		return 0, errgo.Wrap(err, "dal")
	}

	return i, nil
}

func (r mysqlRepo) ListSubjects(
	ctx context.Context,
	id uint32,
	subjectType model.SubjectType,
	limit, offset int,
) ([]domain.IndexSubject, error) {
	q := r.q.IndexSubject.WithContext(ctx).Joins(r.q.IndexSubject.Subject).Preload(r.q.IndexSubject.Subject.Fields).
		Where(r.q.IndexSubject.Rid.Eq(id)).
		Order(r.q.IndexSubject.Order).Limit(limit).Offset(offset)
	if subjectType != 0 {
		q = q.Where(r.q.IndexSubject.Type.Eq(subjectType))
	}

	d, err := q.Find()
	if err != nil {
		return nil, errgo.Wrap(err, "dal")
	}

	var results = make([]domain.IndexSubject, len(d))
	for i, s := range d {
		sub, err := subject.ConvertDao(&s.Subject)
		if err != nil {
			return nil, errgo.Wrap(err, "subject.ConvertDao")
		}

		results[i] = domain.IndexSubject{
			AddedAt: time.Unix(int64(s.Dateline), 0),
			Comment: s.Comment,
			Subject: sub,
		}
	}

	return results, nil
}

func (r mysqlRepo) New(ctx context.Context, i *model.Index) error {
	dao := modelToDAO(i)
	if err := r.q.Index.WithContext(ctx).Create(dao); err != nil {
		return errgo.Wrap(err, "failed to create index in db")
	}
	i.ID = dao.ID
	return nil
}

func (r mysqlRepo) Update(ctx context.Context, id uint32, title string, desc string) error {
	query := r.q.Index.WithContext(ctx)
	result, err := query.Where(r.q.Index.ID.Eq(id)).Updates(dao.Index{
		Title: title,
		Desc:  desc,
	})
	if err != nil {
		return errgo.Wrap(err, "failed to update index info")
	}
	if err = result.Error; err != nil {
		return errgo.Wrap(err, "failed to update index info")
	}
	return nil
}

func daoToModel(index *dao.Index) *model.Index {
	return &model.Index{
		ID:          index.ID,
		Title:       index.Title,
		Description: index.Desc,
		CreatorID:   index.CreatorID,
		Total:       index.SubjectTotal,
		Comments:    index.Replies,
		Collects:    index.Collects,
		Ban:         index.Ban,
		NSFW:        false, // check nsfw outside of this function
		CreatedAt:   time.Unix(int64(index.Dateline), 0),
	}
}

func modelToDAO(index *model.Index) *dao.Index {
	return &dao.Index{
		ID:        index.ID,
		Type:      0,
		Title:     index.Title,
		Desc:      index.Description,
		CreatorID: index.CreatorID,
		Ban:       index.Ban,
		Dateline:  int32(index.CreatedAt.Unix()),
	}
}
