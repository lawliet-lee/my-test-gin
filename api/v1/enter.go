package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/dtx"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/my_test"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	My_testApiGroup my_test.ApiGroup
	DtxApiGroup     dtx.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
