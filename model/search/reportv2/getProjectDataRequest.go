package reportv2

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

type GetProjectDataRequest struct {
	ReportType string `json:"reportType" dc:"报表固定值"`
	//支持的时间单位： //HOUR：小时 //DAY：天 //WEEK：周 //MONTH：月 //SUMMARY：时间段汇总
	TimeUnit string `json:"timeUnit" dc:"时间单位"`
	//不同时间单位对应的日期格式：
	//DAY: 2021-03-31
	//HOUR: 2021-03-31 23:00
	//WEEK MONTH SUMMARY: 2021-03-01至2021-03-31，按照时间单位对应的周期汇总
	Date     string `json:"date" dc:"时间"`
	UserName string `json:"userName" dc:"账户"`
	//报告数据中包含了已删除物料(计划、单元等)，已删除的物料，会在名称后面标记上"[已删除]"，比如"推广计划[已删除]"。
	//未删除物料正常显示名称，没有额外标记。
	CampaignNameStatus string `json:"campaignNameStatus,omitempty" dc:"推广计划"`
	BusinessPointId    int64  `json:"businessPointId,omitempty" dc:"推广业务ID"`
	BusinessPointName  string `json:"businessPointName,omitempty" dc:"推广业务"`
	TransPrice         string `json:"transPrice,omitempty" dc:"目标转化成本	"`
	//查询时使用key，数据返回value。 枚举值： 0:计算机 1:移动设备 取值与旧版报告文档不一致，请注意区分。
	CampaignId int64 `json:"campaignId,omitempty" dc:"推广计划ID"`
	//0：未删除，1：已删除。查询时使用key，数据返回value。枚举值：0:0 1:1
	CampaignStatus int    `json:"campaignStatus,omitempty" dc:"推广计划状态"`
	CampaignName   string `json:"campaignName,omitempty" dc:"推广计划"`
	Device         int    `json:"device,omitempty" dc:"推广设备"`
	//查询时使用key，数据返回value。
	//枚举值： 0:网站链接 1:应用推广 2:门店推广 3:推广营销活动 4:电商店铺推广 5:商品目录
	TargetingType int `json:"targetingType,omitempty" dc:"购买方式"`
	//该指标不支持小时
	ProvinceName     string `json:"provinceName,omitempty" dc:"省"`
	ProvinceCityName string `json:"provinceCityName,omitempty" dc:"城市"`
}

func (r GetProjectDataRequest) Url() string {
	return fmt.Sprintf("%sOpenApiReportService/getReportData", model.BASE_URL_SMS)
}
