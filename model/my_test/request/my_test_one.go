package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/my_test"
)

type MyTestOneSearch struct {
	my_test.MyTestOne
	request.PageInfo
}
