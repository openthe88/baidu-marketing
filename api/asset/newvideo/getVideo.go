package newvideo

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	video2 "github.com/openthe88/baidu-marketing/model/asset/newvideo"
)

// GetVideo 视频获取接口，上传后搜索推广与信息流推广可共用
func GetVideo(clt *core.SDKClient, auth model.RequestHeader, reqBody *video2.GetVideoRequest) (*model.ResponseHeader, []video2.GetVideoResponse, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp []video2.GetVideoResponse
	header, err := clt.Do(req, &resp)
	return header, resp, err
}
