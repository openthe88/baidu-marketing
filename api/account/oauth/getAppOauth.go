package oauth

import (
	"github.com/openthe88/baidu-marketing/model/account/oauth"
)

// GetAppOauth 授权地址
// 返回授权地址
func GetAppOauth(req oauth.GetAppOauthRequest) (string, error) {
	url := req.Url()
	return url, nil
}
