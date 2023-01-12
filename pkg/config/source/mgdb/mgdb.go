/*
 * @Author: lwnmengjing
 * @Date: 2023/1/13 04:03:10
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/1/13 04:03:10
 */

package mgdb

import (
	"context"
	"errors"
	"io/fs"
	"strings"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/mss-boot-io/mss-boot/pkg/config/source"
	"go.mongodb.org/mongo-driver/bson"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type Source struct {
	opt source.Options
}

func (s *Source) Open(string) (fs.File, error) {
	return nil, errors.New("method Get not implemented")
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	if strings.Index(name, ".") > -1 {
		name = name[:strings.Index(name, ".")]
	}
	m := SystemConfig{}
	ctx, cancel := context.WithTimeout(
		context.TODO(),
		s.opt.Timeout)
	defer cancel()
	err := mgm.Coll(&m).FirstWithCtx(ctx, bson.M{"name": name}, &m)
	if err != nil {
		return nil, err
	}
	return m.GenerateYAML()
}

// New source
func New(options ...source.Option) (*Source, error) {
	s := &Source{}
	for _, opt := range options {
		opt(&s.opt)
	}
	if s.opt.Timeout == 0 {
		s.opt.Timeout = 5 * time.Second
	}
	serverAPIOptions := mongoOptions.ServerAPI(mongoOptions.ServerAPIVersion1)
	clientOptions := mongoOptions.Client().
		ApplyURI(s.opt.MongoDBURL).
		SetServerAPIOptions(serverAPIOptions)
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: s.opt.Timeout}, s.opt.MongoDBName, clientOptions)
	if err != nil {
		return nil, err
	}
	return s, nil
}
