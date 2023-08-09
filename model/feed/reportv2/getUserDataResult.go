package reportv2

import "encoding/json"

// GetUserDataResult 账户报表返回参
type GetUserDataResult struct {
	Data []UserData `json:"data,omitempty" dc:"数据"`
}

type UserData struct {
	Impression        int64       `json:"impression,omitempty" dc:"展现"`
	Click             int64       `json:"click,omitempty" dc:"点击"`
	Cost              json.Number `json:"cost,omitempty" dc:"消费"`
	Ctr               json.Number `json:"ctr,omitempty" dc:"消费"`
	Cpc               json.Number `json:"cpc,omitempty" dc:"平均点击价格"`
	Cpm               json.Number `json:"cpm,omitempty" dc:"千次展现消费"`
	PhoneButtonClicks int64       `json:"phoneButtonClicks,omitempty" dc:"组件点击次数"`
	Interaction       int64       `json:"interaction,omitempty" dc:"互动"`
}
