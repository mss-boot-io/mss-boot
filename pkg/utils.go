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

func BuildMap(keys []string, value string) map[string]any {
	data := make(map[string]any)
	if len(keys) > 1 {
		data[keys[0]] = BuildMap(keys[1:], value)
	} else {
		return map[string]any{keys[0]: value}
	}
	return data
}

// MergeMapsDepth deep merge multi map
func MergeMapsDepth(ms ...map[string]any) map[string]any {
	data := make(map[string]any)
	for i := range ms {
		data = MergeMapDepth(data, ms[i])
	}
	return data
}

// MergeMapDepth deep merge map
func MergeMapDepth(m1, m2 map[string]any) map[string]any {
	for k := range m2 {
		if v, ok := m1[k]; ok {
			if m, ok := v.(map[string]any); ok {
				m1[k] = MergeMapDepth(m, m2[k].(map[string]any))
			} else {
				m1[k] = m2[k]
			}
		} else {
			m1[k] = m2[k]
		}
	}
	return m1
}

// MergeMap merge map
func MergeMap(m1, m2 map[string]any) map[string]any {
	for k := range m2 {
		m1[k] = m2[k]
	}
	return m1
}
