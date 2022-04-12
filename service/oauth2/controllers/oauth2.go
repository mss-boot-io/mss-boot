/*
 * @Author: lwnmengjing
 * @Date: 2021/6/23 2:17 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/23 2:17 下午
 */

package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/oauth2"
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
	return "/"
}

func (e OAuth2) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func (e OAuth2) Create(_ *gin.Context) {
}

func (e OAuth2) Update(_ *gin.Context) {
}

func (e OAuth2) Delete(_ *gin.Context) {
}

func (e OAuth2) Get(_ *gin.Context) {
}

func (e OAuth2) List(_ *gin.Context) {
}

func (e OAuth2) Other(r *gin.RouterGroup) {
	r.GET("/login", e.Login)
	r.POST("/login", e.Login)
	r.GET("/auth", e.Auth)
	r.GET("/oauth/authorize", e.Authorize)
	r.POST("/oauth/authorize", e.Authorize)
	r.GET("/oauth/token", e.Token)
	r.POST("/oauth/token", e.Token)
	r.GET("/test", e.Check)
	r.GET("/url", e.URL)
}

func genCodeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))
	return base64.URLEncoding.EncodeToString(s256[:])
}

func (e OAuth2) URL(c *gin.Context) {
	config := oauth2.Config{
		ClientID:     c.Query("id"),
		ClientSecret: c.Query("secret"),
		Scopes:       c.QueryArray("scopes"),
		RedirectURL:  c.Query("redirect"),
	}
	e.OK(config.AuthCodeURL(c.Query("state"), oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256("s256example")),
		oauth2.SetAuthURLParam("code_challenge_method", "S256")))
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
	c.Redirect(http.StatusFound, "http://localhost:8000/#/mss-boot-frontend/user/login")
	//outputHTML(c, "static/login.html")
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
