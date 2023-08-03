package oauth

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

// GetAppOauthRequest 生成授权链接
type GetAppOauthRequest struct {
	//platformId：固定为 4960345965958561794
	//appId：填充为需要授权的应用 ID，可从应用管理界面获取。
	AppId string `json:"appId,omitempty" dc:"授权的应用ID"`
	//scope：应用权限代码，建议从应用管理界面系统生成的授权链接中获取。
	Scope string `json:"scope,omitempty" dc:"权限码"`
	//state：开发者自定义参数，长度限制 512 个字符，特殊字符需要 URLEncode。
	State string `json:"state,omitempty" dc:"自定义内容"`
	//callback：应用回调链接。
	Callback string `json:"callback,omitempty" dc:"回调地址"`
}

func (r GetAppOauthRequest) Url() string {
	return fmt.Sprintf("%spage/index?platformId=%s&appId=%s&scope=%s&state=%s&callback=%s", model.BASE_URL_OAUTH, `4960345965958561794`, r.AppId, r.Scope, r.State, r.Callback)
}
