package reportv2

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

type GetUserDataRequest struct {
	ReportType string `json:"reportType" dc:"报表固定值"`
	//支持的时间单位： //HOUR：小时 //DAY：天 //WEEK：周 //MONTH：月 //SUMMARY：时间段汇总
	TimeUnit string `json:"timeUnit" dc:"时间单位"`
	//不同时间单位对应的日期格式：
	//DAY: 2021-03-31
	//HOUR: 2021-03-31 23:00
	//WEEK MONTH SUMMARY: 2021-03-01至2021-03-31，按照时间单位对应的周期汇总
	Date     string `json:"date" dc:"时间"`
	UserName string `json:"userName" dc:"账户"`
	//查询时使用key，数据返回value。
	//枚举值： 1:百度信息流 2:贴吧 3:好看视频 4:百度小说 5:百青藤 6:默认
	FeedFlowTypeEnum int    `json:"feedFlowTypeEnum,omitempty" dc:"流量类型"`
	WBudget          string `json:"WBudget,omitempty" dc:"预算"`
	BudgetEfficiency string `json:"budgetEfficiency,omitempty" dc:"预算利用率"`
	//查询时使用key，数据返回value。 枚举值： 0:计算机 1:移动设备 取值与旧版报告文档不一致，请注意区分。
	Device int `json:"device,omitempty" dc:"推广设备"`
	//该字段仅支持用于Filter中筛选数据使用。
	//查询时使用key，数据返回value。
	//枚举值：
	//0:全部 1:网站链接 2:应用推广(iOS) 3:应用推广(Android) 4:小程序 5:商品目录 6:门店推广 7:电商店铺
	FeedSubjectEnum int `json:"feedSubjectEnum,omitempty" dc:"FeedSubject"`
}

func (r GetUserDataRequest) Url() string {
	return fmt.Sprintf("%sOpenApiReportService/getReportData", model.BASE_URL_SMS)
}
