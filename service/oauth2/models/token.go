/*
 * @Author: lwnmengjing
 * @Date: 2022/3/9 11:10
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/9 11:10
 */

package models

import "C"
import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"oauth2/common"
	"time"
)

// NewToken create to token model instance
func NewToken() *Token {
	return &Token{}
}

// Token token model
type Token struct {
	ClientID            string        `bson:"ClientID"`
	UserID              string        `bson:"UserID"`
	RedirectURI         string        `bson:"RedirectURI"`
	Scope               string        `bson:"Scope"`
	Code                string        `bson:"Code"`
	CodeChallenge       string        `bson:"CodeChallenge"`
	CodeChallengeMethod string        `bson:"CodeChallengeMethod"`
	CodeCreateAt        time.Time     `bson:"CodeCreateAt"`
	CodeExpiresIn       time.Duration `bson:"CodeExpiresIn"`
	Access              string        `bson:"Access"`
	AccessCreateAt      time.Time     `bson:"AccessCreateAt"`
	AccessExpiresIn     time.Duration `bson:"AccessExpiresIn"`
	Refresh             string        `bson:"Refresh"`
	RefreshCreateAt     time.Time     `bson:"RefreshCreateAt"`
	RefreshExpiresIn    time.Duration `bson:"RefreshExpiresIn"`
}

func (t *Token) C() *mgo.Collection {
	return common.DB.C("token")
}

// Create and store the new token information
func (t *Token) Create(ctx context.Context, info oauth2.TokenInfo) error {
	now := time.Now()
	info.SetAccessCreateAt(now)
	info.SetRefreshCreateAt(now)
	return t.C().Insert(info)

}

// remove key
func (t *Token) remove(key string) error {
	return t.C().Remove(bson.M{"Code": key})
}

// RemoveByCode use the authorization code to delete the token information
func (t *Token) RemoveByCode(ctx context.Context, code string) error {
	return t.remove(code)
}

// RemoveByAccess use the access token to delete the token information
func (t *Token) RemoveByAccess(ctx context.Context, access string) error {
	return t.remove(access)
}

// RemoveByRefresh use the refresh token to delete the token information
func (t *Token) RemoveByRefresh(ctx context.Context, refresh string) error {
	return t.remove(refresh)
}

// GetByCode use the authorization code for token information data
func (t *Token) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	return t.getByCode(code)
}

// GetByAccess use the access token for token information data
func (t *Token) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	return t.getByAccess(access)
}

func (t *Token) getByAccess(access string) (oauth2.TokenInfo, error) {
	var tm models.Token
	err := t.C().Find(bson.M{"Access": access}).One(&tm)
	if err != nil {
		return nil, err
	}
	return &tm, nil
}

// GetByRefresh use the refresh token for token information data
func (t *Token) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	return t.getByRefresh(refresh)
}

func (t *Token) getByRefresh(refresh string) (oauth2.TokenInfo, error) {
	var tm models.Token
	err := t.C().Find(bson.M{"Refresh": refresh}).One(&tm)
	if err != nil {
		return nil, err
	}
	return &tm, nil
}

func (t *Token) getByCode(key string) (oauth2.TokenInfo, error) {
	var tm models.Token
	err := t.C().Find(bson.M{"Code": key}).One(&tm)
	if err != nil {
		return nil, err
	}
	return &tm, err
}

// New create to token model instance
func (t *Token) New() oauth2.TokenInfo {
	return NewToken()
}

// GetClientID the client id
func (t *Token) GetClientID() string {
	return t.ClientID
}

// SetClientID the client id
func (t *Token) SetClientID(clientID string) {
	t.ClientID = clientID
}

// GetUserID the user id
func (t *Token) GetUserID() string {
	return t.UserID
}

// SetUserID the user id
func (t *Token) SetUserID(userID string) {
	t.UserID = userID
}

// GetRedirectURI redirect URI
func (t *Token) GetRedirectURI() string {
	return t.RedirectURI
}

// SetRedirectURI redirect URI
func (t *Token) SetRedirectURI(redirectURI string) {
	t.RedirectURI = redirectURI
}

// GetScope get scope of authorization
func (t *Token) GetScope() string {
	return t.Scope
}

// SetScope get scope of authorization
func (t *Token) SetScope(scope string) {
	t.Scope = scope
}

// GetCode authorization code
func (t *Token) GetCode() string {
	return t.Code
}

// SetCode authorization code
func (t *Token) SetCode(code string) {
	t.Code = code
}

// GetCodeCreateAt create Time
func (t *Token) GetCodeCreateAt() time.Time {
	return t.CodeCreateAt
}

// SetCodeCreateAt create Time
func (t *Token) SetCodeCreateAt(createAt time.Time) {
	t.CodeCreateAt = createAt
}

// GetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) GetCodeExpiresIn() time.Duration {
	return t.CodeExpiresIn
}

// SetCodeExpiresIn the lifetime in seconds of the authorization code
func (t *Token) SetCodeExpiresIn(exp time.Duration) {
	t.CodeExpiresIn = exp
}

// GetCodeChallenge challenge code
func (t *Token) GetCodeChallenge() string {
	return t.CodeChallenge
}

// SetCodeChallenge challenge code
func (t *Token) SetCodeChallenge(code string) {
	t.CodeChallenge = code
}

// GetCodeChallengeMethod challenge method
func (t *Token) GetCodeChallengeMethod() oauth2.CodeChallengeMethod {
	return oauth2.CodeChallengeMethod(t.CodeChallengeMethod)
}

// SetCodeChallengeMethod challenge method
func (t *Token) SetCodeChallengeMethod(method oauth2.CodeChallengeMethod) {
	t.CodeChallengeMethod = string(method)
}

// GetAccess access Token
func (t *Token) GetAccess() string {
	return t.Access
}

// SetAccess access Token
func (t *Token) SetAccess(access string) {
	t.Access = access
}

// GetAccessCreateAt create Time
func (t *Token) GetAccessCreateAt() time.Time {
	return t.AccessCreateAt
}

// SetAccessCreateAt create Time
func (t *Token) SetAccessCreateAt(createAt time.Time) {
	t.AccessCreateAt = createAt
}

// GetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) GetAccessExpiresIn() time.Duration {
	return t.AccessExpiresIn
}

// SetAccessExpiresIn the lifetime in seconds of the access token
func (t *Token) SetAccessExpiresIn(exp time.Duration) {
	t.AccessExpiresIn = exp
}

// GetRefresh refresh Token
func (t *Token) GetRefresh() string {
	return t.Refresh
}

// SetRefresh refresh Token
func (t *Token) SetRefresh(refresh string) {
	t.Refresh = refresh
}

// GetRefreshCreateAt create Time
func (t *Token) GetRefreshCreateAt() time.Time {
	return t.RefreshCreateAt
}

// SetRefreshCreateAt create Time
func (t *Token) SetRefreshCreateAt(createAt time.Time) {
	t.RefreshCreateAt = createAt
}

// GetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) GetRefreshExpiresIn() time.Duration {
	return t.RefreshExpiresIn
}

// SetRefreshExpiresIn the lifetime in seconds of the refresh token
func (t *Token) SetRefreshExpiresIn(exp time.Duration) {
	t.RefreshExpiresIn = exp
}
