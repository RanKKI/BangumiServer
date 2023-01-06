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

package collection

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/trim21/go-phpserialize"
	"go.uber.org/zap"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"github.com/bangumi/server/dal/dao"
	"github.com/bangumi/server/dal/query"
	"github.com/bangumi/server/dal/utiltype"
	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/errgo"
	"github.com/bangumi/server/internal/pkg/gstr"
	"github.com/bangumi/server/internal/subject"
)

var _ Repo = mysqlRepo{}

func NewMysqlRepo(q *query.Query, log *zap.Logger) (Repo, error) {
	return mysqlRepo{
		q:   q,
		log: log.Named("collection.mysqlRepo"),
	}, nil
}

type mysqlRepo struct {
	q   *query.Query
	log *zap.Logger
}

func (r mysqlRepo) WithQuery(query *query.Query) Repo {
	return mysqlRepo{q: query, log: r.log}
}

func (r mysqlRepo) CountSubjectCollections(
	ctx context.Context,
	userID model.UserID,
	subjectType model.SubjectType,
	collectionType model.SubjectCollection,
	showPrivate bool,
) (int64, error) {
	q := r.q.SubjectCollection.WithContext(ctx).
		Where(r.q.SubjectCollection.UserID.Eq(userID))

	if subjectType != model.SubjectTypeAll {
		q = q.Where(r.q.SubjectCollection.SubjectType.Eq(subjectType))
	}

	if collectionType != model.SubjectCollectionAll {
		q = q.Where(r.q.SubjectCollection.Type.Eq(uint8(collectionType)))
	}

	if !showPrivate {
		q = q.Where(r.q.SubjectCollection.Private.Eq(uint8(model.CollectPrivacyNone)))
	}

	c, err := q.Count()
	if err != nil {
		return 0, errgo.Wrap(err, "dal")
	}

	return c, nil
}

func (r mysqlRepo) ListSubjectCollection(
	ctx context.Context,
	userID model.UserID,
	subjectType model.SubjectType,
	collectionType model.SubjectCollection,
	showPrivate bool,
	limit, offset int,
) ([]model.UserSubjectCollection, error) {
	q := r.q.SubjectCollection.WithContext(ctx).
		Order(r.q.SubjectCollection.UpdatedTime.Desc()).
		Where(r.q.SubjectCollection.UserID.Eq(userID)).Limit(limit).Offset(offset)

	if subjectType != model.SubjectTypeAll {
		q = q.Where(r.q.SubjectCollection.SubjectType.Eq(subjectType))
	}

	if collectionType != model.SubjectCollectionAll {
		q = q.Where(r.q.SubjectCollection.Type.Eq(uint8(collectionType)))
	}

	if !showPrivate {
		q = q.Where(r.q.SubjectCollection.Private.Eq(uint8(model.CollectPrivacyNone)))
	}

	collections, err := q.Find()
	if err != nil {
		r.log.Error("unexpected error happened", zap.Error(err))
		return nil, errgo.Wrap(err, "dal")
	}

	var results = make([]model.UserSubjectCollection, len(collections))
	for i, c := range collections {
		results[i] = model.UserSubjectCollection{
			UpdatedAt:   time.Unix(int64(c.UpdatedTime), 0),
			Comment:     string(c.Comment),
			Tags:        gstr.Split(c.Tag, " "),
			SubjectType: c.SubjectType,
			Rate:        c.Rate,
			SubjectID:   c.SubjectID,
			EpStatus:    c.EpStatus,
			VolStatus:   c.VolStatus,
			Type:        model.SubjectCollection(c.Type),
			Private:     model.CollectPrivacy(c.Private) != model.CollectPrivacyNone,
		}
	}

	return results, nil
}

func (r mysqlRepo) GetSubjectCollection(
	ctx context.Context, userID model.UserID, subjectID model.SubjectID,
) (model.UserSubjectCollection, error) {
	c, err := r.q.SubjectCollection.WithContext(ctx).
		Where(r.q.SubjectCollection.UserID.Eq(userID), r.q.SubjectCollection.SubjectID.Eq(subjectID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserSubjectCollection{}, domain.ErrSubjectNotCollected
		}

		return model.UserSubjectCollection{}, errgo.Wrap(err, "dal")
	}

	return model.UserSubjectCollection{
		UpdatedAt:   time.Unix(int64(c.UpdatedTime), 0),
		Comment:     string(c.Comment),
		Tags:        gstr.Split(c.Tag, " "),
		SubjectType: c.SubjectType,
		Rate:        c.Rate,
		SubjectID:   c.SubjectID,
		EpStatus:    c.EpStatus,
		VolStatus:   c.VolStatus,
		Type:        model.SubjectCollection(c.Type),
		Private:     model.CollectPrivacy(c.Private) != model.CollectPrivacyNone,
	}, nil
}

func (r mysqlRepo) GetSubjectEpisodesCollection(
	ctx context.Context,
	userID model.UserID,
	subjectID model.SubjectID,
) (model.UserSubjectEpisodesCollection, error) {
	d, err := r.q.EpCollection.WithContext(ctx).Where(
		r.q.EpCollection.UserID.Eq(userID),
		r.q.EpCollection.SubjectID.Eq(subjectID),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserSubjectEpisodesCollection{}, nil
		}
		return nil, errgo.Wrap(err, "query.EpCollection.Find")
	}

	e, err := deserializePhpEpStatus(d.Status)
	if err != nil {
		return nil, err
	}

	return e.toModel(), nil
}

func (r mysqlRepo) UpdateSubjectCollection(
	ctx context.Context,
	userID model.UserID,
	subjectID model.SubjectID,
	data Update,
	at time.Time,
) error {
	t := r.q.SubjectCollection
	old, err := t.WithContext(ctx).Where(t.SubjectID.Eq(subjectID), t.UserID.Eq(userID)).First()
	if err != nil {
		return errgo.Wrap(err, "failed to get old collection record")
	}

	var updater = make([]field.AssignExpr, 0, 11)
	updater = append(updater, t.UpdatedTime.Value(uint32(at.Unix())), t.LastUpdateIP.Value(data.IP))

	if data.Comment.Set {
		updater = append(updater,
			t.Comment.Value(utiltype.HTMLEscapedString(data.Comment.Value)),
			t.HasComment.Value(data.Comment.Value != ""))
	}
	if data.Tags != nil {
		updater = append(updater, t.Tag.Value(strings.Join(data.Tags, " ")))
	}
	if data.EpStatus.Set {
		updater = append(updater, t.EpStatus.Value(data.EpStatus.Value))
	}
	if data.VolStatus.Set {
		updater = append(updater, t.VolStatus.Value(data.VolStatus.Value))
	}
	if data.Privacy.Set {
		updater = append(updater, t.Private.Value(uint8(data.Privacy.Value)))
	}
	if data.Rate.Set {
		updater = append(updater, t.Rate.Value(data.Rate.Value))
	}

	if data.Type.Set {
		updater = append(updater, t.Type.Value(uint8(data.Type.Value)))
		if uint8(data.Type.Value) != old.Type {
			u, e := r.subjectCollectionUpdater(data.Type.Value, at)
			if e != nil {
				return e
			}
			updater = append(updater, u)
		}
	}

	_, err = t.WithContext(ctx).Where(t.SubjectID.Eq(subjectID), t.UserID.Eq(userID)).UpdateSimple(updater...)
	if err != nil {
		return errgo.Wrap(err, "SubjectCollection.Update")
	}

	r.updateSubject(ctx, subjectID)

	return nil
}

func (r mysqlRepo) updateSubject(ctx context.Context, subjectID model.SubjectID) {
	if err := r.updateSubjectTags(ctx, subjectID); err != nil {
		r.log.Error("failed to update subject tags", zap.Error(err))
	}

	if err := r.reCountSubjectCollection(ctx, subjectID); err != nil {
		r.log.Error("failed to update collection counts", zap.Error(err))
	}
}

func (r mysqlRepo) reCountSubjectCollection(ctx context.Context, subjectID model.SubjectID) error {
	var counts []struct {
		Type  uint8  `gorm:"type"`
		Total uint32 `gorm:"total"`
	}

	err := r.q.SubjectCollection.WithContext(ctx).
		Select(r.q.SubjectCollection.Type.As("type"), r.q.SubjectCollection.Type.Count().As("total")).
		Group(r.q.SubjectCollection.Type).
		Where(r.q.SubjectCollection.SubjectID.Eq(subjectID)).Group(r.q.SubjectCollection.Type).Scan(&counts)
	if err != nil {
		return errgo.Wrap(err, "dal")
	}

	var updater = make([]field.AssignExpr, 0, 5)

	for _, count := range counts {
		switch model.SubjectCollection(count.Type) { //nolint:exhaustive
		case model.SubjectCollectionDropped:
			updater = append(updater, r.q.Subject.Dropped.Value(count.Total))

		case model.SubjectCollectionWish:
			updater = append(updater, r.q.Subject.Wish.Value(count.Total))

		case model.SubjectCollectionDoing:
			updater = append(updater, r.q.Subject.Doing.Value(count.Total))

		case model.SubjectCollectionOnHold:
			updater = append(updater, r.q.Subject.OnHold.Value(count.Total))

		case model.SubjectCollectionDone:
			updater = append(updater, r.q.Subject.Done.Value(count.Total))
		}
	}

	_, err = r.q.Subject.WithContext(ctx).Where(r.q.Subject.ID.Eq(subjectID)).UpdateSimple(updater...)
	if err != nil {
		return errgo.Wrap(err, "dal")
	}

	return nil
}

func (r mysqlRepo) updateSubjectTags(ctx context.Context, subjectID model.SubjectID) error {
	collections, err := r.q.SubjectCollection.WithContext(ctx).
		Where(
			r.q.SubjectCollection.SubjectID.Eq(subjectID),
			r.q.SubjectCollection.Private.Neq(uint8(model.CollectPrivacyBan)),
		).Find()
	if err != nil {
		return errgo.Wrap(err, "failed to get all collection")
	}

	var tags = make(map[string]int)
	for _, collection := range collections {
		for _, s := range strings.Split(collection.Tag, " ") {
			if s == "" {
				continue
			}
			tags[s]++
		}
	}

	var phpTags = make([]subject.Tag, 0, len(tags))

	for name, count := range tags {
		name := name
		phpTags = append(phpTags, subject.Tag{
			Name:  &name,
			Count: count,
		})
	}

	sort.Slice(phpTags, func(i, j int) bool {
		if phpTags[i].Count != phpTags[j].Count {
			return phpTags[i].Count > phpTags[j].Count
		}

		return *phpTags[i].Name > *phpTags[j].Name
	})

	newTag, err := phpserialize.Marshal(lo.Slice(phpTags, 0, 30)) //nolint:gomnd
	if err != nil {
		return errgo.Wrap(err, "php.Marshal")
	}

	_, err = r.q.SubjectField.WithContext(ctx).Where(r.q.SubjectField.Sid.Eq(subjectID)).
		UpdateSimple(r.q.SubjectField.Tags.Value(newTag))

	return errgo.Wrap(err, "failed to update subject field")
}

func (r mysqlRepo) subjectCollectionUpdater(t model.SubjectCollection, at time.Time) (field.AssignExpr, error) {
	switch t {
	case model.SubjectCollectionAll:
		return nil, errgo.Wrap(domain.ErrInput, "can't set collection type to SubjectCollectionAll")
	case model.SubjectCollectionWish:
		return r.q.SubjectCollection.WishTime.Value(uint32(at.Unix())), nil
	case model.SubjectCollectionDone:
		return r.q.SubjectCollection.DoneTime.Value(uint32(at.Unix())), nil
	case model.SubjectCollectionDoing:
		return r.q.SubjectCollection.DoingTime.Value(uint32(at.Unix())), nil
	case model.SubjectCollectionDropped:
		return r.q.SubjectCollection.DroppedTime.Value(uint32(at.Unix())), nil
	case model.SubjectCollectionOnHold:
		return r.q.SubjectCollection.OnHoldTime.Value(uint32(at.Unix())), nil
	}

	return nil, errgo.Wrap(domain.ErrInput, fmt.Sprintln("invalid subject collection type", t))
}

func (r mysqlRepo) UpdateEpisodeCollection(
	ctx context.Context,
	userID model.UserID,
	subjectID model.SubjectID,
	episodeIDs []model.EpisodeID,
	collectionType model.EpisodeCollection,
	at time.Time,
) (model.UserSubjectEpisodesCollection, error) {
	table := r.q.EpCollection
	where := []gen.Condition{table.UserID.Eq(userID), table.SubjectID.Eq(subjectID)}

	d, err := table.WithContext(ctx).Where(where...).First()
	if err != nil {
		// 章节表在用到时才会创建
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.createEpisodeCollection(ctx, userID, subjectID, episodeIDs, collectionType, at)
		}

		r.log.Error("failed to get episode collection record", zap.Error(err))
		return nil, errgo.Wrap(err, "dal")
	}

	e, err := deserializePhpEpStatus(d.Status)
	if err != nil {
		r.log.Error("failed to deserialize php-serialized bytes to go data", zap.Error(err))
		return nil, err
	}

	if updated := updateMysqlEpisodeCollection(e, episodeIDs, collectionType); !updated {
		return e.toModel(), nil
	}

	bytes, err := serializePhpEpStatus(e)
	if err != nil {
		return nil, err
	}

	_, err = table.WithContext(ctx).Where(where...).
		UpdateColumnSimple(table.Status.Value(bytes), table.UpdatedTime.Value(uint32(at.Unix())))
	if err != nil {
		return nil, errgo.Wrap(err, "EpCollection.UpdateColumnSimple")
	}

	return e.toModel(), nil
}

func (r mysqlRepo) createEpisodeCollection(
	ctx context.Context,
	userID model.UserID,
	subjectID model.SubjectID,
	episodeIDs []model.EpisodeID,
	collectionType model.EpisodeCollection,
	at time.Time,
) (model.UserSubjectEpisodesCollection, error) {
	var e = make(mysqlEpCollection, len(episodeIDs))
	updateMysqlEpisodeCollection(e, episodeIDs, collectionType)

	bytes, err := serializePhpEpStatus(e)
	if err != nil {
		return nil, err
	}

	table := r.q.EpCollection
	err = table.WithContext(ctx).Where(table.UserID.Eq(userID), table.SubjectID.Eq(subjectID)).Create(&dao.EpCollection{
		UserID:      userID,
		SubjectID:   subjectID,
		Status:      bytes,
		UpdatedTime: uint32(at.Unix()),
	})
	if err != nil {
		r.log.Error("failed to create episode collection record", zap.Error(err))
		return nil, errgo.Wrap(err, "dal")
	}

	return e.toModel(), nil
}

func updateMysqlEpisodeCollection(
	e mysqlEpCollection,
	episodeIDs []model.EpisodeID,
	collectionType model.EpisodeCollection,
) bool {
	var updated bool

	if collectionType == model.EpisodeCollectionNone {
		// remove episode collection
		for _, episodeID := range episodeIDs {
			_, ok := e[episodeID]
			if !ok {
				continue
			}

			delete(e, episodeID)
			updated = true
		}
	} else {
		for _, episodeID := range episodeIDs {
			v, ok := e[episodeID]
			if ok && v.Type == collectionType {
				continue
			}

			e[episodeID] = mysqlEpCollectionItem{EpisodeID: episodeID, Type: collectionType}
			updated = true
		}
	}

	return updated
}
