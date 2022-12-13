/*
 * @Author: lwnmengjing
 * @Date: 2022/7/18 10:06:11
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2022/7/18 10:06:11
 */

package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Source struct {
	opt Options
}

func (s *Source) Open(string) (fs.File, error) {
	return nil, errors.New("method Get not implemented")
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), s.opt.timeout)
	defer cancel()
	object, err := s.opt.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.opt.bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", s.opt.dir, name)),
	})
	if err != nil {
		return nil, err
	}
	defer object.Body.Close()
	return io.ReadAll(object.Body)
}

// New source
func New(options ...Option) (*Source, error) {
	s := &Source{}
	for _, opt := range options {
		opt(&s.opt)
	}
	if s.opt.timeout == 0 {
		s.opt.timeout = 5 * time.Second
	}
	if s.opt.client != nil {
		return s, nil
	}

	ctx, cancel := context.WithTimeout(context.TODO(), s.opt.timeout)
	defer cancel()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(s.opt.region))
	if err != nil {
		return nil, err
	}
	s.opt.client = s3.NewFromConfig(cfg)
	return s, nil
}
