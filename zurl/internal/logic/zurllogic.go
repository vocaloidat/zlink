package logic

import (
	"context"

	"zlink/zurl/internal/svc"
	"zlink/zurl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZurlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZurlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZurlLogic {
	return &ZurlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZurlLogic) Zurl(req *types.Request) (resp *types.Response, err error) {
	// 根据短链接请求标识符，查找长链接

	// 查询到
	if req.ShortUrl == "123456" {
		return &types.Response{
			LongUrl: `https://www.liwenzhou.com/posts/Go/golang-menu/`,
		}, nil
	}
	// 查询不到
	return &types.Response{
		LongUrl: `https://www.liwenzhou.com/`,
	}, nil
}
