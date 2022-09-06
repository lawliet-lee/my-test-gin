package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/dtx"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/my_test"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	My_testServiceGroup my_test.ServiceGroup
	DtxServiceGroup     dtx.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
