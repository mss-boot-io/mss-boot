/*
 * @Author: lwnmengjing
 * @Date: 2021/6/23 2:17 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/23 2:17 下午
 */

package controller

import (
	"net/http"
	"net/url"
	"oauth2/common"
	"oauth2/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
	"github.com/mss-boot-io/mss-boot/pkg/response"
)

func init() {
	response.AppendController(&OAuth2{})
}

type OAuth2 struct {
	response.Api
}

func (OAuth2) Path() string {
	return "/tenant"
}

func (e OAuth2) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func (e OAuth2) Create(c *gin.Context) {
}

func (e OAuth2) Update(c *gin.Context) {
}

func (e OAuth2) Delete(c *gin.Context) {
}

func (e OAuth2) Get(c *gin.Context) {
}

func (e OAuth2) List(c *gin.Context) {
}

func (e OAuth2) Other(r *gin.RouterGroup) {
	r.Any("/login", e.Login)
	r.GET("/auth", e.Auth)
	r.Any("/oauth/authorize", e.Authorize)
	r.Any("/oauth/token", e.Token)
	r.Any("/test", e.Check)
}

func (e OAuth2) Login(c *gin.Context) {
	store, err := session.Start(c, c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if c.Request.Method == http.MethodPost {
		if c.Request.Form == nil {
			if err = c.Request.ParseForm(); err != nil {
				http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		u := models.User{Username: c.Request.Form.Get("username")}
		if !u.VerifyPassword(c.Request.Form.Get("password")) {
			http.Error(c.Writer, "password incorrect", http.StatusUnauthorized)
			return
		}
		store.Set("LoggedInUserID", c.Request.Form.Get("username"))
		err = store.Save()
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Redirect(http.StatusFound, "/oauth2/auth")
		return
	}
	outputHTML(c, "static/login.html")
}

func (e OAuth2) Auth(c *gin.Context) {
	store, err := session.Start(c, c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		c.Redirect(http.StatusFound, "/oauth2/login")
		return
	}
	outputHTML(c, "static/auth.html")
}

func (e OAuth2) Authorize(c *gin.Context) {
	store, err := session.Start(c, c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	c.Request.Form = form

	store.Delete("ReturnUri")
	err = store.Save()
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = common.OAuth2Srv.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (e OAuth2) Token(c *gin.Context) {
	err := common.OAuth2Srv.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e OAuth2) Check(c *gin.Context) {
	token, err := common.OAuth2Srv.ValidationBearerToken(c.Request)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	})
}

func outputHTML(c *gin.Context, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(c.Writer, c.Request, file.Name(), fi.ModTime(), file)
}
