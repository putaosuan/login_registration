package entity

type Device struct {
	Id     uint64 // 主键
	Uid    uint64 // 用户主键
	Client string // 客户端
	Model  string // 设备型号
	Ip     uint32 // ip地址
	Ext    string // 扩展信息
	Ctime  uint32 // 注册时间
}
