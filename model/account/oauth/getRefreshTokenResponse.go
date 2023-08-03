package oauth

// GetRefreshTokenResponse 刷新token接口
type GetRefreshTokenResponse struct {
	//请求处理状态，0：处理成功，非0：处理失败
	Code int `json:"code,omitempty"`
	//请求处理话术
	Message string `json:"message,omitempty"`
	//请求处理结果
	Data RefreshTokenData
}

type RefreshTokenData struct {
	AccessToken        string `json:"accessToken,omitempty" dc:"授权令牌"`
	RefreshToken       string `json:"refreshToken,omitempty" dc:"刷新令牌"`
	ExpiresTime        string `json:"expiresTime,omitempty" dc:"授权令牌到期时间"`
	RefreshExpiresTime string `json:"refreshExpiresTime,omitempty" dc:"更新令牌到期时间"`
	ExpiresIn          int    `json:"expiresIn,omitempty" dc:"授权令牌剩余有效时间"`
	RefreshExpiresIn   int    `json:"refreshExpiresIn,omitempty" dc:"更新令牌剩余有效时间"`
}
