package my_test

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MyTestRouter struct {
}

// InitMyTestRouter 初始化 MyTest 路由信息
func (s *MyTestRouter) InitMyTestRouter(Router *gin.RouterGroup) {
	myTestRouter := Router.Group("myTest").Use(middleware.OperationRecord())
	myTestRouterWithoutRecord := Router.Group("myTest")
	var myTestApi = v1.ApiGroupApp.My_testApiGroup.MyTestApi
	{
		myTestRouter.POST("createMyTest", myTestApi.CreateMyTest)             // 新建MyTest
		myTestRouter.DELETE("deleteMyTest", myTestApi.DeleteMyTest)           // 删除MyTest
		myTestRouter.DELETE("deleteMyTestByIds", myTestApi.DeleteMyTestByIds) // 批量删除MyTest
		myTestRouter.PUT("updateMyTest", myTestApi.UpdateMyTest)              // 更新MyTest
	}
	{
		myTestRouterWithoutRecord.GET("findMyTest", myTestApi.FindMyTest)       // 根据ID获取MyTest
		myTestRouterWithoutRecord.GET("getMyTestList", myTestApi.GetMyTestList) // 获取MyTest列表
	}
}
