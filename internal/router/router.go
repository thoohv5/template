// Package router 路由
package router

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/google/wire"

	"github.com/thoohv5/template/internal/pkg/app"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/internal/server"
	"github.com/thoohv5/template/pkg/hpx"
)

// ProviderSet is app providers.
var ProviderSet = wire.NewSet(
	New,
)

// New 创建
func New(cf config.IConfig, server server.IServer) hpx.RegisterRouter {
	return func(r *gin.Engine, trans ut.Translator) error {

		if cf.GetHttp().SwaggerEnabled {
			app.InitSwagRouter(r, cf.GetHttp().LocalAddr)
		}

		user := r.Group("/user")

		user.GET("/get-example", server.GetExample)
		user.POST("/post-example", server.PostExample)

		return nil
	}
}
