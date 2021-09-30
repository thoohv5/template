package hpx

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	"github.com/thoohv5/template/pkg/hpx/middleware"
	pkgvalidator "github.com/thoohv5/template/pkg/validator"
)

type hpx struct{}

// IHpx Hpx标准
type IHpx interface {
	Handle(registerRouter RegisterRouter) (*gin.Engine, error)
}

type RegisterRouter func(engine *gin.Engine, trans ut.Translator) error

var (
	trans ut.Translator
)

func New() IHpx {
	return new(hpx)
}

// Handle 处理
func (g *hpx) Handle(registerRouter RegisterRouter) (*gin.Engine, error) {
	router := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := pkgvalidator.Register(v, "zh"); nil != err {
			return nil, errors.Wrap(err, "validator Watch err")
		}
	}

	gin.SetMode(gin.DebugMode)

	router.Use(middleware.Recovery(), middleware.Logger(), middleware.Cors())

	if err := registerRouter(router, trans); nil != err {
		return router, err
	}

	return router, nil
}

// ErrResp handler中调用的错误翻译方法
func ErrResp(err error) error {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	eMap := removeStructPre(errs.Translate(trans))
	for _, e := range eMap {
		err = errors.New(fmt.Sprint(e))
	}
	return err
}

func addValueToMap(fields map[string]string) map[string]interface{} {
	res := make(map[string]interface{})
	for field, err := range fields {
		fieldArr := strings.SplitN(field, ".", 2)
		if len(fieldArr) > 1 {
			NewFields := map[string]string{fieldArr[1]: err}
			returnMap := addValueToMap(NewFields)
			if res[fieldArr[0]] != nil {
				for k, v := range returnMap {
					res[fieldArr[0]].(map[string]interface{})[k] = v
				}
			} else {
				res[fieldArr[0]] = returnMap
			}
			continue
		} else {
			res[field] = err
			continue
		}
	}
	return res
}

// 去掉结构体名称前缀
func removeStructPre(fields map[string]string) map[string]interface{} {
	lowerMap := map[string]string{}
	for field, err := range fields {
		fieldArr := strings.SplitN(field, ".", 2)
		lowerMap[fieldArr[1]] = err
	}
	res := addValueToMap(lowerMap)
	return res
}
