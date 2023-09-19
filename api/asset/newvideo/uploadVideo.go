package newvideo

import (
	"encoding/json"
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	video2 "github.com/openthe88/baidu-marketing/model/asset/newvideo"
)

// UploadVideo 视频上传接口，上传后搜索推广与信息流推广可共用
func UploadVideo(clt *core.SDKClient, auth model.RequestHeader, reqBody *video2.UploadVideoRequest) (*model.ResponseHeader, []video2.Video, error) {
	reqHeader := &model.Request{
		Header: auth,
		Body:   reqBody,
	}

	reqFile := &model.FileRequest{
		File:  reqBody.File,
		Name:  reqBody.Params.VideoName,
		Filed: `file`,
	}
	req := make(map[string]string)
	req[`userName`] = auth.TargetUserNameV2
	req[`accessToken`] = auth.AccessTokenV2
	req[`signature`] = reqBody.Signature
	params, _ := json.Marshal(reqBody.Params)
	req[`params`] = string(params)

	var resp video2.UploadVideoResponse
	header, err := clt.DoFile(reqHeader, &resp, reqFile, req)
	return header, resp.Data, err
}
