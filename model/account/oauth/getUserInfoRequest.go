package oauth

import (
	"fmt"
	"github.com/openthe88/baidu-marketing/model"
)

// GetUserInfoRequest 查询授权用户信息
type GetUserInfoRequest struct {
	OpenId      string `json:"openId" dc:"授权用户查询标识"`
	AccessToken string `json:"accessToken" dc:"已有的授权令牌"`
	UserId      int    `json:"userId" dc:"授权账号 ucid"`
	NeedSubList bool   `json:"needSubList,omitempty" dc:"是否需要子账号列表，值为true时返回subUserList"`
	PageSize    int    `json:"pageSize,omitempty" dc:"分页数量，默认100，最大不超过500"`
	//上一页返回的最大userid，用于子账号列表分页查询子账号列表时，该字段为必填。第一次获取子账户列表时，该字段需要设置为1
	LastPageMaxUcId int `json:"lastPageMaxUcId,omitempty" dc:"上一页返回的最大userid"`
}

func (r GetUserInfoRequest) Url() string {
	return fmt.Sprintf("%sgetUserInfo", model.BASE_URL_OAUTH)
}
