package creative

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/feed/creative"
)

// GetCreativeFeed 查询创意
func GetCreativeFeed(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.GetCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.GetCreativeResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
