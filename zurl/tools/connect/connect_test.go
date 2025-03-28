package connect

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCheckWebsite(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "基本示例",
			args: args{
				url: `https://www.liwenzhou.com/posts/Go/golang-menu/`,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckWebsite(tt.args.url); got != tt.want {
				t.Errorf("CheckWebsite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckWebsite2(t *testing.T) {
	c.Convey("基础用例", t, func() {
		var (
			url = `https://www.liwenzhou.com/posts/Go/golang-menu/`
		)
		got := CheckWebsite(url)
		c.So(got, c.ShouldEqual, true) // 断言

		// 或者 结果等于TRUE
		//c.ShouldBeTrue(got)
	})

	c.Convey("参数错误用例", t, func() {
		var (
			url = `/posts/Go/golang-menu/`
		)
		got := CheckWebsite(url)
		c.So(got, c.ShouldEqual, false) // 断言

		// 或者 结果等于TRUE
		//c.ShouldBeTrue(got)
	})
}
