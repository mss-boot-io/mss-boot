/*
 * @Author: lwnmengjing
 * @Date: 2022/3/8 17:24
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/8 17:24
 */

package manage

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-session/session"
	"github.com/sanity-io/litter"

	"oauth2/common"
	"oauth2/controllers"
	"oauth2/models"
)

func Init() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(models.NewToken(), nil)
	//manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	// manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))
	manager.MapAccessGenerate(generates.NewAccessGenerate())

	clientStore := &models.Client{}
	//clientStore.Set(&models.Client{
	//	ID:          "222222",
	//	Secret:      "22222222",
	//	HomepageURL: "http://localhost:9094",
	//	CreatedAt:   time.Now(),
	//	UpdatedAt:   time.Now(),
	//})
	//u := &models.User{
	//	Username: "test",
	//	Nickname: "test",
	//}
	//if err := u.Register("test"); err != nil {
	//	panic(err)
	//}
	manager.MapClientStorage(clientStore)

	common.OAuth2Srv = server.NewServer(server.NewConfig(), manager)

	common.OAuth2Srv.SetClientAuthorizedHandler(controllers.ClientAuthorizedHandler)
	common.OAuth2Srv.SetPasswordAuthorizationHandler(controllers.PasswordAuthorizationHandler)

	common.OAuth2Srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	common.OAuth2Srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	common.OAuth2Srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	if true {
		_ = dumpRequest(os.Stdout, "userAuthorizeHandler", r) // Ignore the error
	}
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		println(11)
		litter.Dump(r.Form)

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/oauth2/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	println(22)
	litter.Dump(uid)

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

func dumpRequest(writer io.Writer, header string, r *http.Request) error {
	data, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	writer.Write([]byte("\n" + header + ": \n"))
	writer.Write(data)
	return nil
}
