// 自动生成模板MyTestOne
package my_test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MyTestOne 结构体
type MyTestOne struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;size:32;"`
	Code string `json:"code" form:"code" gorm:"column:code;comment:编码;size:32;"`
	Msg  string `json:"msg" form:"msg" gorm:"column:msg;comment:备注;size:255;"`
	Age  *bool  `json:"age" form:"age" gorm:"column:age;comment:数字测试;"`
}

// TableName MyTestOne 表名
func (MyTestOne) TableName() string {
	return "my_test_one"
}
