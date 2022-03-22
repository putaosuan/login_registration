package entity

import (
	"regexp"
	"time"
)

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

//校验手机号格式
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

//校验邮箱格式
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
