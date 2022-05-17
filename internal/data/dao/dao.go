package dao

import (
	"context"
	"math/rand"
	"time"

	"gorm.io/gorm/schema"

	"github.com/thoohv5/template/pkg/dbx/standard"
)

type (
	IModel = schema.Tabler
	IDao   interface {
		Close()
	}
	dao struct {
		build standard.IBuilder
	}
	options struct {
		build standard.IBuilder
		br    *BaseRequest
		limit int
	}
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (d *dao) Close() {

}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithBaseRequest(br *BaseRequest) Option {
	return optionFunc(func(o *options) { o.br = br })
}

func WithLimit(limit int) Option {
	return optionFunc(func(o *options) { o.limit = limit })
}

func WithBuild(build standard.IBuilder) Option {
	return optionFunc(func(o *options) { o.build = build })
}

func (d *dao) buildFilterQuery(ctx context.Context, opts ...Option) standard.IBuilder {
	optos := options{}
	for _, o := range opts {
		o.apply(&optos)
	}

	return d.baseConn(d.build, optos.br, optos.limit)
}

func (d *dao) baseConn(build standard.IBuilder, condition *BaseRequest, appointLimit int) standard.IBuilder {
	if nil != condition {
		// // Fields
		// if fields := condition.Fields; len(fields) > 0 {
		// 	build = build.Select(fields)
		// }
		// Start
		if start := condition.Start; start > 0 {
			build = build.Offset(start)
		}
		// Limit
		if limit := condition.Limit; limit > 0 {
			build = build.Limit(int(limit))
		}
		// 排序
		for _, sort := range condition.Sorts {
			if string(sort[0]) == "+" {
				sort = string(sort[1:]) + " ASC"
			} else if string(sort[0]) == "-" {
				sort = string(sort[1:]) + " DESC"
			} else {
				sort = sort + " ASC"
			}
			build = build.Order(sort)
		}
		// // GroupBy
		// if groupBy := condition.GroupBy; len(groupBy) > 0 {
		// 	build = build.Group(groupBy)
		// }
	} else {
		build = build.Limit(1)
	}

	if appointLimit > 0 {
		build = build.Limit(appointLimit)
	}

	return build
}

// BaseRequest 公共请求
type BaseRequest struct {
	// 数据开始位置
	Start int `json:"start,omitempty"`
	// 返回数据条数
	Limit int `json:"limit,omitempty"`
	// 字段按照逗号隔开
	Fields string `json:"fields,omitempty"`
	// 排序：sort=otc_type,-created_at,*custom
	// 以符号开头，可选符号：(+或空 正序）（- 倒序）（* 自定义复杂排序标识关键词）
	Sorts []string `json:"sorts,omitempty"`
	// 分组
	GroupBy string `json:"group_by,omitempty"`
}
