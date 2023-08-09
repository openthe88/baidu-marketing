package reportv2

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/search/reportv2"
)

// GetUserData 账户报表
func GetUserData(clt *core.SDKClient, auth model.RequestHeader, userDataRequest *reportv2.GetUserDataRequest) (*model.ResponseHeader, []reportv2.UserData, error) {
	userDataRequest.ReportType = `2208157`
	req := &model.Request{
		Header: auth,
		Body:   userDataRequest,
	}
	var resp reportv2.GetUserDataResult
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
