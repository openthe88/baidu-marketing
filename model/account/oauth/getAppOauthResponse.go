package oauth

// GetAppOauthResponse 返回授权地址
type GetAppOauthResponse struct {
	//授权地址
	Url string `json:"url,omitempty"`
}
