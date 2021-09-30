package http

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

type Base struct {
	Ctx *gin.Context
}

type BasePage struct {
	options
	Start uint32 `json:"start" form:"start,default=0"`
	Limit uint32 `json:"limit" form:"limit,default=15"`
}

type BaseResp struct {
	HasMore bool        `json:"hasMore"`
	Start   uint32      `json:"start"`
	List    interface{} `json:"list"`
}

type options struct {
	// 0
	defaultStart uint32
	// 50
	maxLimit uint32
}

type Option interface {
	apply(*BasePage)
}

type optionFunc func(*BasePage)

func (f optionFunc) apply(page *BasePage) {
	f(page)
}

func WithDefaultStart(defaultStart uint32) Option {
	return optionFunc(func(page *BasePage) {
		page.defaultStart = defaultStart
	})
}

func WithMaxLimit(maxLimit uint32) Option {
	return optionFunc(func(page *BasePage) {
		page.maxLimit = maxLimit
	})
}

func ParsePage(page BasePage, opts ...Option) BasePage {
	page.options = options{
		defaultStart: 0,
		maxLimit:     50,
	}
	for _, o := range opts {
		o.apply(&page)
	}

	if page.Limit > page.maxLimit {
		page.Limit = page.maxLimit
	}

	page.Limit++

	return page
}

func BuildResponse(list interface{}, page BasePage) *BaseResp {
	var (
		recordNum uint32
		start     uint32
		hasMore   bool
		result    *BaseResp
	)
	v := reflect.ValueOf(list)
	if v.Kind() != reflect.Slice {
		panic("build response fail")
	}

	if 0 == v.Len() {
		list = []interface{}{}
	}

	recordNum = uint32(v.Len())
	if recordNum >= page.Limit {
		start = page.Start + recordNum - 1
		hasMore = true
	} else {
		start = page.Start + recordNum
	}

	if hasMore && page.Limit > 0 {
		list = v.Slice(0, int(page.Limit-1)).Interface()
	}
	result = new(BaseResp)

	result.List = list
	result.Start = start
	result.HasMore = hasMore

	return result
}
