package http

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"github.com/thoohv5/template/pkg/hpx"
)

type Response struct {
	gtx      *gin.Context
	basePage BasePage
}

type IResponse interface {
	Success(data interface{}, msg string)
	DefaultSuccess(data interface{})
	Bind(param interface{}, opts ...Option) error
	Error(err error)
}

func NewResponse(gtx *gin.Context) IResponse {
	return &Response{
		gtx: gtx,
	}
}

// Success 响应成功请求
func (r *Response) Success(data interface{}, msg string) {
	if len(msg) == 0 {
		msg = "success"
	}

	if r.basePage.Limit > 0 {
		data = BuildResponse(data, r.basePage)
	}

	// 响应数据
	r.gtx.JSON(http.StatusOK, ResponseEntity{
		ErrorCode: 0,
		Message:   msg,
		Data:      data,
	})
	return
}

func (r *Response) DefaultSuccess(data interface{}) {
	r.Success(data, "成功")
}

// 响应错误
func (r *Response) Error(err error) {
	if err == nil {
		return
	}

	r.gtx.JSON(http.StatusOK, ResponseEntity{
		ErrorCode: 0,
		Message:   hpx.ErrResp(err).Error(),
		Data:      nil,
	})
	return
}

func (r *Response) Bind(param interface{}, opts ...Option) error {
	if err := r.gtx.ShouldBind(param); nil != err {
		return err
	}
	sf := reflect.Indirect(reflect.ValueOf(param)).FieldByName("BasePage")
	if sf.IsValid() {
		r.basePage = ParsePage(BasePage{
			Start: uint32(sf.FieldByName("Start").Uint()),
			Limit: uint32(sf.FieldByName("Limit").Uint()),
		}, opts...)
		sf.Set(reflect.ValueOf(r.basePage))
	}
	return nil
}
