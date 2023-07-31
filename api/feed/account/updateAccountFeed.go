package account

import (
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model"
	"github.com/openthe88/baidu-marketing/model/feed/account"
)

// UpdateAccountFeed 更新账户信息
func UpdateAccountFeed(clt *core.SDKClient, auth model.RequestHeader, budget float64) (*model.ResponseHeader, []account.Account, error) {
	req := &model.Request{
		Header: auth,
		Body: account.UpdateAccountFeedRequest{
			Budget: budget,
		},
	}
	var resp account.UpdateAccountFeedResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
