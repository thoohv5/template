// Package server 业务
package server

import (
	"github.com/thoohv5/template/internal/service/example"
	"github.com/thoohv5/template/pkg/http"

	"github.com/gin-gonic/gin"

	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/pkg/log"
)

type server struct {
	// 配置
	cf config.IConfig
	// 日志
	log log.ILog

	// 业务
	svr example.IService
}

// IServer 服务标准
type IServer interface {
	GetExample(gtx *gin.Context)
	PostExample(gtx *gin.Context)
}

// New 创建
func New(cf config.IConfig, log log.ILog, svr example.IService) IServer {
	return &server{
		cf:  cf,
		log: log,

		svr: svr,
	}
}

// GetExample 示例
// PingExample godoc
// @Summary 示例
// @Description 获取swagger格式
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param id query int true "id"
// @Param name query string false "姓名"
// @Success 200 {object} http.ResponseEntity{data=GetExampleResp}
// @Router /get-example [get]
func (s *server) GetExample(gtx *gin.Context) {
	resp := http.NewResponse(gtx)

	req := new(GetExampleReq)
	if err := gtx.Bind(req); nil != err {
		resp.Error(err)
		return
	}

	resp.Error(s.svr.Test(gtx, req.Name))
}

// PostExample 示例
// PingExample godoc
// @Summary 示例
// @Description 获取swagger格式
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param req body PostExampleReq true "请求值"
// @Success 200 {object} http.ResponseEntity{data=PostExampleResp}
// @Router /post-example [post]
func (s *server) PostExample(gtx *gin.Context) {
	resp := http.NewResponse(gtx)

	req := new(PostExampleReq)
	if err := gtx.Bind(req); nil != err {
		resp.Error(err)
		return
	}
	ret := new(PostExampleResp)
	ret.Id = req.Id
	ret.Name = req.Name

	resp.DefaultSuccess(ret)
}
