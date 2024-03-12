package gorm

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/3/4 01:30:34
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/3/4 01:30:34
 */

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/pkg"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/response"
	"gorm.io/gorm"
)

// Control action
type Control struct {
	Base
	Key string
}

// String action name
func (*Control) String() string {
	return "control"
}

// NewControl new control action
func NewControl(b Base, key string) *Control {
	return &Control{
		Base: b,
		Key:  key,
	}
}

func (e *Control) Handler() gin.HandlersChain {
	h := func(c *gin.Context) {
		if e.Model == nil {
			response.Make(c).Err(http.StatusNotImplemented, "not implemented")
			return
		}
		switch c.Request.Method {
		case http.MethodPost:
			e.create(c)
		case http.MethodPut:
			e.update(c)
		default:
			response.Make(c).Err(http.StatusNotImplemented, "not implemented")
		}
	}
	if e.Handlers != nil {
		return append(e.Handlers, h)
	}
	return gin.HandlersChain{h}
}

func (e *Control) create(c *gin.Context) {
	m := pkg.TablerDeepCopy(e.Model)
	api := response.Make(c).Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	err := gormdb.DB.WithContext(c).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(m).Error
		if err != nil {
			api.AddError(err).Log.ErrorContext(c, "Create error", "error", err)
			api.Err(http.StatusInternalServerError)
			return err
		}
		if pkg.SupportCreator(m) {
			verify := response.VerifyHandler(c)
			if verify == nil {
				api.Err(http.StatusUnauthorized)
				return errors.New("verify handler is nil")
			}
			err = tx.Model(m).Update(pkg.GetCreatorField(), verify.GetUserID()).Error
			if err != nil {
				api.AddError(err).Log.ErrorContext(c, "Create error", "error", err)
				api.Err(http.StatusInternalServerError)
				return err
			}
		}
		return nil
	})

	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Create error", "error", err)
		return
	}

	api.OK(m)
}

func (e *Control) update(c *gin.Context) {
	m := pkg.TablerDeepCopy(e.Model)
	id := c.Param(e.Key)
	api := response.Make(c)
	if id == "" {
		api.AddError(errors.New("id is empty"))
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	query := gormdb.DB.Where(e.Key, id)
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, m))
	}
	//find object
	err := query.First(m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.AddError(fmt.Errorf("%s(%s) record not found", e.Key, id))
			api.Err(http.StatusNotFound)
			return
		}
		api.AddError(err).Log.ErrorContext(c, "Update error", "error", err.Error())
		api.Err(http.StatusInternalServerError)
		return
	}

	api = api.Bind(m)
	if api.Error != nil {
		api.Err(http.StatusUnprocessableEntity)
		return
	}
	query = gormdb.DB.WithContext(c)
	if e.Scope != nil {
		query = query.Scopes(e.Scope(c, m))
	}
	err = query.Save(m).Error
	if err != nil {
		api.AddError(err).Log.ErrorContext(c, "Update error", "error", err.Error())
		api.Err(http.StatusInternalServerError)
		return
	}
	api.OK(m)
}
