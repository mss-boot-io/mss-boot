package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

const (
	TrafficKey = "X-Request-Id"
	LoggerKey  = "_go-admin-logger-request"
)

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// GenerateMsgIDFromContext 生成msgID
func GenerateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(TrafficKey, requestId)
	}
	return requestId
}

// ModelDeepCopy model deep copy
func ModelDeepCopy(m mgm.Model) mgm.Model {
	return reflect.New(reflect.TypeOf(m).Elem()).Interface().(mgm.Model)
}

// DeepCopy deep copy
func DeepCopy(d any) any {
	return reflect.New(reflect.TypeOf(d).Elem()).Interface()
}
