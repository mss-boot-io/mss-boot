/*
 * @Author: lwnmengjing
 * @Date: 2021/6/22 4:48 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/22 4:48 下午
 */

package response

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mss-boot-io/mss-boot/core/errcode"
	"github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/language"
)

var DefaultLanguage = "zh-CN"

type Api struct {
	Context *gin.Context
	Logger  *logger.Helper
	Error   error
	engine  *gin.RouterGroup
}

func (e *Api) SetEngine(engine *gin.RouterGroup) {
	e.engine = engine
}

func (e *Api) AddError(err error) {
	if e.Error == nil {
		e.Error = err
	} else if err != nil {
		e.Logger.Error(err)
		e.Error = fmt.Errorf("%v; %w", e.Error, err)
	}
}

// Make 设置http上下文
func (e *Api) Make(c *gin.Context) *Api {
	e.Context = c
	e.Logger = GetRequestLogger(c)
	return e
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	switch e.Context.Request.Method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		//去除json、yaml、xml
		for i := range bindings {
			switch bindings[i] {
			case binding.JSON, binding.XML, binding.YAML:
				bindings = append(bindings[:i], bindings[i:]...)
			}
		}
	}
	needValidateNum := len(bindings) - 1
	for i := range bindings {
		switch bindings[i] {
		case nil:
			err = e.Context.ShouldBindUri(d)
		case binding.Query:
			err = e.Context.BindQuery(d)
		default:
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			e.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			errs, ok := err.(validator.ValidationErrors)
			if ok && i < needValidateNum {
				err = nil
				continue
			}
			if err != nil {
				e.AddError(err)
				return e
			}
			trans, errT := transInit(e.getAcceptLanguage())
			if errT != nil {
				err = fmt.Errorf(errT.Error()+", %w", err)
				e.AddError(err)
				return e
			}
			validatorErrs := errs.Translate(trans)
			strArr := make([]string, 0)
			for k, v := range validatorErrs {
				strArr = append(strArr, k+":"+v)
			}
			if len(strArr) != 0 {
				err = errors.New(strings.Join(strArr, ","))
				e.AddError(err)
				return e
			}
		}
	}
	return e
}

// Err 通常错误数据处理
func (e Api) Err(code errcode.ErrorCoder, err error, msg ...string) {
	Error(e.Context, code, err, msg...)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg ...string) {
	OK(e.Context, data, msg...)
}

// PageOK 分页数据处理
func (e Api) PageOK(result interface{}, count int64, pageIndex int, pageSize int, msg ...string) {
	PageOK(e.Context, result, count, pageIndex, pageSize, msg...)
}

// getAcceptLanguage 获取当前语言
func (e *Api) getAcceptLanguage() string {
	languages := language.ParseAcceptLanguage(e.Context.GetHeader("Accept-Language"), nil)
	if len(languages) == 0 {
		return DefaultLanguage
	}
	return languages[0]
}

// GetRequestLogger 获取上下文提供的日志
func GetRequestLogger(c *gin.Context) *logger.Helper {
	var log *logger.Helper
	l, ok := c.Get(pkg.LoggerKey)
	if ok {
		log, ok = l.(*logger.Helper)
		if ok && log != nil {
			return log
		}
	}
	//如果没有在上下文中放入logger
	requestId := pkg.GenerateMsgIDFromContext(c)
	log = logger.NewHelper(logger.DefaultLogger).WithFields(map[string]interface{}{
		strings.ToLower(pkg.TrafficKey): requestId,
	})
	return log
}
