/*
 * @Author: lwnmengjing
 * @Date: 2022/3/9 9:17
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/9 9:17
 */

package controller

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"log"
	"oauth2/models"
)

func ClientAuthorizedHandler(clientID string, _ oauth2.GrantType) (allowed bool, err error) {
	log.Println("client")
	client := &models.Client{}
	cm, err := client.GetByID(context.TODO(), clientID)
	if err != nil || cm == nil {
		return false, nil
	}
	return true, nil
}
