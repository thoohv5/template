package dao

import (
	"context"

	"github.com/thoohv5/template/internal/data/dao/model"
	ds "github.com/thoohv5/template/internal/data/standard"
	dbs "github.com/thoohv5/template/pkg/dbx/standard"
)

type (
	ITest interface {
		// Detail 查询一条
		Detail(ctx context.Context, condition *TestCondition, opts ...Option) (info *model.Test, err error)
		// List 查询多条
		List(ctx context.Context, condition *TestCondition, opts ...Option) (list []*model.Test, err error)
		// Add 添加
		Add(ctx context.Context, data *model.Test, opts ...Option) (err error)
		// Update 更新
		Update(ctx context.Context, condition *TestCondition, updateData map[string]interface{}, opts ...Option) (err error)
	}
	test struct {
		dao
	}
	TestCondition struct {
		*BaseRequest

		// 查询参数
		Id  int32
		Ids []int32
	}
)

// NewTest 创建
func NewTest(data ds.IData) ITest {
	m := new(test)
	m.build = data.GetEdb().Model(&model.Test{})
	return m
}

// Detail 查询一条
func (m *test) Detail(ctx context.Context, condition *TestCondition, opts ...Option) (info *model.Test, err error) {
	info = new(model.Test)
	return info, m.buildFilterQuery(ctx, condition, opts...).Find(ctx, info)
}

// List 查询多条
func (m *test) List(ctx context.Context, condition *TestCondition, opts ...Option) (list []*model.Test, err error) {
	list = make([]*model.Test, 0)
	return list, m.buildFilterQuery(ctx, condition, opts...).Get(ctx, list)
}

// Add 添加
func (m *test) Add(ctx context.Context, data *model.Test, opts ...Option) (err error) {
	return m.buildFilterQuery(ctx, &TestCondition{}, opts...).Add(ctx, data)
}

// Update 更新
func (m *test) Update(ctx context.Context, condition *TestCondition, updateData map[string]interface{}, opts ...Option) (err error) {
	return m.buildFilterQuery(ctx, condition, opts...).Update(ctx, updateData)
}

// 参数构建
func (m *test) buildFilterQuery(ctx context.Context, condition *TestCondition, opts ...Option) (build dbs.IBuilder) {
	// 处理公共参数
	if condition != nil && condition.BaseRequest != nil {
		opts = append(opts, WithBaseRequest(condition.BaseRequest))
	}
	// 预处理
	opts = append(opts, WithBuild(m.build))
	build = m.dao.buildFilterQuery(ctx, opts...)

	if condition == nil {
		return
	}

	// 查询判断
	ids := condition.Ids
	if id := condition.Id; id > 0 {
		ids = append(ids, id)
	}
	if len(ids) > 0 {
		if len(ids) == 1 {
			build = build.Where("id = ?", ids[0])
		} else {
			build = build.Where("id IN(?)", ids)
		}
	}

	return
}
