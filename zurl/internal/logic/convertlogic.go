package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zlink/zurl/internal/svc"
	"zlink/zurl/internal/types"
	"zlink/zurl/tools/connect"
	"zlink/zurl/tools/md5"
	"zlink/zurl/tools/urlPath"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Convert 转换短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. 校验数据
	// 1.1 校验参数合法性
	// 在ConvertHandler已经完成

	// 1.2 判断传入的参数是否是个正常的网址
	ok := connect.CheckWebsite(req.LongUrl)
	if !ok {
		return nil, errors.New("传入的长连接，无法到达，请检查。")
	}
	// 1.3 判断之前是否转链过
	lUrl := md5.GetMd5(req.LongUrl)
	_, err = l.svcCtx.ZUrlModel.FindOneByMd5(l.ctx, sql.NullString{
		String: lUrl,
		Valid:  true,
	})
	if !errors.Is(err, sqlx.ErrNotFound) {
		logx.Error("数据库查询失败，记录错误:err:", err)
		return nil, errors.New("数据库查询失败，记录错误")
	}

	if err == nil {
		return nil, errors.New("传入的长连接，已存在对应的短链接，请检查管理中心,err:" + err.Error())
	}
	// 1.4 输入的是否已经是个短链接了。
	// http://www.baidu.com/oskfhg?name=wyc
	// 这时候，要拿到的是，oskfhg
	basePath, err := urlPath.GetBasePath(req.LongUrl)
	if err != nil {
		return nil, err
	}
	// 判断是否短链接已经存在
	_, err = l.svcCtx.ZUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: basePath,
		Valid:  true,
	})
	if !errors.Is(err, sqlx.ErrNotFound) {
		logx.Error("数据库查询失败，记录错误:err:", err)
		return nil, errors.New("数据库查询失败，记录错误")
	}
	if err == nil {
		return nil, errors.New("传入的连接，已经是短链接。无法继续生成，err:" + err.Error())
	}
	// 2. 取号器取号
	seq, err := l.svcCtx.Sequence.Next()
	if err != nil {
		return nil, errors.New("Sequence.Next()无法继续生成 err:" + err.Error())
	}
	fmt.Println("seq:", seq)
	// 3. 生成短链接
	// 4. 插入短链接表
	// 5. 返回响应

	return
}
