package creative

import (
	"fmt"

	"github.com/openthe88/baidu-marketing/model"
)

// UpdateCreativeRequest 修改推广创意 API Request
type UpdateCreativeRequest struct {
	// CreativeTypes 更新推广创意字段;集合长度限制：[1, 3000];建议分批多次请求
	CreativeTypes []Creative `json:"creativeTypes,omitempty"`
}

func (r UpdateCreativeRequest) Url() string {
	return fmt.Sprintf("%sCreativeService/updateCreative", model.BASE_URL_SMS)
}
