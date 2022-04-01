package entity

import (
	"my_sso/internal/user/domain/valobj"
)

type Trace struct {
	Id    uint64           // 主键
	Uid   uint64           // 用户主键
	Type  valobj.TraceType // 类型(0:注册1::登录2:退出3:修改4:删除)
	Ip    uint32           // ip
	Ext   string           // 扩展字段
	Ctime uint32           // 注册时间
}
