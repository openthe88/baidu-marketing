package oauth

// GetAccessTokenResponse 返回获取授权令牌接口
type GetAccessTokenResponse struct {
	//请求处理状态，0：处理成功，非0：处理失败
	Code int `json:"code,omitempty"`
	//请求处理话术
	Message string `json:"message,omitempty"`
	//请求处理结果
	Data AccessTokenData
}

type AccessTokenData struct {
	AccessToken        string `json:"accessToken,omitempty" dc:"授权令牌"`
	RefreshToken       string `json:"refreshToken,omitempty" dc:"刷新令牌"`
	OpenId             string `json:"openId,omitempty" dc:"获取授权用户信息标识"`
	ExpiresTime        string `json:"expiresTime,omitempty" dc:"授权令牌到期时间"`
	RefreshExpiresTime string `json:"refreshExpiresTime,omitempty" dc:"更新令牌到期时间"`
	ExpiresIn          int    `json:"expiresIn,omitempty" dc:"授权令牌剩余有效时间"`
	RefreshExpiresIn   int    `json:"refreshExpiresIn,omitempty" dc:"更新令牌剩余有效时间"`
	UserId             int    `json:"userId,omitempty" dc:"授权账号 ucid"`
}
