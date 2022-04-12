/*
 * @Author: lwnmengjing
 * @Date: 2022/3/9 9:17
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/9 9:17
 */

package controllers

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/mss-boot-io/mss-boot/pkg/response"

	"oauth2/models"
)

func init() {
	response.AppendController(&Client{})
}

type Client struct {
	response.Api
}

func (Client) Path() string {
	return "/client"
}

func (e Client) Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{}
}

func (e Client) Create(c *gin.Context) {

}

func (e Client) Update(c *gin.Context) {
}

func (e Client) Delete(c *gin.Context) {
}

func (e Client) Get(c *gin.Context) {
}

func (e Client) List(c *gin.Context) {
}

func (e Client) Other(r *gin.RouterGroup) {

}

func ClientAuthorizedHandler(clientID string, _ oauth2.GrantType) (allowed bool, err error) {
	log.Println("client")
	client := &models.Client{}
	cm, err := client.GetByID(context.TODO(), clientID)
	if err != nil || cm == nil {
		return false, nil
	}
	return true, nil
}