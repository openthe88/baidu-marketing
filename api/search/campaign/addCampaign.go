package campaign

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/search/campaign"
)

// AddCampaign 添加计划
// 新增推广计划，新增时可设置计划的属性设置。
func AddCampaign(clt *core.SDKClient, auth model.RequestHeader, reqBody *campaign.AddCampaignRequest) (*model.ResponseHeader, []campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp campaign.AddCampaignResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
