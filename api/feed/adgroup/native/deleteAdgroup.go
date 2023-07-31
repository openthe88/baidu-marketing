package native

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/feed/adgroup/native"
)

// DeleteAdgroup 删除单元
func DeleteAdgroup(clt *core.SDKClient, auth model.RequestHeader, adgroupFeedIds []int64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: native.DeleteAdgroupRequest{
			AdgroupFeedIds: adgroupFeedIds,
		},
	}

	return clt.Do(req, nil)
}
