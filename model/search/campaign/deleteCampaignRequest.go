package campaign

import (
	"fmt"

	"github.com/openthe88/baidu-marketing/model"
)

// DeleteCampaignRequest 删除计划 API Request
type DeleteCampaignRequest struct {
	// CampaignIds 计划ID
	CampaignIds []int64 `json:"campaignIds,omitempty"`
}

func (r DeleteCampaignRequest) Url() string {
	return fmt.Sprintf("%sCampaignService/deleteCampaign", model.BASE_URL_SMS)
}
