package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zlink/zurl/internal/config"
	"zlink/zurl/internal/svc"
	"zlink/zurl/internal/types"
	"zlink/zurl/model"
	"zlink/zurl/tools/base62"
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
	lUrl_md5 := md5.GetMd5(req.LongUrl)
	_, err = l.svcCtx.ZUrlModel.FindOneByMd5(l.ctx, sql.NullString{
		String: lUrl_md5,
		Valid:  true,
	})
	if !errors.Is(err, sqlx.ErrNotFound) {
		logx.Error("传入的长连接，已存在对应的短链接，请检查管理中心:err:", err)
		return nil, errors.New("传入的长连接，已存在对应的短链接，请检查管理中心")
	}

	if err == nil {
		return nil, errors.New("传入的长连接，已存在对应的短链接，请检查管理中心,err")
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
		logx.Error("传入的长连接，已存在对应的短链接，请检查管理中心:err:", err)
		return nil, errors.New("传入的长连接，已存在对应的短链接，请检查管理中心")
	}
	if err == nil {
		return nil, errors.New("传入的连接，已经是短链接。无法继续生成，err:" + err.Error())
	}
	var short string
	for {
		// 2. 取号器取号
		seq, err := l.svcCtx.Sequence.Next()
		if err != nil {
			return nil, errors.New("Sequence.Next()无法继续生成 err:" + err.Error())
		}
		fmt.Println("seq:", seq)
		// 3. 生成短链接
		// 3.1 安全性 可以通过自定义初始化解决
		// 3.2 敏感词检测
		short = base62.EncodeIntToBase62(seq)
		//fmt.Println("short:", short)
		_, ok := config.BlackLCMap[short]
		// 敏感词没找到 返回false
		if !ok {
			break
		}
	}
	// 4. 插入短链接表
	_, err = l.svcCtx.ZUrlModel.Insert(l.ctx, &model.ShortUrlMap{
		CreateBy: "admin",
		IsDel:    0,
		Lurl: sql.NullString{
			String: req.LongUrl,
			Valid:  true,
		},
		Md5: sql.NullString{
			String: lUrl_md5,
			Valid:  true,
		},
		Surl: sql.NullString{
			String: short,
			Valid:  true,
		},
	})
	if err != nil {
		logx.Error("数据库写入失败，记录错误:err:", err)
		return nil, errors.New("数据库写入失败，记录错误,err:" + err.Error())
	}
	// 5. 返回响应
	// 返回完整的短域名
	resp = &types.ConvertResponse{
		ShortUrl: l.svcCtx.Config.ZUrlDoamin + "/" + short,
	}
	return resp, nil
}
