// 自动生成模板FieldDictionaryType
package dtx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FieldDictionaryType 结构体
type FieldDictionaryType struct {
	global.GVA_MODEL
	OriginId string `json:"originId" form:"originId" gorm:"column:origin_id;comment:原始id;size:32;"`
	TypeCode string `json:"typeCode" form:"typeCode" gorm:"column:type_code;comment:类型编码;size:32;"`
	TypeName string `json:"typeName" form:"typeName" gorm:"column:type_name;comment:类型名称;size:32;"`
	Remark   string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:200;"`
	//CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
	//UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`
	IsDelete *bool  `json:"isDelete" form:"isDelete" gorm:"column:is_delete;comment:删除标识，0：未删除，1：已删除;"`
	Status   *int   `json:"status" form:"status" gorm:"column:status;comment:状态;"`
	CreateBy string `json:"createBy" form:"createBy" gorm:"column:create_by;comment:创建者名称;size:32;"`
	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"column:update_by;comment:修改者名称;size:32;"`
}

// TableName FieldDictionaryType 表名
func (FieldDictionaryType) TableName() string {
	return "field_dictionary_type"
}
