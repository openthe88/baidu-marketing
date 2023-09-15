package newvideo

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	video2 "github.com/openthe88/baidu-marketing/model/asset/newvideo"
)

// UploadVideo 视频上传接口，上传后搜索推广与信息流推广可共用
func UploadVideo(clt *core.SDKClient, auth model.RequestHeader, reqBody *video2.UploadVideoRequest) (*model.ResponseHeader, []video2.UploadVideoResponse, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp []video2.UploadVideoResponse
	header, err := clt.Do(req, &resp)
	return header, resp, err
}
