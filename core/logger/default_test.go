package logger

/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2023/5/29 07:20:43
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2023/5/29 07:20:43
 */

import "testing"

func Test_logCallerFilePath(t *testing.T) {
	type args struct {
		loggingFilePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test linux/unix",
			args: args{
				loggingFilePath: "/root/project/test.go",
			},
			want: "project/test.go",
		},
		{
			name: "test windows",
			args: args{
				loggingFilePath: "C:\\root\\project\\test.go",
			},
			want: "project\\test.go",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logCallerFilePath(tt.args.loggingFilePath); got != tt.want {
				t.Errorf("logCallerFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
