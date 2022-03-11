/*
 * @Author: lwnmengjing
 * @Date: 2022/3/9 9:44
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/9 9:44
 */

package models

import (
	"errors"
	"github.com/mss-boot-io/mss-boot/pkg/security"
	"github.com/sanity-io/litter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"

	"oauth2/common"
)

type User struct {
	Username string
	Nickname string
	Avatar   string
	Email    string
	Phone    string
	Status   uint8
	PWD      UserPwd
}

type UserPwd struct {
	Salt string
	Hash string
}

func (e *User) C() *mgo.Collection {
	return common.DB.C("user")
}

func (e *User) Encrypt(pwd string) (err error) {
	if pwd == "" {
		return errors.New("password can't be empty")
	}
	e.PWD.Salt = security.GenerateRandomKey16()
	e.PWD.Hash, err = security.SetPassword(pwd, e.PWD.Salt)
	return
}

func (e *User) Register(pwd string) error {
	n, err := e.C().Find(User{
		Username: e.Username}).
		Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return errors.New("username already exists")
	}
	if err = e.Encrypt(pwd); err != nil {
		return err
	}
	return e.C().Insert(e)
}

func (e *User) GetUserByUsername(username string) error {
	e.Username = username
	return e.C().Find(bson.M{
		"username": e.Username}).One(e)
}

func (e *User) VerifyPassword(pwd string) bool {
	if e.C().Find(bson.M{"username": e.Username}).One(e) != nil {
		return false
	}
	litter.Dump(e)
	verify, err := security.SetPassword(pwd, e.PWD.Salt)
	println(verify)
	if err != nil {
		log.Println(err)
		return false
	}
	return verify == e.PWD.Hash
}
