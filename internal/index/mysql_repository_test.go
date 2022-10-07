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

package index_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/bangumi/server/internal/dal/query"
	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/index"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/test"
)

func getRepo(t *testing.T) domain.IndexRepo {
	t.Helper()
	repo, err := index.NewMysqlRepo(query.Use(test.GetGorm(t)), zap.NewNop())
	require.NoError(t, err)

	return repo
}

func TestMysqlRepo_Get(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	i, err := repo.Get(context.Background(), 15045)
	require.NoError(t, err)

	require.EqualValues(t, 15045, i.ID)
	require.EqualValues(t, 14127, i.CreatorID)
	require.False(t, i.NSFW)
}

func TestMysqlRepo_ListSubjects(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	subjects, err := repo.ListSubjects(context.Background(), 15045, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 20)
}

func TestMysqlRepo_NewIndex(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	// 存入的时间戳是 int32 类型， nanosecond 会被忽略掉
	// TODO: 数据库时间戳是否应该改成 uint32 或者 uint64 类型
	now := time.Now()

	index := &model.Index{
		ID:          0,
		Title:       "test",
		Description: "Test Index",
		CreatorID:   382951,
		CreatedAt:   now,
		UpdateAt:    now,
		Total:       0,
		Comments:    0,
		Collects:    0,
		Ban:         false,
		NSFW:        false,
	}
	err := repo.New(context.Background(), index)
	require.NoError(t, err)
	require.NotEqualValues(t, 0, index.ID)

	i, err := repo.Get(context.Background(), index.ID)
	require.NoError(t, err)

	require.EqualValues(t, 382951, i.CreatorID)
	require.EqualValues(t, "test", i.Title)
	require.EqualValues(t, now.Truncate(time.Second), i.CreatedAt)
	require.EqualValues(t, now.Truncate(time.Second), i.UpdateAt)
}

func TestMysqlRepo_UpdateIndex(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)
	_, err := repo.Get(context.Background(), 15045)

	defaultTitle := "日本动画最高收视率TOP100"
	defaultDesc := "[url]http://www.tudou.com/programs/view/" +
		"W6eIoxnHs6g/[/url]\r\n有美国动画混入，所以准确的说是在日本播放的" +
		"动画最高收视率（而且是关东地区的\r\n基本大部分是70年代的，那个年代娱乐贫乏优势真大"

	// check current
	require.NoError(t, err)

	// update index information
	err = repo.Update(context.Background(), 15045, "日本动画", "日本动画的介绍")
	require.NoError(t, err)

	// check updated index information
	i, err := repo.Get(context.Background(), 15045)
	require.NoError(t, err)
	require.EqualValues(t, "日本动画", i.Title)
	require.EqualValues(t, "日本动画的介绍", i.Description)

	// revert update
	err = repo.Update(context.Background(), 15045, defaultTitle, defaultDesc)
	require.NoError(t, err)
}

func TestMysqlRepo_DeleteIndex(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)
	index := &model.Index{
		ID:          0,
		Title:       "test",
		Description: "Test Index",
		CreatorID:   382951,
		CreatedAt:   time.Now(),
		Total:       0,
		Comments:    0,
		Collects:    0,
		Ban:         false,
		NSFW:        false,
	}
	_ = repo.New(context.Background(), index)
	require.NotEqual(t, 0, index.ID)

	i, err := repo.Get(context.Background(), index.ID)
	require.NoError(t, err)
	require.EqualValues(t, index.ID, i.ID)

	err = repo.Delete(context.Background(), index.ID)
	require.NoError(t, err)

	_, err = repo.Get(context.Background(), index.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrNotFound)
}

// 删除目录会把所属的 subject 全部删掉
func TestMysqlRepo_DeleteIndex2(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)
	index := &model.Index{
		ID:          0,
		Title:       "test",
		Description: "Test Index",
		CreatorID:   382951,
		CreatedAt:   time.Now(),
		Total:       0,
		Comments:    0,
		Collects:    0,
		Ban:         false,
		NSFW:        false,
	}

	ctx := context.Background()

	err := repo.New(ctx, index)
	require.NoError(t, err)

	for i := 10; i < 20; i++ {
		_, err = repo.AddIndexSubject(ctx, index.ID, model.SubjectID(i),
			uint32(i), fmt.Sprintf("comment %d", i))
		require.NoError(t, err)
	}

	i, err := repo.Get(ctx, index.ID)
	require.NoError(t, err)
	require.EqualValues(t, 10, i.Total)

	subjects, err := repo.ListSubjects(context.Background(), index.ID, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 10)

	err = repo.Delete(ctx, index.ID)
	require.NoError(t, err)

	i, err = repo.Get(ctx, index.ID)
	require.Equal(t, err, domain.ErrNotFound)

	subjects, err = repo.ListSubjects(context.Background(), index.ID, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 0)

	// 确保不会影响到其他目录
	subjects, err = repo.ListSubjects(context.Background(), 15045, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 20)
}

func TestMysqlRepo_AddIndexSubject(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)
	index := &model.Index{
		ID:          0,
		Title:       "test",
		Description: "Test Index",
		CreatorID:   382951,
		CreatedAt:   time.Now(),
		Total:       0,
		Comments:    0,
		Collects:    0,
		Ban:         false,
		NSFW:        false,
	}

	ctx := context.Background()

	err := repo.New(ctx, index)
	require.NotEqual(t, 0, index.ID)
	require.NoError(t, err)

	_, err = repo.AddIndexSubject(ctx, index.ID, 3, 1, "comment 1")
	require.NoError(t, err)

	_, err = repo.AddIndexSubject(ctx, index.ID, 4, 3, "comment 3")
	require.NoError(t, err)

	i, err := repo.Get(ctx, index.ID)
	require.NoError(t, err)
	require.EqualValues(t, index.ID, i.ID)

	require.EqualValues(t, 2, i.Total)

	subjects, err := repo.ListSubjects(context.Background(), index.ID, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 2)

	cache := map[model.SubjectID]domain.IndexSubject{}
	for _, s := range subjects {
		cache[s.Subject.ID] = s
	}
	require.EqualValues(t, cache[3].Comment, "comment 1")
	require.EqualValues(t, cache[4].Comment, "comment 3")

	err = repo.Delete(ctx, index.ID)
	require.NoError(t, err)
}

func TestMysqlRepo_DeleteIndexSubject(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)
	index := &model.Index{
		ID:          0,
		Title:       "test",
		Description: "Test Index",
		CreatorID:   382951,
		CreatedAt:   time.Now(),
		Total:       0,
		Comments:    0,
		Collects:    0,
		Ban:         false,
		NSFW:        false,
	}

	ctx := context.Background()

	err := repo.New(ctx, index)
	require.NotEqual(t, 0, index.ID)
	require.NoError(t, err)

	for i := 10; i < 20; i++ {
		_, err = repo.AddIndexSubject(ctx, index.ID, model.SubjectID(i),
			uint32(i), fmt.Sprintf("comment %d", i))
		require.NoError(t, err)
	}

	i, err := repo.Get(ctx, index.ID)
	require.NoError(t, err)
	require.EqualValues(t, 10, i.Total)

	subjects, err := repo.ListSubjects(context.Background(), index.ID, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 10)

	err = repo.DeleteIndexSubject(ctx, index.ID, 15)
	require.NoError(t, err)

	i, err = repo.Get(ctx, index.ID)
	require.NoError(t, err)
	require.EqualValues(t, 9, i.Total)

	subjects, err = repo.ListSubjects(context.Background(), index.ID, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 9)

	for _, v := range subjects {
		require.NotEqualValues(t, v.Subject.ID, 15)
	}

	err = repo.Delete(ctx, index.ID)
	require.NoError(t, err)
}

func TestMysqlRepo_DeleteNonExistsIndexSubject(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	ctx := context.Background()

	_ = repo.Delete(ctx, 99999999)

	err := repo.DeleteIndexSubject(ctx, 99999999, 15)
	require.Error(t, err)
	require.Error(t, err, domain.ErrNotFound)
}

func TestMysqlRepo_FailedAddedToNonExists(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	ctx := context.Background()
	_ = repo.Delete(ctx, 99999999) // in case index(99999999) exists

	_, err := repo.AddIndexSubject(ctx, 99999999, 5, 5, "test")
	require.Error(t, err)
	require.Equal(t, err, domain.ErrNotFound)
}

func TestMysqlRepo_UpdateSubjectInfo(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)
	index := &model.Index{
		ID:          0,
		Title:       "test",
		Description: "Test Index",
		CreatorID:   382951,
		CreatedAt:   time.Now(),
		Total:       0,
		Comments:    0,
		Collects:    0,
		Ban:         false,
		NSFW:        false,
	}
	ctx := context.Background()

	err := repo.New(ctx, index)
	require.NoError(t, err)

	_, err = repo.AddIndexSubject(ctx, index.ID, 5, 5, "test")
	require.NoError(t, err)
	subjects, err := repo.ListSubjects(context.Background(), index.ID, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 1)
	require.EqualValues(t, subjects[0].Subject.ID, 5)
	require.EqualValues(t, subjects[0].Comment, "test")

	err = repo.UpdateIndexSubject(ctx, index.ID, 5, 5, "test22222")
	require.NoError(t, err)

	subjects, err = repo.ListSubjects(context.Background(), index.ID, model.SubjectTypeAll, 20, 0)
	require.NoError(t, err)
	require.Len(t, subjects, 1)
	require.EqualValues(t, subjects[0].Subject.ID, 5)
	require.EqualValues(t, subjects[0].Comment, "test22222")
}

func TestMysqlRepo_AddExists(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	index := &model.Index{
		ID:          0,
		Title:       "test",
		Description: "Test Index",
		CreatorID:   382951,
		CreatedAt:   time.Now(),
		Total:       0,
		Comments:    0,
		Collects:    0,
		Ban:         false,
		NSFW:        false,
	}
	ctx := context.Background()

	_ = repo.New(ctx, index)

	_, err := repo.AddIndexSubject(ctx, index.ID, 5, 5, "test")
	require.NoError(t, err)

	_, err = repo.AddIndexSubject(ctx, index.ID, 5, 5, "test")
	require.Error(t, err)
	require.Equal(t, err, domain.ErrExists)
}

func TestMysqlRepo_AddNoneExistsSubject(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	ctx := context.Background()

	_, err := repo.AddIndexSubject(ctx, 15045, 999999999, 5, "test")
	require.Error(t, err)
	require.Equal(t, err, domain.ErrSubjectNotFound)
}
