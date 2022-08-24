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

type MyTestOneApi struct {
}

var myTestOneService = service.ServiceGroupApp.My_testServiceGroup.MyTestOneService

// CreateMyTestOne 创建MyTestOne
// @Tags MyTestOne
// @Summary 创建MyTestOne
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_test.MyTestOne true "创建MyTestOne"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myTestOne/createMyTestOne [post]
func (myTestOneApi *MyTestOneApi) CreateMyTestOne(c *gin.Context) {
	var myTestOne my_test.MyTestOne
	_ = c.ShouldBindJSON(&myTestOne)
	if err := myTestOneService.CreateMyTestOne(myTestOne); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMyTestOne 删除MyTestOne
// @Tags MyTestOne
// @Summary 删除MyTestOne
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_test.MyTestOne true "删除MyTestOne"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /myTestOne/deleteMyTestOne [delete]
func (myTestOneApi *MyTestOneApi) DeleteMyTestOne(c *gin.Context) {
	var myTestOne my_test.MyTestOne
	_ = c.ShouldBindJSON(&myTestOne)
	if err := myTestOneService.DeleteMyTestOne(myTestOne); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMyTestOneByIds 批量删除MyTestOne
// @Tags MyTestOne
// @Summary 批量删除MyTestOne
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MyTestOne"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /myTestOne/deleteMyTestOneByIds [delete]
func (myTestOneApi *MyTestOneApi) DeleteMyTestOneByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := myTestOneService.DeleteMyTestOneByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMyTestOne 更新MyTestOne
// @Tags MyTestOne
// @Summary 更新MyTestOne
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body my_test.MyTestOne true "更新MyTestOne"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /myTestOne/updateMyTestOne [put]
func (myTestOneApi *MyTestOneApi) UpdateMyTestOne(c *gin.Context) {
	var myTestOne my_test.MyTestOne
	_ = c.ShouldBindJSON(&myTestOne)
	if err := myTestOneService.UpdateMyTestOne(myTestOne); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMyTestOne 用id查询MyTestOne
// @Tags MyTestOne
// @Summary 用id查询MyTestOne
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query my_test.MyTestOne true "用id查询MyTestOne"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /myTestOne/findMyTestOne [get]
func (myTestOneApi *MyTestOneApi) FindMyTestOne(c *gin.Context) {
	var myTestOne my_test.MyTestOne
	_ = c.ShouldBindQuery(&myTestOne)
	if remyTestOne, err := myTestOneService.GetMyTestOne(myTestOne.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remyTestOne": remyTestOne}, c)
	}
}

// GetMyTestOneList 分页获取MyTestOne列表
// @Tags MyTestOne
// @Summary 分页获取MyTestOne列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query my_testReq.MyTestOneSearch true "分页获取MyTestOne列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /myTestOne/getMyTestOneList [get]
func (myTestOneApi *MyTestOneApi) GetMyTestOneList(c *gin.Context) {
	var pageInfo my_testReq.MyTestOneSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := myTestOneService.GetMyTestOneInfoList(pageInfo); err != nil {
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
