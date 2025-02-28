package creative

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/search/creative"
)

// AddCreative 新增推广创意
func AddCreative(clt *core.SDKClient, auth model.RequestHeader, reqBody *creative.AddCreativeRequest) (*model.ResponseHeader, []creative.Creative, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp creative.AddCreativeResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
