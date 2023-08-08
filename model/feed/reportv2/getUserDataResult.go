package reportv2

import "encoding/json"

// GetUserDataResult 账户报表返回参
type GetUserDataResult struct {
	// ID 请求对象的ID app下载报告/推广电话报告，请求对象的ID，为null
	ID json.Number `json:"ID,omitempty"`
	// Name 请求对象的name，数组形式，不同报告形式的不同定义规则定义如下文
	Name []string `json:"name,omitempty"`
	// RelatedId 依赖ID仅在 reportType=9，relatedId为adgroupId reportType=3，relatedId为userId reportType=5，relatedId为省市一级地域代码 其余情况为Null。
	RelatedId int `json:"relatedId,omitempty"`
	// Date 统计开始时间
	Date string `json:"date,omitempty"`
	// KPIs 按照请求顺序，返回KPI数据数组
	KPIs []json.Number `json:"KPIs,omitempty"`
	// TotalNumber 记录总条数
	TotalNumber int64 `json:"totalNumber,omitempty"`
	// TotalRowNumber 记录总条数
	TotalRowNumber json.Number `json:"totalRowNumber,omitempty"`
	// PageIndex 当前页码
	PageIndex int        `json:"pageIndex,omitempty"`
	Data      []UserData `json:"data,omitempty" dc:"数据"`
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
