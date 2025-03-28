package md5

import "testing"

func TestGetMd5(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "基本用例",
			args: args{
				str: "123456",
			},
			want: "e10adc3949ba59abbe56e057f20f883e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMd5(tt.args.str); got != tt.want {
				t.Errorf("GetMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}
