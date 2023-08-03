package oauth

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

// GetRefreshTokenRequest 刷新token
type GetRefreshTokenRequest struct {
	//appId：填充为需要授权的应用 ID，可从应用管理界面获取。
	AppId string `json:"appId,omitempty" dc:"appid"`
	//已有的更新令牌 refreshToken
	RefreshToken string `json:"refreshToken,omitempty" dc:"已有的更新令牌 refreshToken"`
	//应用持有的 secretKey
	SecretKey string `json:"secretKey,omitempty"`
	//从授权回调地址中获取的管家账户ID
	UserId int64 `json:"userId,omitempty"`
}

func (r GetRefreshTokenRequest) Url() string {
	return fmt.Sprintf("%srefreshToken", model.BASE_URL_OAUTH)
}
