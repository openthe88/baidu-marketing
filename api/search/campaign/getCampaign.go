package campaign

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/search/campaign"
)

// GetCampaign 查询计划
// 根据指定的计划ID获取推广计划(ID可批量)
func GetCampaign(clt *core.SDKClient, auth model.RequestHeader, reqBody *campaign.GetCampaignRequest) (*model.ResponseHeader, []campaign.Campaign, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp campaign.GetCampaignResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
