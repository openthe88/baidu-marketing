package report

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/enum"
	"github.com/openthe88/baidu-marketing/model"
)

type GetReportDataRequest struct {
	// ReportType 报告类型，唯一标识一个报告
	ReportType enum.ReportType `json:"reportType"`
	// UserIds 查询的用户ID，用于查询超管账户所管辖的多个账户数据。可为空，默认只查询当前账户数据。
	UserIds []int64 `json:"userIds,omitempty"`
	// TimeUnit 支持的时间单位： HOUR：小时 DAY：天 WEEK：周 MONTH：月 SUMMARY：时间段汇总
	TimeUnit enum.TimeUnit `json:"timeUnit"`
	// StartDate 数据的起始日期，格式 2020-05-28
	StartDate string `json:"startDate"`
	// EndDate 数据的结束日期，格式 2020-05-29
	EndDate string `json:"endDate"`
	// Columns 查询的列，包含属性和转化指标，必填项，至少要带一个转化指标。 每个报告的说明文档里都有支持的columns列表及说明。
	Columns []string `json:"columns"`
	// Sorts 排序信息，详见下方排序说明
	Sorts []Sort `json:"sorts,omitempty"`
	// Filters 筛选条件集合，非必填，详见下方过滤条件说明。
	Filters []Filter `json:"filters,omitempty"`
	// StartRow 从第几行开始获取结果
	StartRow int64 `json:"startRow,omitempty"`
	// RowCount 要获取多少行，和startRow配合使用，用于分页获取数据。
	RowCount int64 `json:"rowCount,omitempty"`
	// NeedSum 	是否需总计，非必填，默认不需要总计。
	NeedSum bool `json:"needSum,omitempty"`
}

func (r GetReportDataRequest) Url() string {
	return fmt.Sprintf("%sOpenApiReportService/getReportData", model.BASE_URL_SMS)
}

type Sort struct {
	// Column 排序的列名。最多支持添加2个排序规则；排序的列必须在请求的列里
	Column string `json:"column,omitempty"`
	// SortRule 排序规则。ASC:正序，DESC：倒序
	SortRule enum.SortRule `json:"sortRule,omitempty"`
}

type Filter struct {
	// Column 过滤的列名。
	Column string `json:"column,omitempty"`
	// Operator 过滤的操作符。
	Operator enum.Operator `json:"operator,omitempty"`
	// Values 过滤值。 对于Enum类型的列，使用key作为value过滤。
	Values []string `json:"values,omitempty"`
}
