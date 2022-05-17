package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/thoohv5/template/api/docs"
	"github.com/thoohv5/template/internal/pkg/config"
	"github.com/thoohv5/template/pkg/app"
	"github.com/thoohv5/template/pkg/hpx"
)

// ProviderSet is app providers.
var ProviderSet = wire.NewSet(
	New,
)

type application struct {
	cf             config.IConfig
	registerRouter hpx.RegisterRouter
}

func New(cf config.IConfig, registerRouter hpx.RegisterRouter) app.IApp {
	return &application{
		cf:             cf,
		registerRouter: registerRouter,
	}
}

func (p *application) GetConfig() config.IConfig {
	return p.cf
}

func (p *application) Run(addr ...string) error {
	gen, err := hpx.New().Handle(p.registerRouter)
	if nil != err {
		panic(err)
	}
	_, _ = fmt.Fprintf(os.Stdout, "http://%s\n", p.cf.GetHttp().LocalAddr)
	return gen.Run(addr...)
}

func InitSwagRouter(r *gin.Engine, localAddr string) {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "文档"
	docs.SwaggerInfo.Description = "开发文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = localAddr
	docs.SwaggerInfo.BasePath = "/user"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_, _ = fmt.Fprintf(os.Stdout, "http://%s/swagger/index.html\n", localAddr)
}
