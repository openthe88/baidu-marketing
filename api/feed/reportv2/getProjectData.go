package reportv2

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/feed/reportv2"
)

// GetProjectData 计划报表
func GetProjectData(clt *core.SDKClient, auth model.RequestHeader, userDataRequest *reportv2.GetProjectDataRequest) (*model.ResponseHeader, []reportv2.ProjectData, error) {
	userDataRequest.ReportType = `2276038`
	req := &model.Request{
		Header: auth,
		Body:   userDataRequest,
	}
	var resp reportv2.GetProjectDataResult
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
