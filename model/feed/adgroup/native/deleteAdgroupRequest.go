package native

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

// DeleteAdgroupRequest 删除单元 API Request
type DeleteAdgroupRequest struct {
	// AdgroupFeedIds 推广单元ID
	AdgroupFeedIds []int64 `json:"adgroupFeedIds"`
}

func (r DeleteAdgroupRequest) Url() string {
	return fmt.Sprintf("%sAdgroupFeedService/deleteAdgroupFeed", model.BASE_URL_FEED)
}
