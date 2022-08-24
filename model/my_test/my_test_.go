// 自动生成模板MyTest
package my_test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MyTest 结构体
type MyTest struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:32;"`
	Code string `json:"code" form:"code" gorm:"column:code;comment:编码;size:32;"`
	Msg  string `json:"msg" form:"msg" gorm:"column:msg;comment:备注;size:255;"`
}

// TableName MyTest 表名
func (MyTest) TableName() string {
	return "my_test"
}
