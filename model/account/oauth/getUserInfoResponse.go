package oauth

// GetUserInfoResponse 查询授权用户信息
type GetUserInfoResponse struct {
	//请求处理状态，0：处理成功，非0：处理失败
	Code int `json:"code,omitempty"`
	//请求处理话术
	Message string `json:"message,omitempty"`
	//请求处理结果
	Data UserInfoData
}

type UserInfoData struct {
	MasterUid    int               `json:"masterUid,omitempty" dc:"同意授权用户ucid"`
	MasterName   string            `json:"masterName,omitempty" dc:"同意授权用户对应的ucName"`
	UserAcctType int               `json:"userAcctType,omitempty" dc:"授权账户类型 1普通 2超管"`
	HasNext      bool              `json:"hasNext,omitempty" dc:"是否有下一页子账号列表"`
	SubUserList  []subUserListItem `json:"subUserList" dc:"同意授权用户关联的子账号列表"`
}

type subUserListItem struct {
	Ucid   int    `json:"ucid,omitempty" dc:"同意授权用户关联的子账号ucid"`
	UcName string `json:"ucName,omitempty" dc:"同意授权用户关联的子账号ucName"`
}
