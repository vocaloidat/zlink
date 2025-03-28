package logic

import (
	"context"
	"database/sql"
	"errors"

	"zlink/zurl/internal/svc"
	"zlink/zurl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Show 展示短链接
func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 不查数据库了。直接查找布隆过滤器
	exists, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Error("布隆过滤器查询错误:err:", err)
		return nil, errors.New("布隆过滤器查询错误")
	}
	// 没有查到
	if !exists {
		return nil, errors.New("布隆过滤器没有这个短链接")
	}

	// 数据库查询短链接对应长连接
	u, err := l.svcCtx.ZUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: req.ShortUrl,
		Valid:  true,
	})
	if err != nil {
		logx.Error("数据库查询错误:err:", err)
		return nil, errors.New("请输入正确的链接")
	}

	// 返回短链接对应的长链接
	return &types.ShowResponse{
		LongUrl: u.Lurl.String,
	}, nil
}
