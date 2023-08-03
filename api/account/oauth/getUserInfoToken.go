package oauth

import (
	"encoding/json"
	"errors"
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model/account/oauth"
)

func GetUserInfo(clt *core.SDKClient, req oauth.GetUserInfoRequest) (oauth.UserInfoData, error) {
	reqBytes, _ := json.Marshal(req)
	url := req.Url()

	var resp oauth.GetUserInfoResponse
	err := clt.Post(url, reqBytes, &resp)

	if resp.Code != 0 {
		return oauth.UserInfoData{}, errors.New(resp.Message)
	}

	return resp.Data, err
}
