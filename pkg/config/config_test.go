package config

import (
	"fmt"
	"testing"
)

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/12/19 11:24:58
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/12/19 11:24:58
 */

func Test_parseTemplateWithEnv(t *testing.T) {
	type args struct {
		rb []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test0",
			args: args{
				rb: []byte(`{{.Env.TEST}}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTemplateWithEnv(tt.args.rb)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTemplateWithEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(got))
		})
	}
}
