package app

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/search/app"
)

// GetAppList 获取APP素材
func GetAppList(clt *core.SDKClient, auth model.RequestHeader, reqBody *app.GetAppListRequest) (*model.ResponseHeader, []app.AppInfoItemList, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp app.GetAppListResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
