package my_test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MyTestOneRouter struct {
}

// InitMyTestOneRouter 初始化 MyTestOne 路由信息
func (s *MyTestOneRouter) InitMyTestOneRouter(Router *gin.RouterGroup) {
	myTestOneRouter := Router.Group("myTestOne").Use(middleware.OperationRecord())
	myTestOneRouterWithoutRecord := Router.Group("myTestOne")
	var myTestOneApi = v1.ApiGroupApp.My_testApiGroup.MyTestOneApi
	{
		myTestOneRouter.POST("createMyTestOne", myTestOneApi.CreateMyTestOne)             // 新建MyTestOne
		myTestOneRouter.DELETE("deleteMyTestOne", myTestOneApi.DeleteMyTestOne)           // 删除MyTestOne
		myTestOneRouter.DELETE("deleteMyTestOneByIds", myTestOneApi.DeleteMyTestOneByIds) // 批量删除MyTestOne
		myTestOneRouter.PUT("updateMyTestOne", myTestOneApi.UpdateMyTestOne)              // 更新MyTestOne
	}
	{
		myTestOneRouterWithoutRecord.GET("findMyTestOne", myTestOneApi.FindMyTestOne)       // 根据ID获取MyTestOne
		myTestOneRouterWithoutRecord.GET("getMyTestOneList", myTestOneApi.GetMyTestOneList) // 获取MyTestOne列表
	}
}
