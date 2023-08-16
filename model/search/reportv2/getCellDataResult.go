package reportv2

import "encoding/json"

// GetCellDataResult 单元报表返回参
type GetCellDataResult struct {
	Data []CellData `json:"data,omitempty" dc:"数据"`
}

type CellData struct {
	RowCount      int            `json:"rowCount,omitempty" dc:"记录数"`
	TotalRowCount int            `json:"totalRowCount,omitempty" dc:"总记录数"`
	Rows          []CellDataInfo `json:"rows,omitempty" dc:"数据列表"`
}

type CellDataInfo struct {
	Date              string      `json:"date,omitempty" dc:"日期"`
	UserName          string      `json:"userName,omitempty" dc:"账号名称"`
	CampaignId        int64       `json:"campaignId,omitempty" dc:"推广计划ID"`
	CampaignStatus    string      `json:"campaignStatus,omitempty" dc:"推广计划状态"`
	CampaignName      string      `json:"campaignName,omitempty" dc:"推广计划"`
	AdGroupId         int64       `json:"adGroupId,omitempty" dc:"推广单元ID"`
	AdGroupStatus     string      `json:"adGroupStatus,omitempty" dc:"推广单元状态"`
	AdGroupName       string      `json:"adGroupName,omitempty" dc:"推广单元"`
	Impression        int64       `json:"impression,omitempty" dc:"展现"`
	Click             int64       `json:"click,omitempty" dc:"点击"`
	Cost              json.Number `json:"cost,omitempty" dc:"消费"`
	Ctr               json.Number `json:"ctr,omitempty" dc:"消费"`
	Cpc               json.Number `json:"cpc,omitempty" dc:"平均点击价格"`
	Cpm               json.Number `json:"cpm,omitempty" dc:"千次展现消费"`
	TopPvWinA         json.Number `json:"topPvWinA,omitempty" dc:"上方展现胜出率"`
	TopPvWinP         json.Number `json:"topPvWinP,omitempty" dc:"上方展现胜出率（精确)"`
	TopPageViews      json.Number `json:"topPageViews,omitempty" dc:"上方位展现"`
	TopPClicks        json.Number `json:"topPClicks,omitempty" dc:"上方位点击"`
	TopPay            json.Number `json:"topPay,omitempty" dc:"上方位消费"`
	TopFirstPageViews json.Number `json:"topFirstPageViews,omitempty" dc:"上方位首位展现"`
	TopFirstPvWinA    json.Number `json:"topFirstPvWinA,omitempty" dc:"上方首位胜出率"`
	TopFirstPvWinP    json.Number `json:"topFirstPvWinP,omitempty" dc:"上方首位胜出率（精确)"`
}
