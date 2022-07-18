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
	"io/fs"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Source struct {
	bucket  string
	dir     string
	timeout time.Duration
	client  *s3.Client
}

func (s *Source) Open(string) (fs.File, error) {
	return nil, errors.New("method Get not implemented")
}

func (s *Source) ReadFile(name string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), s.timeout)
	defer cancel()
	object, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", s.dir, name)),
	})
	if err != nil {
		return nil, err
	}
	defer object.Body.Close()
	return ioutil.ReadAll(object.Body)
}

// New source
func New(region, bucket, dir string, timeout time.Duration) (*Source, error) {
	s := &Source{
		bucket:  bucket,
		dir:     dir,
		timeout: timeout,
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	s.client = s3.NewFromConfig(cfg)
	return s, nil
}
