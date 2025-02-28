package reportv2

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

type GetUserDataRequest struct {
	ReportType string `json:"reportType" dc:"报表固定值"`
	//支持的时间单位： //HOUR：小时 //DAY：天 //WEEK：周 //MONTH：月 //SUMMARY：时间段汇总
	TimeUnit  string   `json:"timeUnit" dc:"时间单位"`
	StartDate string   `json:"startDate" dc:"开始时间"`
	EndDate   string   `json:"endDate" dc:"结束时间"`
	Columns   []string `json:"columns" dc:"属性字段"`
	Sorts     []Sorts  `json:"sorts" dc:"排序"`
	StartRow  int      `json:"startRow" dc:"开始记录数"`
	RowCount  int      `json:"rowCount" dc:"获取记录数"`
	NeedSum   bool     `json:"needSum" dc:"是否需要统计"`

	//不同时间单位对应的日期格式：
	//DAY: 2021-03-31
	//HOUR: 2021-03-31 23:00
	//WEEK MONTH SUMMARY: 2021-03-01至2021-03-31，按照时间单位对应的周期汇总
	//Date              string `json:"date" dc:"时间"`
	//UserName          string `json:"userName" dc:"账户"`
	//BusinessPointId   int64  `json:"businessPointId,omitempty" dc:"推广业务ID"`
	//BusinessPointName string `json:"businessPointName,omitempty" dc:"推广业务"`
	//WBudget           string `json:"WBudget,omitempty" dc:"预算"`
	//BudgetEfficiency  string `json:"budgetEfficiency,omitempty" dc:"预算利用率"`
	//查询时使用key，数据返回value。 枚举值： 0:计算机 1:移动设备 取值与旧版报告文档不一致，请注意区分。
	//Device int `json:"device,omitempty" dc:"推广设备"`
	//查询时使用key，数据返回value。
	//枚举值： 0:网站链接 1:应用推广 2:门店推广 3:推广营销活动 4:电商店铺推广 5:商品目录
	//MarketingTargetEnum int `json:"marketingTargetEnum,omitempty" dc:"营销目标"`
	//查询时使用key，数据返回value。
	//枚举值： 0:网站链接 1:应用推广 2:门店推广 3:推广营销活动 4:电商店铺推广 5:商品目录
	//TargetingType int `json:"targetingType,omitempty" dc:"购买方式"`
	//该指标不支持小时
	//ProvinceName     string `json:"provinceName,omitempty" dc:"省"`
	//ProvinceCityName string `json:"provinceCityName,omitempty" dc:"城市"`
}

type Sorts struct {
	Column   string `json:"column" dc:"排序字段"`
	SortRule string `json:"sortRule" dc:"排序类型"`
}

func (r GetUserDataRequest) Url() string {
	return fmt.Sprintf("%sOpenApiReportService/getReportData", model.BASE_URL_SMS)
}
