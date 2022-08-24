package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/my_test"
)

type MyTestSearch struct {
	my_test.MyTest
	request.PageInfo
}
