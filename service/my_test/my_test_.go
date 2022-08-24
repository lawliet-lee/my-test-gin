package my_test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/my_test"
	my_testReq "github.com/flipped-aurora/gin-vue-admin/server/model/my_test/request"
)

type MyTestService struct {
}

// CreateMyTest 创建MyTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestService *MyTestService) CreateMyTest(myTest my_test.MyTest) (err error) {
	err = global.GVA_DB.Create(&myTest).Error
	return err
}

// DeleteMyTest 删除MyTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestService *MyTestService) DeleteMyTest(myTest my_test.MyTest) (err error) {
	err = global.GVA_DB.Delete(&myTest).Error
	return err
}

// DeleteMyTestByIds 批量删除MyTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestService *MyTestService) DeleteMyTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]my_test.MyTest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMyTest 更新MyTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestService *MyTestService) UpdateMyTest(myTest my_test.MyTest) (err error) {
	err = global.GVA_DB.Save(&myTest).Error
	return err
}

// GetMyTest 根据id获取MyTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestService *MyTestService) GetMyTest(id uint) (myTest my_test.MyTest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&myTest).Error
	return
}

// GetMyTestInfoList 分页获取MyTest记录
// Author [piexlmax](https://github.com/piexlmax)
func (myTestService *MyTestService) GetMyTestInfoList(info my_testReq.MyTestSearch) (list []my_test.MyTest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&my_test.MyTest{})
	var myTests []my_test.MyTest
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&myTests).Error
	return myTests, total, err
}
