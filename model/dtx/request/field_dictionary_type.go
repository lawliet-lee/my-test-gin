package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dtx"
)

type FieldDictionaryTypeSearch struct {
	dtx.FieldDictionaryType
	request.PageInfo
}
