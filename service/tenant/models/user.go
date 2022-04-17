/*
 * @Author: lwnmengjing
 * @Date: 2022/3/14 13:25
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/14 13:25
 */

package models

import (
	"context"
	"errors"
	"github.com/casdoor/casdoor-go-sdk/auth"
	"github.com/mss-boot-io/mss-boot/pkg/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	log "github.com/mss-boot-io/mss-boot/core/logger"
	"github.com/mss-boot-io/mss-boot/pkg/config/mongodb"
	"github.com/mss-boot-io/mss-boot/pkg/security"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type User struct {
	ID        string      `bson:"_id" json:"id"`
	TenantID  string      `bson:"tenantID" json:"tenantID"`
	Username  string      `bson:"username" json:"username"`
	Nickname  string      `bson:"nickname" json:"nickname"`
	Avatar    string      `bson:"avatar" json:"avatar"`
	Email     string      `bson:"email" json:"email"`
	Phone     string      `bson:"phone" json:"phone"`
	Status    enum.Status `bson:"status" json:"status"`
	PWD       UserPwd     `bson:"pwd" json:"pwd"`
	Metadata  interface{} `bson:"metadata" json:"metadata"`
	CreatedAt time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt" bson:"updatedAt"`
}

type UserPwd struct {
	Salt string `bson:"salt" json:"salt"`
	Hash string `bson:"hash" json:"hash"`
}

func (User) TableName() string {
	return "user"
}

func (e *User) Make() {
	ops := options.Index()
	ops.SetName("name")
	ops.SetUnique(true)
	_, err := e.C().Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bsonx.Doc{
			{"username", bsonx.Int32(1)},
			{"tenantID", bsonx.Int32(1)},
		},
		Options: ops,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (e *User) C() *mongo.Collection {
	return mongodb.DB.Collection(e.TableName())
}

func (e *User) Encrypt(pwd string) (err error) {
	if pwd == "" {
		return errors.New("password can't be empty")
	}
	e.PWD.Salt = security.GenerateRandomKey16()
	e.PWD.Hash, err = security.SetPassword(pwd, e.PWD.Salt)
	return
}

func (e *User) VerifyPassword(pwd string) bool {
	verify, err := security.SetPassword(pwd, e.PWD.Salt)
	if err != nil {
		return false
	}
	return verify == e.PWD.Hash
}

// CreateOrUpdateUser create or update user
func CreateOrUpdateUser(claims *auth.Claims) {
	tenant := &Tenant{}
	err := tenant.C().FindOne(context.TODO(), bson.M{"name": claims.Owner}).Decode(tenant)
	if err != nil {
		log.Errorf("find tenant error: %s", err)
		return
	}
	usernamePostfix := ""
	if claims.Github != "" {
		usernamePostfix = "-" + claims.Github
	}
	user := &User{
		TenantID:  tenant.ID,
		Username:  claims.Name + usernamePostfix,
		Nickname:  claims.Name + usernamePostfix,
		Avatar:    claims.Avatar,
		Email:     claims.Email,
		Phone:     claims.Phone,
		Status:    enum.Enabled,
		Metadata:  *claims,
		CreatedAt: time.Now(),
	}
	user.UpdatedAt = user.CreatedAt
	user.ID = primitive.NewObjectID().Hex()
	_, err = user.C().InsertOne(context.TODO(), user)
	if err != nil {
		log.Errorf("insert user error: %s", err)
		return
	}
}
