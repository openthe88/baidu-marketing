package newimage

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	image2 "github.com/openthe88/baidu-marketing/model/asset/newimage"
)

// GetImage 图片拉取接口
func GetImage(clt *core.SDKClient, auth model.RequestHeader, reqBody *image2.GetImageRequest) (*model.ResponseHeader, image2.GetImageResponse, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp image2.GetImageResponse
	header, err := clt.Do(req, &resp)
	return header, resp, err
}
