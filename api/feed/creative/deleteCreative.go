package creative

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/feed/creative"
)

// DeleteCreative 删除推广创意
func DeleteCreative(clt *core.SDKClient, auth model.RequestHeader, creativeFeedIds []int64) (*model.ResponseHeader, error) {
	req := &model.Request{
		Header: auth,
		Body: creative.DeleteCreativeRequest{
			CreativeFeedIds: creativeFeedIds,
		},
	}

	return clt.Do(req, nil)
}
