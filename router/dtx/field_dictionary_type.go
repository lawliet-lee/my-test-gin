package dtx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FieldDictionaryTypeRouter struct {
}

// InitFieldDictionaryTypeRouter 初始化 FieldDictionaryType 路由信息
func (s *FieldDictionaryTypeRouter) InitFieldDictionaryTypeRouter(Router *gin.RouterGroup) {
	fieldDictionaryTypeRouter := Router.Group("fieldDictionaryType").Use(middleware.OperationRecord())
	fieldDictionaryTypeRouterWithoutRecord := Router.Group("fieldDictionaryType")
	var fieldDictionaryTypeApi = v1.ApiGroupApp.DtxApiGroup.FieldDictionaryTypeApi
	{
		fieldDictionaryTypeRouter.POST("createFieldDictionaryType", fieldDictionaryTypeApi.CreateFieldDictionaryType)             // 新建FieldDictionaryType
		fieldDictionaryTypeRouter.DELETE("deleteFieldDictionaryType", fieldDictionaryTypeApi.DeleteFieldDictionaryType)           // 删除FieldDictionaryType
		fieldDictionaryTypeRouter.DELETE("deleteFieldDictionaryTypeByIds", fieldDictionaryTypeApi.DeleteFieldDictionaryTypeByIds) // 批量删除FieldDictionaryType
		fieldDictionaryTypeRouter.PUT("updateFieldDictionaryType", fieldDictionaryTypeApi.UpdateFieldDictionaryType)              // 更新FieldDictionaryType
	}
	{
		fieldDictionaryTypeRouterWithoutRecord.GET("findFieldDictionaryType", fieldDictionaryTypeApi.FindFieldDictionaryType)       // 根据ID获取FieldDictionaryType
		fieldDictionaryTypeRouterWithoutRecord.GET("getFieldDictionaryTypeList", fieldDictionaryTypeApi.GetFieldDictionaryTypeList) // 获取FieldDictionaryType列表
	}
}
