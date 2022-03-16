package entity

import "time"

type Users struct {
	Id     uint64    // 主键
	Name   string    // 用户名
	Email  string    // 邮箱
	Mobile string    // 手机号
	Passwd string    // 密码
	Salt   string    // 盐值
	Ext    string    // 扩展字段
	Status int8      // 状态（0：未审核,1:通过 10删除）
	Ctime  uint32    // 创建时间
	Mtime  time.Time // 修改时间
}

type Trace struct {
	Id    uint64 // 主键
	Uid   uint64 // 用户主键
	Type  int8   // 类型(0:注册1::登录2:退出3:修改4:删除)
	Ip    uint32 // ip
	Ext   string // 扩展字段
	Ctime uint32 // 注册时间
}

type Device struct {
	Id     uint64 // 主键
	Uid    uint64 // 用户主键
	Client string // 客户端
	Model  string // 设备型号
	Ip     uint32 // ip地址
	Ext    string // 扩展信息
	Ctime  uint32 // 注册时间
}
