package logic

import (
	"context"

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
	//

	return
}
