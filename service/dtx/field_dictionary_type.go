package dtx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dtx"
	dtxReq "github.com/flipped-aurora/gin-vue-admin/server/model/dtx/request"
)

type FieldDictionaryTypeService struct {
}

// CreateFieldDictionaryType 创建FieldDictionaryType记录
// Author [piexlmax](https://github.com/piexlmax)
func (fieldDictionaryTypeService *FieldDictionaryTypeService) CreateFieldDictionaryType(fieldDictionaryType dtx.FieldDictionaryType) (err error) {
	err = global.GVA_DB.Create(&fieldDictionaryType).Error
	return err
}

// DeleteFieldDictionaryType 删除FieldDictionaryType记录
// Author [piexlmax](https://github.com/piexlmax)
func (fieldDictionaryTypeService *FieldDictionaryTypeService) DeleteFieldDictionaryType(fieldDictionaryType dtx.FieldDictionaryType) (err error) {
	err = global.GVA_DB.Delete(&fieldDictionaryType).Error
	return err
}

// DeleteFieldDictionaryTypeByIds 批量删除FieldDictionaryType记录
// Author [piexlmax](https://github.com/piexlmax)
func (fieldDictionaryTypeService *FieldDictionaryTypeService) DeleteFieldDictionaryTypeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dtx.FieldDictionaryType{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFieldDictionaryType 更新FieldDictionaryType记录
// Author [piexlmax](https://github.com/piexlmax)
func (fieldDictionaryTypeService *FieldDictionaryTypeService) UpdateFieldDictionaryType(fieldDictionaryType dtx.FieldDictionaryType) (err error) {
	err = global.GVA_DB.Save(&fieldDictionaryType).Error
	return err
}

// GetFieldDictionaryType 根据id获取FieldDictionaryType记录
// Author [piexlmax](https://github.com/piexlmax)
func (fieldDictionaryTypeService *FieldDictionaryTypeService) GetFieldDictionaryType(id uint) (fieldDictionaryType dtx.FieldDictionaryType, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fieldDictionaryType).Error
	return
}

// GetFieldDictionaryTypeInfoList 分页获取FieldDictionaryType记录
// Author [piexlmax](https://github.com/piexlmax)
func (fieldDictionaryTypeService *FieldDictionaryTypeService) GetFieldDictionaryTypeInfoList(info dtxReq.FieldDictionaryTypeSearch) (list []dtx.FieldDictionaryType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dtx.FieldDictionaryType{})
	var fieldDictionaryTypes []dtx.FieldDictionaryType
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TypeCode != "" {
		db = db.Where("type_code LIKE ?", "%"+info.TypeCode+"%")
	}
	if info.TypeName != "" {
		db = db.Where("type_name LIKE ?", "%"+info.TypeName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&fieldDictionaryTypes).Error
	return fieldDictionaryTypes, total, err
}
