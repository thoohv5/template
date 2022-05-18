package app

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uber/jaeger-client-go"

	jc "github.com/uber/jaeger-client-go/config"

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

func New(cf config.IConfig, registerRouter hpx.RegisterRouter) (app.IApp, func(), error) {
	closer, err := initJaeger()
	return &application{
			cf:             cf,
			registerRouter: registerRouter,
		}, func() {
			if closer == nil {
				if err = closer.Close(); nil != err {
					panic(err)
				}
			}
		}, err
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

func initJaeger() (closer io.Closer, err error) {
	// 根据配置初始化Tracer 返回Closer
	tracer, closer, err := (&jc.Configuration{
		ServiceName: "tracing",
		Disabled:    false,
		Sampler: &jc.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &jc.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}).NewTracer()
	if err != nil {
		return
	}

	// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
	opentracing.SetGlobalTracer(tracer)
	return
}
