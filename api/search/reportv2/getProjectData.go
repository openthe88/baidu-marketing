package reportv2

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/search/reportv2"
)

// GetProjectData 计划报表
func GetProjectData(clt *core.SDKClient, auth model.RequestHeader, projectDataRequest *reportv2.GetProjectDataRequest) (*model.ResponseHeader, []reportv2.ProjectData, error) {
	projectDataRequest.ReportType = `2290316`
	req := &model.Request{
		Header: auth,
		Body:   projectDataRequest,
	}
	var resp reportv2.GetProjectDataResult
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
