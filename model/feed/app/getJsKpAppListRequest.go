package app

import (
	"fmt"

	"github.com/openthe88/baidu-marketing/model"
)

// GetJsKpAppListRequest 查询APP信息 API Request
type GetJsKpAppListRequest struct {
}

func (r GetJsKpAppListRequest) Url() string {
	return fmt.Sprintf("%sAppFeedService/getJsKpAppList", model.BASE_URL_FEED)
}
