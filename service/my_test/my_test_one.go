package my_test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/my_test"
	my_testReq "github.com/flipped-aurora/gin-vue-admin/server/model/my_test/request"
)

type MyTestOneService struct {
}

// CreateMyTestOne 创建MyTestOne记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestOneService *MyTestOneService) CreateMyTestOne(myTestOne my_test.MyTestOne) (err error) {
	err = global.GVA_DB.Create(&myTestOne).Error
	return err
}

// DeleteMyTestOne 删除MyTestOne记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestOneService *MyTestOneService) DeleteMyTestOne(myTestOne my_test.MyTestOne) (err error) {
	err = global.GVA_DB.Delete(&myTestOne).Error
	return err
}

// DeleteMyTestOneByIds 批量删除MyTestOne记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestOneService *MyTestOneService) DeleteMyTestOneByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]my_test.MyTestOne{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMyTestOne 更新MyTestOne记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestOneService *MyTestOneService) UpdateMyTestOne(myTestOne my_test.MyTestOne) (err error) {
	err = global.GVA_DB.Save(&myTestOne).Error
	return err
}

// GetMyTestOne 根据id获取MyTestOne记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestOneService *MyTestOneService) GetMyTestOne(id uint) (myTestOne my_test.MyTestOne, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&myTestOne).Error
	return
}

// GetMyTestOneInfoList 分页获取MyTestOne记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestOneService *MyTestOneService) GetMyTestOneInfoList(info my_testReq.MyTestOneSearch) (list []my_test.MyTestOne, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&my_test.MyTestOne{})
	var myTestOnes []my_test.MyTestOne
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Code != "" {
		db = db.Where("code = ?", info.Code)
	}
	if info.Msg != "" {
		db = db.Where("msg LIKE ?", "%"+info.Msg+"%")
	}
	if info.Age != nil {
		db = db.Where("age = ?", info.Age)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&myTestOnes).Error
	return myTestOnes, total, err
}
