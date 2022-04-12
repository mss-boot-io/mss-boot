/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/3/27 5:59
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/3/27 5:59
 */

package session

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStore struct {
}

type model struct {
	Key   string      `bson:"key"`
	Value interface{} `bson:"value"`
}

type store struct {
	sid     string
	expired int64
	c       mongo.Collection
	ctx     context.Context
}

func (s *store) Context() context.Context {
	return s.ctx
}

func (s *store) SessionID() string {
	return s.sid
}

func (s *store) Set(key string, value interface{}) {
	if errors.Is(s.c.FindOne(s.ctx, bson.M{"key": key}).Err(), mongo.ErrNoDocuments) {
		_, _ = s.c.InsertOne(s.ctx, &model{key, value})
		return
	}
	_, _ = s.c.UpdateOne(s.ctx, bson.M{"key": key}, &model{key, value})
	return
}

func (s *store) Get(key string) (interface{}, bool) {
	m := &model{}
	err := s.c.FindOne(s.ctx, bson.M{"key": key}).Decode(m)
	if err != nil || m.Key == "" {
		return nil, false
	}
	return m, true
}

func (s *store) Delete(key string) interface{} {
	m := &model{}
	err := s.c.FindOne(s.ctx, bson.M{"key": key}).Decode(m)
	if err != nil {
		return nil
	}
	_, _ = s.c.DeleteOne(s.ctx, bson.M{"key": key})
	return m.Value
}

func (s *store) Flush() error {
	return nil
}

func (s *store) Save() error {
	return nil
}
