package oauth

import (
	"encoding/json"
	"errors"
	"github.com/openthe88/baidu-marketing/core"
	"github.com/openthe88/baidu-marketing/model/account/oauth"
)

func GetAccessToken(clt *core.SDKClient, req oauth.GetAccessTokenRequest) (oauth.AccessTokenData, error) {
	reqBytes, _ := json.Marshal(req)
	url := req.Url()

	var resp oauth.GetAccessTokenResponse
	err := clt.Post(url, reqBytes, &resp)

	if resp.Code != 0 {
		return oauth.AccessTokenData{}, errors.New(resp.Message)
	}

	return resp.Data, err
}
