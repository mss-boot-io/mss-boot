/*
 * @Author: lwnmengjing
 * @Date: 2022/3/10 14:06
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/3/10 14:06
 */

package config

import "golang.org/x/oauth2"

type OAuth2 struct {
	ID          string         `yaml:"id" json:"id"`
	Secret      string         `yaml:"secret" json:"secret"`
	Scopes      []string       `yaml:"scopes" json:"scopes"`
	RedirectURL string         `yaml:"redirectURL" json:"redirectURL"`
	Endpoint    OAuth2Endpoint `yaml:"endpoint" json:"endpoint"`
}

type OAuth2Endpoint struct {
	AuthURL   string `yaml:"authURL" json:"authURL"`
	TokenURL  string `yaml:"tokenURL" json:"tokenURL"`
	AuthStyle int    `yaml:"authStyle" json:"authStyle"`
}

func (e OAuth2) Init() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     e.ID,
		ClientSecret: e.Secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   e.Endpoint.AuthURL,
			TokenURL:  e.Endpoint.TokenURL,
			AuthStyle: oauth2.AuthStyle(e.Endpoint.AuthStyle),
		},
		RedirectURL: e.RedirectURL,
		Scopes:      e.Scopes,
	}
}
