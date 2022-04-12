/*
 * @Author: lwnmengjing
 * @Date: 2022/3/9 9:23
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/9 9:23
 */

package models

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/sanity-io/litter"
	"strings"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"oauth2/common"
)

type Client struct {
	ID          string `bson:"_id"`
	Secret      string
	Scopes      []string
	CallbackURL string
	HomepageURL string
	Application Application
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Application struct {
	Name        string
	Logo        string
	Description string
}

func (e *Client) C() *mgo.Collection {
	return common.DB.C("client")
}

func (e *Client) Insert() error {
	e.Secret = fmt.Sprintf("%X",
		sha256.Sum256([]byte(uuid.New().String())))
	e.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	return e.C().Insert(e)
}

func (e *Client) Delete() error {
	return e.C().RemoveId(e.ID)
}

func (e *Client) VerifyPassword(secret string) bool {
	println("verify")
	return e.Secret == secret
}

func (e *Client) GetID() string {
	return e.ID
}

func (e *Client) GetSecret() string {
	return e.Secret
}

func (e *Client) GetDomain() string {
	return e.HomepageURL
}

func (e *Client) GetUserID() string {
	return ""
}

func (e *Client) GetByID(_ context.Context, id string) (oauth2.ClientInfo, error) {
	err := e.C().Find(bson.M{"id": id}).One(e)
	litter.Dump(e)
	return e, err
}

func (e *Client) Set(cli oauth2.ClientInfo) error {
	n, err := e.C().FindId(cli.GetID()).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return nil
	}
	return e.C().Insert(cli)
}
