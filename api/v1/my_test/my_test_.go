package my_test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/my_test"
	my_testReq "github.com/flipped-aurora/gin-vue-admin/server/model/my_test/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MyTestApi struct {
}

var myTestService = service.ServiceGroupApp.My_testServiceGroup.MyTestService

// CreateMyTest 创建MyTest
// @Tags MyTest
// @Summary 创建MyTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_test.MyTest true "创建MyTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myTest/createMyTest [post]
func (myTestApi *MyTestApi) CreateMyTest(c *gin.Context) {
	var myTest my_test.MyTest
	_ = c.ShouldBindJSON(&myTest)
	if err := myTestService.CreateMyTest(myTest); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMyTest 删除MyTest
// @Tags MyTest
// @Summary 删除MyTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_test.MyTest true "删除MyTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /myTest/deleteMyTest [delete]
func (myTestApi *MyTestApi) DeleteMyTest(c *gin.Context) {
	var myTest my_test.MyTest
	_ = c.ShouldBindJSON(&myTest)
	if err := myTestService.DeleteMyTest(myTest); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMyTestByIds 批量删除MyTest
// @Tags MyTest
// @Summary 批量删除MyTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MyTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /myTest/deleteMyTestByIds [delete]
func (myTestApi *MyTestApi) DeleteMyTestByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := myTestService.DeleteMyTestByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMyTest 更新MyTest
// @Tags MyTest
// @Summary 更新MyTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_test.MyTest true "更新MyTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /myTest/updateMyTest [put]
func (myTestApi *MyTestApi) UpdateMyTest(c *gin.Context) {
	var myTest my_test.MyTest
	_ = c.ShouldBindJSON(&myTest)
	if err := myTestService.UpdateMyTest(myTest); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMyTest 用id查询MyTest
// @Tags MyTest
// @Summary 用id查询MyTest
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query my_test.MyTest true "用id查询MyTest"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /myTest/findMyTest [get]
func (myTestApi *MyTestApi) FindMyTest(c *gin.Context) {
	var myTest my_test.MyTest
	_ = c.ShouldBindQuery(&myTest)
	if remyTest, err := myTestService.GetMyTest(myTest.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remyTest": remyTest}, c)
	}
}

// GetMyTestList 分页获取MyTest列表
// @Tags MyTest
// @Summary 分页获取MyTest列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query my_testReq.MyTestSearch true "分页获取MyTest列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myTest/getMyTestList [get]
func (myTestApi *MyTestApi) GetMyTestList(c *gin.Context) {
	var pageInfo my_testReq.MyTestSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := myTestService.GetMyTestInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
