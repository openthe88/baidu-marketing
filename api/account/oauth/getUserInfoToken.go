package oauth

import (
	"encoding/json"
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model/account/oauth"
)

func GetUserInfo(clt *core.SDKClient, req oauth.GetUserInfoRequest) (oauth.GetUserInfoResponse, error) {
	reqBytes, _ := json.Marshal(req)
	url := req.Url()

	var resp oauth.GetUserInfoResponse
	err := clt.Post(url, reqBytes, &resp)
	return resp, err
}
