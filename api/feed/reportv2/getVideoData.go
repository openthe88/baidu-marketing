package reportv2

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/feed/reportv2"
)

// GetVideoData 视频报表
func GetVideoData(clt *core.SDKClient, auth model.RequestHeader, videoDataRequest *reportv2.GetVideoDataRequest) (*model.ResponseHeader, reportv2.GetVideoDataResult, error) {
	videoDataRequest.ReportType = `114718`
	req := &model.Request{
		Header: auth,
		Body:   videoDataRequest,
	}
	var resp reportv2.GetVideoDataResult
	header, err := clt.Do(req, &resp)
	return header, resp, err
}
