package urlPath

import "testing"

func TestGetBasePath(t *testing.T) {
	type args struct {
		pathStr string
	}
	tests := []struct {
		name         string
		args         args
		wantBasePath string
		wantErr      bool
	}{
		{
			name: "基本示例",
			args: args{
				pathStr: `https://www.liwenzhou.com/posts/Go/golang-menu/`,
			},
			wantBasePath: "golang-menu",
			wantErr:      false,
		},
		{
			name: "带参数的示例",
			args: args{
				pathStr: `https://www.liwenzhou.com/posts/Go/golang-menu/?id=1`,
			},
			wantBasePath: "golang-menu",
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBasePath, err := GetBasePath(tt.args.pathStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBasePath != tt.wantBasePath {
				t.Errorf("GetBasePath() gotBasePath = %v, want %v", gotBasePath, tt.wantBasePath)
			}
		})
	}
}
