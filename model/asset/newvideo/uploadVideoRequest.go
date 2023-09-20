package newvideo

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
	"io"
)

// UploadVideoRequest 上传视频
type UploadVideoRequest struct {
	File      io.Reader   `json:"fileUrl" dc:"必填，文件reader流"`
	Signature string      `json:"signature" dc:"必填，视频文件md5签名"`
	Params    VideoParams `json:"params" dc:"必填，标准json格式，接口业务请求参数"`
}

func (r UploadVideoRequest) Url() string {
	return fmt.Sprintf("%sVideoUploadService/addVideo", model.BASE_URL_SMS)
}

type VideoParams struct {
	Format string `json:"format" dc:"必填，上传视频的格式，如:mp4"`
	//必填，视频来源产品线：
	//2 - 信息流推广
	//13 - 搜索推广
	//20 - 基木鱼
	//35 - 电商
	Source       int    `json:"source" dc:"必填，视频来源产品线"`
	VideoName    string `json:"videoName" dc:"填，视频名称"`
	HttpProtocol string `json:"httpProtocol,omitempty" dc:"选填，默认 HTTP， 可选值:HTTP、 HTTPS"`
}
