package dtx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dtx"
	dtxReq "github.com/flipped-aurora/gin-vue-admin/server/model/dtx/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FieldDictionaryTypeApi struct {
}

var fieldDictionaryTypeService = service.ServiceGroupApp.DtxServiceGroup.FieldDictionaryTypeService

// CreateFieldDictionaryType 创建FieldDictionaryType
// @Tags FieldDictionaryType
// @Summary 创建FieldDictionaryType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dtx.FieldDictionaryType true "创建FieldDictionaryType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fieldDictionaryType/createFieldDictionaryType [post]
func (fieldDictionaryTypeApi *FieldDictionaryTypeApi) CreateFieldDictionaryType(c *gin.Context) {
	var fieldDictionaryType dtx.FieldDictionaryType
	_ = c.ShouldBindJSON(&fieldDictionaryType)
	verify := utils.Rules{
		"TypeCode": {utils.NotEmpty()},
		"TypeName": {utils.NotEmpty()},
	}
	if err := utils.Verify(fieldDictionaryType, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fieldDictionaryTypeService.CreateFieldDictionaryType(fieldDictionaryType); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFieldDictionaryType 删除FieldDictionaryType
// @Tags FieldDictionaryType
// @Summary 删除FieldDictionaryType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dtx.FieldDictionaryType true "删除FieldDictionaryType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fieldDictionaryType/deleteFieldDictionaryType [delete]
func (fieldDictionaryTypeApi *FieldDictionaryTypeApi) DeleteFieldDictionaryType(c *gin.Context) {
	var fieldDictionaryType dtx.FieldDictionaryType
	_ = c.ShouldBindJSON(&fieldDictionaryType)
	if err := fieldDictionaryTypeService.DeleteFieldDictionaryType(fieldDictionaryType); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFieldDictionaryTypeByIds 批量删除FieldDictionaryType
// @Tags FieldDictionaryType
// @Summary 批量删除FieldDictionaryType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FieldDictionaryType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fieldDictionaryType/deleteFieldDictionaryTypeByIds [delete]
func (fieldDictionaryTypeApi *FieldDictionaryTypeApi) DeleteFieldDictionaryTypeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := fieldDictionaryTypeService.DeleteFieldDictionaryTypeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFieldDictionaryType 更新FieldDictionaryType
// @Tags FieldDictionaryType
// @Summary 更新FieldDictionaryType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dtx.FieldDictionaryType true "更新FieldDictionaryType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fieldDictionaryType/updateFieldDictionaryType [put]
func (fieldDictionaryTypeApi *FieldDictionaryTypeApi) UpdateFieldDictionaryType(c *gin.Context) {
	var fieldDictionaryType dtx.FieldDictionaryType
	_ = c.ShouldBindJSON(&fieldDictionaryType)
	verify := utils.Rules{
		"TypeCode": {utils.NotEmpty()},
		"TypeName": {utils.NotEmpty()},
	}
	if err := utils.Verify(fieldDictionaryType, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fieldDictionaryTypeService.UpdateFieldDictionaryType(fieldDictionaryType); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFieldDictionaryType 用id查询FieldDictionaryType
// @Tags FieldDictionaryType
// @Summary 用id查询FieldDictionaryType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dtx.FieldDictionaryType true "用id查询FieldDictionaryType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fieldDictionaryType/findFieldDictionaryType [get]
func (fieldDictionaryTypeApi *FieldDictionaryTypeApi) FindFieldDictionaryType(c *gin.Context) {
	var fieldDictionaryType dtx.FieldDictionaryType
	_ = c.ShouldBindQuery(&fieldDictionaryType)
	if refieldDictionaryType, err := fieldDictionaryTypeService.GetFieldDictionaryType(fieldDictionaryType.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refieldDictionaryType": refieldDictionaryType}, c)
	}
}

// GetFieldDictionaryTypeList 分页获取FieldDictionaryType列表
// @Tags FieldDictionaryType
// @Summary 分页获取FieldDictionaryType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dtxReq.FieldDictionaryTypeSearch true "分页获取FieldDictionaryType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fieldDictionaryType/getFieldDictionaryTypeList [get]
func (fieldDictionaryTypeApi *FieldDictionaryTypeApi) GetFieldDictionaryTypeList(c *gin.Context) {
	var pageInfo dtxReq.FieldDictionaryTypeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := fieldDictionaryTypeService.GetFieldDictionaryTypeInfoList(pageInfo); err != nil {
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
