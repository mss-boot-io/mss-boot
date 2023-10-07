package config

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/12/14 03:12:11
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/12/14 03:12:11
 */

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestStorage_Init(t *testing.T) {
	type fields struct {
		Type            ProviderType
		SigningMethod   string
		Region          string
		Bucket          string
		AccessKeyID     string
		SecretAccessKey string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test",
			fields: fields{
				Type:            ProviderType(os.Getenv("s3_provider")),
				SigningMethod:   "v4",
				Region:          os.Getenv("s3_region"),
				Bucket:          os.Getenv("s3_bucket"),
				AccessKeyID:     os.Getenv("s3_access_key_id"),
				SecretAccessKey: os.Getenv("s3_secret_access_key"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Storage{
				Type:            tt.fields.Type,
				SigningMethod:   tt.fields.SigningMethod,
				Region:          tt.fields.Region,
				Bucket:          tt.fields.Bucket,
				AccessKeyID:     tt.fields.AccessKeyID,
				SecretAccessKey: tt.fields.SecretAccessKey,
			}
			o.Init()
			res, err := o.GetClient().ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
				Bucket:  aws.String(tt.fields.Bucket),
				MaxKeys: 10,
			})
			if err != nil {
				t.Fatalf("failed to list items: %v", err)
			}

			for _, o := range res.Contents {
				fmt.Println(">>> ", *o.Key)
			}
			_, err = o.GetClient().PutObject(context.TODO(), &s3.PutObjectInput{
				Bucket: aws.String(tt.fields.Bucket),
				Key:    aws.String("test.json"),
				Body:   bytes.NewBuffer([]byte(`{"name": "lwx"}`)),
			})
			if err != nil {
				t.Fatalf("failed to put object: %v", err)
			}
		})
	}
}
