/*
 * @Author: lwnmengjing
 * @Date: 2022/3/9 11:47
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/9 11:47
 */

package controller

import (
	"context"
	"errors"
	"log"
	"oauth2/models"
)

// PasswordAuthorizationHandler verify password
func PasswordAuthorizationHandler(
	_ context.Context,
	username, password string) (id string, err error) {
	u := &models.User{Username: username}
	ok := u.VerifyPassword(password)
	log.Println(ok)
	if !ok {
		err = errors.New("password incorrect")
		return
	}
	return u.Username, nil
}
