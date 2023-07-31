package report

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/report"
)

// GetAdgroupReportData 数据报表
func GetAdgroupReportData(clt *core.SDKClient, auth model.RequestHeader, reportRequest report.GetReportDataRequest) (*model.ResponseHeader, []report.GetAdgroupReportData, error) {
	req := &model.Request{
		Header: auth,
		Body:   reportRequest,
	}
	var resp report.GetAdgroupReportDataResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
