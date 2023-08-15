package reportv2

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/search/reportv2"
)

// GetCellData 单元报表
func GetCellData(clt *core.SDKClient, auth model.RequestHeader, cellDataRequest *reportv2.GetCellDataRequest) (*model.ResponseHeader, reportv2.GetCellDataResult, error) {
	cellDataRequest.ReportType = `2284618`
	req := &model.Request{
		Header: auth,
		Body:   cellDataRequest,
	}
	var resp reportv2.GetCellDataResult
	header, err := clt.Do(req, &resp)
	return header, resp, err
}
