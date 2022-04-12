/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 14:06
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 14:06
 */

package config

import "github.com/casdoor/casdoor-go-sdk/auth"

type OAuth2 struct {
	Endpoint         string `yaml:"endpoint" json:"endpoint"`
	ClientID         string `yaml:"clientID" json:"clientID"`
	ClientSecret     string `yaml:"clientSecret" json:"clientSecret"`
	JwtPublicKey     string `yaml:"jwtPublicKey" json:"jwtPublicKey"`
	OrganizationName string `yaml:"organizationName" json:"organizationName"`
	ApplicationName  string `yaml:"applicationName" json:"applicationName"`
}

// Init 初始化
func (e *OAuth2) Init() {
	auth.InitConfig(
		e.Endpoint,
		e.ClientID,
		e.ClientSecret,
		e.JwtPublicKey,
		e.OrganizationName,
		e.ApplicationName)
}
