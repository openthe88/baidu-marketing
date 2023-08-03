package oauth

import (
	"encoding/json"
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model/account/oauth"
)

func GetAccessToken(clt *core.SDKClient, req oauth.GetAccessTokenRequest) (oauth.GetAccessTokenResponse, error) {
	reqBytes, _ := json.Marshal(req)
	url := req.Url()

	var resp oauth.GetAccessTokenResponse
	err := clt.Post(url, reqBytes, &resp)
	return resp, err
}
