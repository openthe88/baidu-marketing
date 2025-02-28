package native

import (
	"github.com/openthe88/baidu-marketing/enum"
	"github.com/openthe88/baidu-marketing/model/feed/adgroup"
)

// Adgroup 推广单元
type Adgroup struct {
	// AdgroupId 推广单元ID
	AdgroupId int64 `json:"adgroupFeedId,omitempty"`
	// AdgroupName 推广单元名称
	AdgroupName string `json:"adgroupFeedName,omitempty"`
	// CampaignId 推广计划ID
	CampaignId int64 `json:"campaignFeedId,omitempty"`
	// Pause 暂停/启用推广单元
	Pause *bool `json:"pause,omitempty"`
	// Status 推广单元状态; 0 - 有效; 1 - 暂停推广; 2 - 推广计划暂停推广
	Status int `json:"status,omitempty"`
	// Audience 此字段为对象类型，对象内每个字段表示一种定向设置，所有字段的值均为String类型。每个单元可以设置0至多个定向设置。全部定向字段列表和详细说明请参考原生定向设置列表。
	Audience map[string]string `json:"audience,omitempty"`
	// Bid 出价，根据优化目标不同，具体含义如下：CPC：单次点击出价; CPM：千次展现出价; oCPC：第一阶段单次点击出价; eCPC：单次点击出价
	Bid float64 `json:"bid,omitempty"`
	// ProductTypes 投放版位
	ProductTypes []int `json:"producttypes,omitempty"`
	// FTypes 投放范围; 1 - 自定义类-百度信息流; 2 - 自定义类-贴吧; 4 - 百青藤; 8 - 自定义类-好看视频（好看视频流量目前在实验阶段，仅限已开通该流量的账户使用）; 64 - 自定义类-百度小说
	FTypes []int `json:"ftypes,omitempty"`
	// BidType 优化目标和付费模式; 1 - 点击（CPC）(应用推广营销目标下不支持，不影响存量物料投放); 2 - 曝光（CPM）(应用推广营销目标下不支持，不影响存量物料投放); 3 - 转化（oCPC/oCPM）; 5 - eCPC
	BidType int `json:"bidtype,omitempty"`
	// Ocpc oCPC信息; 本字段只有当bidtype=3有效。请参考oCPC设置对象;
	Ocpc *adgroup.Ocpc `json:"ocpc,omitempty"`
	// AtpFeedId 定向包ID; 定向包的优先级高于audience字段设置的定向信息。已绑定定向包单元解除绑定需传0
	AtpFeedId int64 `json:"atpFeedId,omitempty"`
	// DeliveryType 投放场景；默认值：0 取值范围：枚举值，列表如下 0 - 不限 1 - 开屏 ；开屏选项仅在投放范围选择百青藤时有效
	DeliveryType []int `json:"deliveryType,omitempty"`
	// ProductSetId 商品组id；仅商品目录营销目标单元可用，新增时与unitProducts至少填写一个，与unitProducts同时使用时unitProducts优先生效。参照商品组管理。
	ProductSetId int64 `json:"productSetId,omitempty"`
	// AddTime 添加时间
	AddTime string `json:"addTime,omitempty"`
	// UnitProducts 单元商品筛选设置；仅商品目录营销目标单元可用，新增时与productSetId至少填写一个，与productSetId同时使用时unitProducts优先生效
	UnitProducts []UnitProducts `json:"unitProducts,omitempty"`
}

type UnitProducts struct {
	// CatalogId 商品目录id；商品目录管理请登录百度商品中心
	CatalogId int64 `json:"catalogId,omitempty"`
	// RuleProducts 商品组规则
	RuleProducts []ProductSetRule `json:"ruleProducts,omitempty"`
}

// ProductSetRule 商品组规则
type ProductSetRule struct {
	// Field 字段名；参照查询商品库字段返回的字段value
	Field string `json:"field,omitempty"`
	// Operator 操作符。
	Operator enum.Operator `json:"operator,omitempty"`
	// Values 操作值。 操作字段值，多个操作值使用'|'分隔，多个之间取并集，最多15个值
	Values []string `json:"values,omitempty"`
}
