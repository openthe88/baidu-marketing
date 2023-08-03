package oauth

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

// GetAccessTokenRequest 获取token
type GetAccessTokenRequest struct {
	//appId：填充为需要授权的应用 ID，可从应用管理界面获取。
	AppId string `json:"appId,omitempty"`
	//临时授权码
	AuthCode string `json:"authCode,omitempty"`
	//应用持有的 secretKey
	SecretKey string `json:"secretKey,omitempty"`
	//授权令牌获取模式，仅限：auth_code（授权码模式）
	GrantType string `json:"grantType,omitempty"`
	//从授权回调地址中获取的管家账户ID
	UserId int64 `json:"userId,omitempty"`
}

func (r GetAccessTokenRequest) Url() string {
	return fmt.Sprintf("%saccessToken", model.BASE_URL_OAUTH)
}
