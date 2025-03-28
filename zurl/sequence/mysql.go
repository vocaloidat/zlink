package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	sqlRepLaceIntoStub = `REPLACE INTO sequence (stub) VALUES ('a')`
)

type Mysql struct {
	conn sqlx.SqlConn
}

func NewMysql(dns string) SequenceSql {
	conn := sqlx.NewMysql(dns)
	return &Mysql{conn}
}

// Next 取出下一个号码
func (m Mysql) Next() (seq uint64, err error) {
	var stmt sqlx.StmtSession
	// 创建一个预编译语句
	stmt, err = m.conn.Prepare(sqlRepLaceIntoStub)
	if err != nil {
		logx.Errorw("mysql prepare error", logx.LogField{Key: "err", Value: err})
		return 0, err
	}
	defer stmt.Close()

	//执行mysql语句
	rest, err := stmt.Exec()
	if err != nil {
		logx.Errorw("mysql exec error", logx.LogField{Key: "err", Value: err})
		return 0, err
	}
	// 获取新增的id
	id, err := rest.LastInsertId()
	if err != nil {
		logx.Errorw("mysql LastInsertId error", logx.LogField{Key: "err", Value: err})
		return 0, err
	}
	return uint64(id), nil
}
