package reportv2

import "encoding/json"

// GetProjectDataResult 单元报表返回参
type GetProjectDataResult struct {
	RowCount      int           `json:"rowCount,omitempty" dc:"记录数"`
	TotalRowCount int           `json:"totalRowCount,omitempty" dc:"总记录数"`
	Rows          []ProjectData `json:"rows,omitempty" dc:"数据"`
}

type ProjectData struct {
	Date              string      `json:"date,omitempty" dc:"日期"`
	UserName          string      `json:"userName,omitempty" dc:"账号名称"`
	CampaignId        int64       `json:"campaignId,omitempty" dc:"推广计划ID"`
	CampaignStatus    string      `json:"campaignStatus,omitempty" dc:"推广计划状态"`
	CampaignName      string      `json:"campaignName,omitempty" dc:"推广计划"`
	Impression        int64       `json:"impression,omitempty" dc:"展现"`
	Click             int64       `json:"click,omitempty" dc:"点击"`
	Cost              json.Number `json:"cost,omitempty" dc:"消费"`
	Ctr               json.Number `json:"ctr,omitempty" dc:"消费"`
	Cpc               json.Number `json:"cpc,omitempty" dc:"平均点击价格"`
	Cpm               json.Number `json:"cpm,omitempty" dc:"千次展现消费"`
	PhoneButtonClicks int64       `json:"phoneButtonClicks,omitempty" dc:"组件点击次数"`
	Interaction       int64       `json:"interaction,omitempty" dc:"互动"`
}
