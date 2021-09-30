package validator

import (
	"errors"
	"fmt"
	"reflect"

	ven "github.com/go-playground/locales/en"
	vzh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	playgroundvalidator "github.com/go-playground/validator/v10"
	translationsen "github.com/go-playground/validator/v10/translations/en"
	translationszh "github.com/go-playground/validator/v10/translations/zh"
)

func Register(v *playgroundvalidator.Validate, local string) (err error) {

	zh := vzh.New()
	cn := ven.New()
	uni := ut.New(cn, cn, zh)
	trans, ok := uni.GetTranslator(local)
	if !ok {
		return errors.New("GetTranslator err")
	}

	// 注册提示语言
	switch local {
	case "zh":
		err = translationszh.RegisterDefaultTranslations(v, trans)
	case "en":
		err = translationsen.RegisterDefaultTranslations(v, trans)
	default:
		err = translationsen.RegisterDefaultTranslations(v, trans)
	}
	if nil != err {
		return fmt.Errorf("RegisterDefaultTranslations err:%w", err)
	}

	// 注册字段名称解释
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})

	// 注册检查规则
	if err = registerValidation(v); nil != err {
		return err
	}

	// 注册字段名称解释
	if err := registerTranslation(v, trans); nil != err {
		return err
	}

	return
}

// registerValidation 注册检查规则
func registerValidation(v *playgroundvalidator.Validate) (err error) {

	// 检查规则test
	err = v.RegisterValidation("test", func(fl playgroundvalidator.FieldLevel) bool {
		str, ok := fl.Field().Interface().(string)
		if ok {
			if str != "test" {
				return false
			}
		}
		return true
	})
	if nil != err {
		return fmt.Errorf("registerValidation err:%w", err)
	}

	return
}

// registerTranslation 注册字段提示
func registerTranslation(v *playgroundvalidator.Validate, trans ut.Translator) (err error) {

	// 注册字段提示
	err = v.RegisterTranslation("test", trans, func(ut ut.Translator) error {
		return ut.Add("test", "{0}测试检查", true)
	}, func(ut ut.Translator, fe playgroundvalidator.FieldError) string {
		t, _ := ut.T("test", fe.Field())
		return t
	})
	if nil != err {
		return fmt.Errorf("registerTranslation err:%w", err)
	}

	return
}
