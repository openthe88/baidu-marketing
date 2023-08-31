package reportv2

import "encoding/json"

// GetVideoDataResult 单元报表返回参
type GetVideoDataResult struct {
	Data []VideoData `json:"data,omitempty" dc:"数据"`
}

type VideoData struct {
	RowCount      int             `json:"rowCount,omitempty" dc:"记录数"`
	TotalRowCount int             `json:"totalRowCount,omitempty" dc:"总记录数"`
	Rows          []VideoDataInfo `json:"rows,omitempty" dc:"数据列表"`
}

type VideoDataInfo struct {
	Date              string      `json:"date,omitempty" dc:"日期"`
	UserName          string      `json:"userName,omitempty" dc:"账号名称"`
	IdeaId            int64       `json:"ideaId,omitempty" dc:"创意ID"`
	VideoId           int64       `json:"videoId,omitempty" dc:"视频ID"`
	VideoMD5          int64       `json:"videoMD5,omitempty" dc:"视频文件签名"`
	VideoName         int64       `json:"videoName,omitempty" dc:"视频名称"`
	VideoNameStatus   int64       `json:"videoNameStatus,omitempty" dc:"视频名称"`
	VideoInfo         VideoInfo   `json:"videoInfo,omitempty" dc:"视频内容预览"`
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
	Interaction       int64       `json:"interaction,omitempty" dc:"互动"`
	CompletePlayCount int64       `json:"completePlayCount,omitempty" dc:"播放完成数"`
	CompletePlayRatio json.Number `json:"completePlayRatio,omitempty" dc:"播放完成率"`
}

type VideoInfo struct {
	VideoThumbnail string `json:"videoThumbnail,omitempty" dc:"VideoThumbnail"`
	VideoURL       string `json:"videoURL,omitempty" dc:"视频URL"`
}
