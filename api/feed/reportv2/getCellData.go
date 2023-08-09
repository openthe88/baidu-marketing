package reportv2

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/feed/reportv2"
)

// GetUserData 单元报表
func GetCellData(clt *core.SDKClient, auth model.RequestHeader, userDataRequest *reportv2.GetCellDataRequest) (*model.ResponseHeader, []reportv2.CellData, error) {
	userDataRequest.ReportType = `2330652`
	req := &model.Request{
		Header: auth,
		Body:   userDataRequest,
	}
	var resp reportv2.GetCellDataResult
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
